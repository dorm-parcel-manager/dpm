package main

import (
	webpush "github.com/SherClockHolmes/webpush-go"
)

func main() {
	private, public, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		panic(err)
	}
	println("Public Key:", public)
	println("Private Key:", private)
}
