package socketcan

import (
	"encoding/binary"
	"fmt"

	"golang.org/x/sys/unix"
)

func NewRawInterface(ifName string) (Interface, error) {
	canIf := Interface{}
	canIf.ifType = IF_TYPE_RAW

	fd, err := unix.Socket(unix.AF_CAN, unix.SOCK_RAW, CAN_RAW)
	if err != nil {
		return canIf, err
	}

	// enable can-fd
	if err := unix.SetsockoptInt(fd, unix.SOL_CAN_RAW, unix.CAN_RAW_FD_FRAMES, 1); err != nil {
		return canIf, fmt.Errorf("set error canfd: %w", err)
	}

	ifIndex, err := getIfIndex(fd, ifName)
	if err != nil {
		return canIf, err
	}

	addr := &unix.SockaddrCAN{Ifindex: ifIndex}
	if err = unix.Bind(fd, addr); err != nil {
		return canIf, err
	}

	canIf.IfName = ifName
	canIf.SocketFd = fd

	return canIf, nil
}

func (i Interface) SendFrame(f CanFrame) error {
	if i.ifType != IF_TYPE_RAW {
		return fmt.Errorf("interface is not raw type")
	}

	var frameBytes []byte
	// assemble a SocketCAN frame
	if f.IsFD == true {
		frameBytes = make([]byte, 72)
	} else {
		frameBytes = make([]byte, 16)
	}
	// bytes 0-3: arbitration ID
	if f.ArbId < 0x800 {
		// standard ID
		binary.LittleEndian.PutUint32(frameBytes[0:4], f.ArbId)
	} else {
		// extended ID
		// set bit 31 (frame format flag (0 = standard 11 bit, 1 = extended 29 bit)
		binary.LittleEndian.PutUint32(frameBytes[0:4], f.ArbId|1<<31)
	}

	// byte 4: data length code
	frameBytes[4] = f.Dlc
	// data
	copy(frameBytes[8:], f.Data)

	_, err := unix.Write(i.SocketFd, frameBytes)
	return err
}

func (i Interface) RecvFrame() (CanFrame, error) {
	f := CanFrame{}

	if i.ifType != IF_TYPE_RAW {
		return f, fmt.Errorf("interface is not raw type")
	}

	// read SocketCAN frame from device
	frameBytes := make([]byte, 72)
	n, err := unix.Read(i.SocketFd, frameBytes)
	if err != nil {
		return f, err
	}
	//根据n判断是否是canfd
	if n == 16 {
		f.IsFD = false
	} else if n == 72 {
		f.IsFD = true
	}
	// bytes 0-3: arbitration ID
	f.ArbId = uint32(binary.LittleEndian.Uint32(frameBytes[0:4]))
	// remove bit 31: extended ID flag
	f.ArbId = f.ArbId & 0x7FFFFFFF
	// byte 4: data length code
	f.Dlc = frameBytes[4]
	// data
	f.Data = make([]byte, f.Dlc)
	copy(f.Data, frameBytes[8:f.Dlc+8])

	return f, nil
}
