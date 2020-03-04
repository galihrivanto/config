package main

import (
	"context"
	"fmt"
	"time"

	"github.com/galihrivanto/config"
)

func main() {
	// // TODO: create runner here which watching remote config. for every changeset, dump its value
	// s := config.New(
	// 	config.Env(),
	// 	config.Yaml("/etc/config.yaml"),
	// 	config.Args(),
	// 	config.Etcd("localhost:2379"),
	// ).Watch()

	// dbhost := s.Get("db.host").String()
	// dbport := s.Get("db.port").Int()

	// <- s.Notify()
	ctx, cancel := context.WithCancel(context.Background())

	c := config.New(
		config.WithSource(
			config.File("test.json", true),
		),
		config.EnableWatcher(ctx, time.Second*5),
	)

	fmt.Println("name: ", c.Get("name").String(""))
	fmt.Println("slogan: ", c.Get("slogan").String(""))

	go func() {
		changed := c.Subscribe()

		for {
			<-changed
			fmt.Println("updated")

			fmt.Println("name: ", c.Get("name").String(""))
			fmt.Println("slogan: ", c.Get("slogan").String(""))

		}
	}()

	fmt.Scanln()
	fmt.Println("canceling")

	cancel()
}
