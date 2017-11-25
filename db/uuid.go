package daryl_db

import (
	"github.com/satori/go.uuid"
)

func UUID() string {
	return uuid.NewV4().String()
}
