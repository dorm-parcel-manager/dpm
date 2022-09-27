package utils

import (
	"log"

	"github.com/pkg/errors"
)

func ErrorToFatal(fn func() error) func() {
	return func() {
		err := fn()
		if err != nil {
			log.Fatalln(errors.WithStack(err))
		}
	}
}
