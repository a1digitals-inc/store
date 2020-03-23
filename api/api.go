package api

import (
	"database/sql"
	"github.com/sergiosegrera/store/db"
	"github.com/stripe/stripe-go"
	"os"
)

var (
	dbc *sql.DB
)

func init() {
	var err error
	dbc, err = db.NewDatabase()

	if err != nil {
		panic(err)
	}

	stripe.Key = os.Getenv("STRIPE_SECRET")

	if err != nil {
		panic(err)
	}

}
