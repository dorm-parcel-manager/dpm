package cmd

import (
	"github.com/pkg/errors"
)

func RunServer() error {
	server, cleanup, err := InitializeServer()
	if err != nil {
		return errors.WithStack(err)
	}
	err = server.Start()
	defer cleanup()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
