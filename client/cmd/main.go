package main

import (
	"flag"
	"log"
	"spark-client/internal/client"
)

func main() {
	clientID := flag.String("id", "", "客户端ID")
	flag.Parse()

	if *clientID == "" {
		log.Fatal("必须提供客户端ID")
	}

	c := client.NewClient(*clientID)
	if err := c.Connect("101.132.157.32:1234"); err != nil {
		log.Fatal(err)
	}

	// 阻塞主线程
	select {}
}
