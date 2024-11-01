package token

import "time"

// Maker is an interface for managing tokens

type Maker interface {
	//Creatername Token creates a new token for a specific username and duratiın
	CreateToken(username string, duration time.Duration) (string, error)

	//VerifyToken checs if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
