package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/marcusolsson/focusd"
)

var browserClass = "Chromium"

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Reset on quit
	go func() {
		<-sigs
		setxkbmap("se")
		os.Exit(1)
	}()

	focus.Listen(func(e focus.Event) {
		switch ev := e.(type) {
		case focus.Gained:
			log.Printf("focus gained: %+v\n", ev.WMClass)
			if anyOf(browserClass, ev.WMClass...) {
				setxkbmap("se")
			}
		case focus.Lost:
			log.Printf("focus lost: %+v\n", ev.WMClass)
			if anyOf(browserClass, ev.WMClass...) {
				setxkbmap("us")
			}
		}
	})
}

func anyOf(str string, val ...string) bool {
	for _, v := range val {
		if str == v {
			return true
		}
	}
	return false
}

func setxkbmap(l string) {
	if err := exec.Command("setxkbmap", l).Run(); err != nil {
		log.Println(err)
	}
}
