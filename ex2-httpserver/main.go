package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartHttpServer(ctx context.Context,port int,name string) error {
	srv := &http.Server{Addr: fmt.Sprintf(":%d",port)}
	http.HandleFunc(name, func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("str"))
	})
	fmt.Printf("http server %s start\n",name)

	go func() {
		// 等待done
		// 外面报错此时会开始停止
		<-ctx.Done()
		fmt.Printf("http server %s stop , reason : %s\n",name,ctx.Err())
		err := srv.Shutdown(ctx)
		if err != nil && !errors.Is(err,ctx.Err()){
			fmt.Printf("stop server %s failed, reason : %s\n",name,err.Error())
		}
	}()

	err := srv.ListenAndServe()
	return err
}

func main() {
	// 使用 errgroup 进行 goroutine 取消
	group, ctx := errgroup.WithContext(context.Background())

	// create two servers

	// normal server
	group.Go(func() error {
		return StartHttpServer(ctx,8080,"server")
	})

	// debug server
	group.Go(func() error {
		return StartHttpServer(ctx,8888,"debug")
	})

	// receive os sig
	signals := make(chan os.Signal, 1)
	signal.Notify(signals)

	group.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("goroutine listening sig exit , reason : %s\n",ctx.Err())
				return nil
			case <-signals:
				// 接收到信号则终止
				fmt.Printf("receive os sig ...\n")
				return errors.New("os sig exit")
			}
		}
	})

	// 模拟运行3s后出error
	time.Sleep(3*time.Second)
	group.Go(func() error {
		return errors.New("?")
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group exit , error : ", err)
	}
}
