package api

import (
	"database/sql"
	"github.com/plutov/paypal/v3"
	"github.com/sergiosegrera/store/db"
	"os"
)

var (
	dbc *sql.DB
	pp  *paypal.Client
)

func init() {
	var err error
	dbc, err = db.NewDatabase()

	if err != nil {
		panic(err)
	}

	pp, err = paypal.NewClient(os.Getenv("PAYPAL_CLIENT"), os.Getenv("PAYPAL_SECRET"), paypal.APIBaseSandBox)

	if err != nil {
		panic(err)
	}

	_, err = pp.GetAccessToken()

	if err != nil {
		panic(err)
	}
}
