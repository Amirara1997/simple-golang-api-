package modelha



import (
	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	User  Users `json:"user"`
	jwt.StandardClaims
}