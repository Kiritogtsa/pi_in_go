package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Kiritogtsa/pi_in_go/conf"
	"github.com/Kiritogtsa/pi_in_go/handles"
)

// func server() {
// 	r := chi.NewRouter()
// 	r.Group(func(r chi.Router) {

//		})
//	}
func tempo(t chan string) {
	for i := range 1000 {
		select {
		case msg := <-t:
			if msg == "foi" {
				fmt.Println("abriu o servidor")
				return
			}
		default:
			fmt.Println("--", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
	log.Fatal("nao foi possivel abrir o servidor")
}

func main() {
	teste, err := conf.NewConf()
	if err != nil {
		fmt.Println("deu erro ", err)
	}
	testech := make(chan string)
	port := teste.App.Port
	go tempo(testech)
	handles.Serverinit(port)
}
