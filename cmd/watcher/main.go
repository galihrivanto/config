package main

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/galihrivanto/config"
	"github.com/galihrivanto/x/lib/security"
)

var (
	configKey = "4755E2EDAB241FB28BC69E34783D2758"
)

func encodeFile(src, dest string) error {
	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		return errors.New("Configuration file not found")
	}

	// Read encrypted file content
	// #nosec
	plainText, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	key, err := hex.DecodeString(configKey)
	if err != nil {
		return err
	}

	// Encrypt the content
	cipherText, err := security.Encrypt(key, plainText)
	if err != nil {
		return err
	}

	encoded := []byte(hex.EncodeToString(cipherText))

	// Write to new file
	return ioutil.WriteFile(dest, encoded, os.ModePerm)
}

func decode(src []byte) []byte {
	cipherText, err := hex.DecodeString(string(src))
	if err != nil {
		return src
	}

	// Decode the encryption key
	key, err := hex.DecodeString(configKey)
	if err != nil {
		return src
	}

	plainText, err := security.Decrypt(key, cipherText)
	if err != nil {
		return src
	}

	return plainText
}

type decoder struct {}

func (d *decoder) Decode(src []byte) []byte {
	return decode(src)
}


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

	// for testing, encode sample file
	if err := encodeFile("test.json", "test-encoded.json"); err != nil {
		log.Fatal(err)
	}

	// <- s.Notify()
	ctx, cancel := context.WithCancel(context.Background())

	c := config.New(
		config.WithSource(
			config.File("test-encoded.json", true),
			&decoder{},
		),
		config.WithSource(
			config.File("test.yaml", true),
		),
		config.EnableWatcher(ctx, time.Second*5),
	)

	fmt.Println(string(c.Bytes()))

	go func() {
		changed := c.Subscribe()

		for {
			<-changed
			fmt.Println("updated")

			fmt.Println(string(c.Bytes()))
		}
	}()

	fmt.Scanln()
	fmt.Println("canceling")

	cancel()
}
