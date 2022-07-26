package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/jyyds/sink/server"
)

func main() {
	go server.Run()

	cmd := startBrowser()

	// 等待中断信号
	chSingnal := listenToInterrupt()

	<-chSingnal
	cmd.Process.Kill()
}

func startBrowser() *exec.Cmd {

	// 启动Chrome
	chromePath := "C:\\Users\\86158\\AppData\\Local\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://localhost:27149/static/")
	cmd.Start()
	return cmd
}

func listenToInterrupt() chan os.Signal {
	// 监听中断信号
	chSingnal := make(chan os.Signal, 1)
	signal.Notify(chSingnal, os.Interrupt)
	return chSingnal
}
