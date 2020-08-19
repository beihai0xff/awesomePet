# @Desperated 已作废，糟糕的设计与编码
# awesomePet

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![Support](https://img.shields.io/badge/ECHO-V4-yellow)](https://echo.labstack.com/)
![Support](https://img.shields.io/badge/Mysql-8.0%2B-green)

一个使用 Golang 编写的后端服务案例——在线宠物分享

## Usage
#### Requirements
Go version >= 1.11 and GO111MODULE=on；mysql >= 5.5.3 and utf8mb4；配置 conf.yaml 文件
#### Build & Run
```bash
$ git clone https://github.com/wingsxdu/awesomePet.git
$ cd Wade && go build
$ export GOPROXY=https://goproxy.io // 存在网络环境问题的可以设置代理
$ chmod a+x awesomePet // linux 下赋予文件执行权限
$ ./awesomePet
```
打开浏览器访问：<https://localhost:443/info> 查看请求信息

## Features
* 跨平台开发与交叉编译（启用CGO需关闭交叉编译）；
* 远程程序调用,跨平台、跨语言进程间通信；
* 高并发处理 ，性能优化；
* 数据可视化；

#### awesomePet 有哪些功能？
* [Echo](https://echo.labstack.com/)：高性能、可扩展、简约的的Go Web框架;
* [GORM](https://gorm.io/)：全功能数据库 orm 引擎；
* [go-echarts](https://go-echarts.chenjiandongx.com/)：Golang 数据可视化第三方库；
* [gRPC](https://grpc.io/)：高性能、开源的RPC框架，可跨语言；
* [Protocol Buffers](https://github.com/protocolbuffers/protobuf)：配合gRPC使用；

## others
#### About Author
欢迎反馈使用过程中遇到的问题，可用以下联系方式跟我交流：
* 邮箱：beihai@wingsxdu.com；
* blog：<https://www.wingsxdu.com> @beihai

#### Noticed
* TLS 使用自签名证书，有效期为一年，更新证书指令：
```bash
$ cd awesomePet && go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
```
* IDE: Goland，为方便作者在多设备上编程，.idea 文件夹一并上传到 github 上，可自行删除。

