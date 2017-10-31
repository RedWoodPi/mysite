package main

import (
	"net/http"
	"log"
	"mysite/controllers"
	"time"
	"mysite/token"
)

//主程序
func main()  {
	go timeTike()
	http.HandleFunc("/check", controllers.CheckSignature)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Listenandserver: ", err)
	}
}

func timeTike(){
	token.SaveToken()
	for range time.Tick(7200 * time.Second){
		token.SaveToken()
	}
}

