package main

import (
	"log"

	"github.com/dorm-parcel-manager/dpm/pkg/user/cmd"
)

func main() {
	err := cmd.RunServer()
	if err != nil {
		log.Fatal(err)
	}
}
