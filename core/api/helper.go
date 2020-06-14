package api

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

func genUID() string {
	u := uuid.Must(uuid.NewV4(), errors.New("UUID Error"))
	return u.String()
}
