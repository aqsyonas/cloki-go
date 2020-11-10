package service

import (
	"github.com/jmoiron/sqlx"
)

// Service : here you tell us what Salutation is
type ServiceData struct {
	Session *sqlx.DB
}
