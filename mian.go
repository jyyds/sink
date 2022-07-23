package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	chromePath := "C:\\Users\\86158\\AppData\\Local\\Google\\Chrome\\Application"
	// var ui lorca.UI
	// currentDir, _ := os.Getwd()
	// dir := filepath.Join(currentDir, ".cache")
	// ui, _ = lorca.New("https://baidu.com", dir, 800, 600, "--disable-sync", "--disable-translate")
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	//case <-ui.Done():
	case <-chSignal:
	}
	//ui.Close()
}
