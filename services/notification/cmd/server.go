package cmd

import (
	"github.com/pkg/errors"
)

func RunServer() error {
	notificationServiceServer, cleanup, err := InitializeServer()
	if err != nil {
		return errors.WithStack(err)
	}
	defer cleanup()
	err = notificationServiceServer.Start()
	if err != nil {
		return err
	}
	return nil
}
