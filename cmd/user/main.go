package main

import (
	"github.com/dorm-parcel-manager/dpm/pkg/user/cmd"
	"log"
)

func main() {
	err := cmd.RunServer()
	if err != nil {
		log.Fatal(err)
	}
}
