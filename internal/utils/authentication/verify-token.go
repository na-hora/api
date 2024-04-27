package authentication

import (
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string) (*jwt.Token, error) {
	return nil, nil
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	return secretKey, nil
	// })

	// // Check for verification errors
	// if err != nil {
	// 	return nil, err
	// }

	// // Check if the token is valid
	// if !token.Valid {
	// 	return nil, fmt.Errorf("invalid token")
	// }

	// // Return the verified token
	// return token, nil
}
