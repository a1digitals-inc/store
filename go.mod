module github.com/sergiosegrera/store

go 1.13

replace github.com/plutov/paypal/v3 => ../paypal

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/gzip v0.0.1
	github.com/gin-gonic/gin v1.5.0
	github.com/lib/pq v1.3.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/plutov/paypal/v3 v3.0.12
	golang.org/x/crypto v0.0.0-20200311171314-f7b00557c8c4
)
