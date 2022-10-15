package utils

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ErrorToFatal(fn func() error) func() {
	return func() {
		err := fn()
		if err != nil {
			log.Fatalln(errors.WithStack(err))
		}
	}
}

func NullTimeToTimestampPb(nullTime sql.NullTime) *timestamppb.Timestamp {
	if !nullTime.Valid {
		return nil
	}
	return timestamppb.New(nullTime.Time)
}
