package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hellodudu/Ultimate/logger"
	ultimate "github.com/hellodudu/Ultimate/server"
)

func main() {
	umt, err := ultimate.NewUltimate()
	if err != nil {
		logger.Fatal(err)
	}

	umt.Run()

	// xmlloader
	// res.NewXmlLoader()

	// server exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-c
		logger.Print(fmt.Sprintf("ultimate server closing down (signal: %v)", sig))

		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			umt.Stop()
			logger.Print("server exit safely")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
	os.Exit(0)
}
