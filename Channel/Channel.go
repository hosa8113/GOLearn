package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	//channel ch erstellen
	ch := make(chan *TransferObjekt)
	//gepufferter channel ch erstellen, buffergrösse 10
	//ch := make(chan *TransferObjekt, 10)

	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println("from channel ", msg.Message, msg.Code)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := 0; i < 30; i++ {
		fmt.Println("from main-loop", i)
		x := TransferObjekt{
			Code:    i,
			Message: "Message-Id" + strconv.Itoa(i),
		}
		//channel ch nebenläufig befüllen, da sonst das pgm hier stehen bleibt, wenn der channel ungepuffert ist
		go func() { ch <- &x }()
	}

	time.Sleep(30 * time.Second)
	close(ch)
	fmt.Println("End of Program")
}

type TransferObjekt struct {
	Code    int
	Message string
}
