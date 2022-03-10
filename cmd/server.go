package cmd

import (
	"fmt"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
	"htSample/EndPoint"
	"htSample/Service"
	"htSample/Transport"
	"htSample/global"
	"net/http"
	"os"
	"os/signal"
	"sync"
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

func init() {
	ServerCommand.Flags = []cli.Flag{}
}

func Server(ctx *cli.Context) error {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(error); ok {
				// 异常捕获的处理
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)
	prepareDB(ctx, &wg)
	prepareWeb(ctx, &wg)
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

func prepareWeb(ctx *cli.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	// 1.先创建我们最开始定义的Server/server.go
	as := Service.AccountServer{}
	ps := Service.PaymentServer{}
	ts := Service.TransferServer{}
	// 2.在用EndPoint/endpoint.go 创建业务服务

	accountByCode := EndPoint.MakeServerEndPointAccountByCode(as)
	accountList := EndPoint.MakeServerEndPointAccountList(as)
	accountBlanceByCode := EndPoint.MakeServerEndPointAccountBalanceByCode(as)

	paymentList := EndPoint.MakeServerEndPointPaymentList(ps)
	paymentListByToken := EndPoint.MakeServerEndPointPaymentListByToken(ps)
	paymentListByAccount := EndPoint.MakeServerEndPointPaymentListByAccount(ps)

	transfer := EndPoint.MakeServerEndPointTransfer(ts)

	// 3.使用 kit 创建 handler
	// 固定格式
	// 传入 业务服务 以及 定义的 加密解密方法
	accountByCodeServer := httpTransport.NewServer(accountByCode, Transport.AccountByCodeDecodeRequest, Transport.AccountByCodeEncodeResponse)
	accountListServer := httpTransport.NewServer(accountList, Transport.AccountListDecodeRequest, Transport.AccountListEncodeResponse)
	accountBlanceByCodeServer := httpTransport.NewServer(accountBlanceByCode, Transport.AccountBalanceByCodeDecodeRequest, Transport.AccountBalanceByCodeEncodeResponse)

	paymentListServer := httpTransport.NewServer(paymentList, Transport.PaymentListDecodeRequest, Transport.PaymentListEncodeResponse)
	paymentListByTokenServer := httpTransport.NewServer(paymentListByToken, Transport.PaymentByTokenDecodeRequest, Transport.PaymentByTokenEncodeResponse)
	paymentListByAccountServer := httpTransport.NewServer(paymentListByAccount, Transport.PaymentByAccountDecodeRequest, Transport.PaymentByAccountEncodeResponse)

	transferServer := httpTransport.NewServer(transfer, Transport.TransferDecodeRequest, Transport.TransferByCodeEncodeResponse)

	//
	//go http.ListenAndServe("0.0.0.0:8000", helloServer)
	//go http.ListenAndServe("0.0.0.0:8001", sayServer)

	r := mux.NewRouter()

	r.Handle("/accountByCode", accountByCodeServer).Methods("GET")
	r.Handle("/accountList", accountListServer).Methods("GET")
	r.Handle("/accountBalanceByCode", accountBlanceByCodeServer).Methods("GET")

	r.Handle("/paymentList", paymentListServer).Methods("GET")
	r.Handle("/paymentListByToken", paymentListByTokenServer).Methods("GET")
	r.Handle("/paymentListByAccount", paymentListByAccountServer).Methods("GET")

	r.Handle("/transfer", transferServer).Methods("POST")

	go http.ListenAndServe(fmt.Sprintf("%v:%v", ctx.String("host"), ctx.Int("port")), r)

}

func prepareDB(ctx *cli.Context, wg *sync.WaitGroup) {
	var e error
	e = global.InitDB(ctx.String("dbHost"),
		ctx.Int("dbPort"),
		ctx.String("dbUser"),
		ctx.String("dbPwd"),
		ctx.String("dbName"),
	)

	if e != nil {
		fmt.Printf("the db connect failure.")
	} else {
		wg.Done()
	}

}
