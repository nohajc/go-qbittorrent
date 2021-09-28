package main

import (
	"fmt"
	"log"

	"github.com/nohajc/go-qbittorrent/qbt"
)

func main() {
	qb := qbt.NewClient("http://localhost:8080")

	err := qb.Login(qbt.LoginOptions{
		Username: "admin",
		Password: "adminadmin",
	})
	if err != nil {
		log.Fatal(err)
	}

	torrents, err := qb.Torrents(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", torrents)
}
