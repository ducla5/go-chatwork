package main

import (
	"flag"
	"fmt"
	chatwork "github.com/ducla5/go-chatwork"
)

//var apiKey string

func init() {
	flag.StringVar(&apiKey, "key", "", "Chatwork API key")
	flag.Parse()
}

func main() {
	c := chatwork.NewClient(apiKey)
	fmt.Printf("%+v\n", c.MyStatus())
	fmt.Printf("%+v\n", c.MyTasks(map[string]string{}))
}
