package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/jyyds/sink/server"
)

func main() {
	go func() {
		server.Run()
	}()

	// 启动Chrome
	chromePath := "C:\\Users\\86158\\AppData\\Local\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://localhost:27149/static/")
	cmd.Start()

	// 监听中断信号
	chSingnal := make(chan os.Signal, 1)
	signal.Notify(chSingnal, os.Interrupt)

	// 等待中断信号
	<-chSingnal
	cmd.Process.Kill()
}
