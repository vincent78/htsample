package cmd

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/urfave/cli/v2"
	"htSample/EndPoint"
	"htSample/Service"
	"htSample/Transport"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	ServerCommand = &cli.Command{
		Action:    Server,
		Name:      "server",
		Usage:     "",
		ArgsUsage: " ",
		Category:  "",
		Description: `
The output of this command is supposed to be machine-readable.
`,
	}
)

func Server(ctx *cli.Context) error {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(error); ok {
				// 异常捕获的处理
			}
		}
	}()

	// 1.先创建我们最开始定义的Server/server.go
	s := Service.Server{}

	// 2.在用EndPoint/endpoint.go 创建业务服务
	hello := EndPoint.MakeServerEndPointHello(s)
	Bye := EndPoint.MakeServerEndPointBye(s)

	// 3.使用 kit 创建 handler
	// 固定格式
	// 传入 业务服务 以及 定义的 加密解密方法
	helloServer := httpTransport.NewServer(hello, Transport.HelloDecodeRequest, Transport.HelloEncodeResponse)
	sayServer := httpTransport.NewServer(Bye, Transport.ByeDecodeRequest, Transport.ByeEncodeResponse)

	// 使用http包启动服务
	go http.ListenAndServe("0.0.0.0:8000", helloServer)

	go http.ListenAndServe("0.0.0.0:8001", sayServer)

	// 堵塞主线程
	BlockBySignal()
	return nil
}

func BlockBySignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	for {
		s := <-c
		switch s {
		case os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGTERM:
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
