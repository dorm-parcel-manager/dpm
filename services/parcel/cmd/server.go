package cmd

import (
	"github.com/pkg/errors"
)

func RunServer() error {
	server, err := InitializeServer()
	if err != nil {
		return errors.WithStack(err)
	}
	err = server.Start()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
