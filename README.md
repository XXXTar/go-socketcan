 # go-socketcan-fd

[![GitHub](https://img.shields.io/github/owner-repo-size/linklayer/go-socketcan.svg)](https://github.com/linklayer/go-socketcan)
[![GitHub Issues](https://img.shields.io/github/issues/linklayer/go-socketcan.svg)](https://github.com/linklayer/go-socketcan/issues)
[![License](https://img.shields.io/github/license/linklayer/go-socketcan.svg)](https://github.com/linklayer/go-socketcan/blob/master/LICENSE)

## 引言

`go-socketcan` 是一个对流行 Go 库 [go-socketcan](github.com/linklayer/go-socketcan) 的二次开发，该库用于与 CAN 套接字进行交互。这个版本增加了对 CAN FD（灵活数据速率）通信的支持，允许在 CAN 总线上以高速度传输数据。

## 功能

- **CAN FD 支持**：支持接收和发送最大 64 字节的 CAN FD 帧。
- **兼容性**：与原始 `go-socketcan` 库完全兼容。
- **易于安装**：可以使用 Go 的包管理器进行安装。
- **详细文档**：提供全面的文档和示例，便于集成。

## 安装

要安装 `go-socketcan`，你需要在系统上安装 Go。然后运行以下命令：

```sh
# set env only for first time
go env -w GOPROXY=https://mirrors.hirain.com/golang/,direct
go env -w GOPRIVATE=source-crdc.hirain.com

# install pkg
go get source-crdc.hirain.com/atx-go/go-socketcan
```

## 使用示例

参考 `cmd/`

