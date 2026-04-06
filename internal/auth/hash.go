package auth

import "golang.org/x/crypto/bcrypt"

func createHash(pw []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pw, 8)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
