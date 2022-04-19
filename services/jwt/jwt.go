package jwt

import (
	"github.com/cristalhq/jwt/v4"
	"log"
	"fmt"
	"time"
	"strconv"
	"encoding/json"
	"rgb/models"
)

var (
	jwtSigner   jwt.Signer
	jwtVerifier jwt.Verifier
)

func JwtSetup(JwtSecret string) {
	var err error
	key := []byte(JwtSecret)

	jwtSigner, err = jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		log.Println("Error creating jwt signer: ", err)
	}

	jwtVerifier, err = jwt.NewVerifierHS(jwt.HS256, key)
	if err != nil {
		log.Println("Error creating jwt verifier: ", err)
	}
}

func GenerateJWT(user models.User,) string {
	claims := &jwt.RegisteredClaims{
	  ID:        fmt.Sprint(user.ID),
	  ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
	}
	builder := jwt.NewBuilder(jwtSigner)
	token, err := builder.Build(claims)
	if err != nil {
	    log.Println("Error building jwt token: ", err)
	}
	return token.String()
  }

func VerifyJWT(tokenStr string) (int, error) {
	key := []byte(`test`)
	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
	if err != nil {
		log.Println("Error creating jwt verifier: ", err)
	}	

	token, err := jwt.Parse([]byte(tokenStr),verifier)
	if err != nil {
		log.Println("Error parsing jwt : ", err)
		return 0, err
	}

	if err := jwtVerifier.Verify(token); err != nil {
		log.Println("Error verifying token : ", err)
		return 0, err
	}

	// get Registered claims
	var claims jwt.RegisteredClaims
	if err := json.Unmarshal(token.Claims(), &claims); err != nil {
		log.Println("Error unmarshalling JWT claims : ", err)
		return 0, err
	}

	var _ bool = claims.IsForAudience("admin")
	var _ bool = claims.IsValidAt(time.Now())

	id, err := strconv.Atoi(claims.ID)

	return id, nil
}
