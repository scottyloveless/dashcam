package main

import (
	"crypto/rand"
	"crypto/sha256"
	"net/http"

	"github.com/scottyloveless/dashcam/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse email and password from the form
	//not sure
	email := "test"
	password := "password"

	// Call GetUserByEmail
	user, err := app.queries.GetUserByEmail(r.Context(), email)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	// Call bcrypt.CompareHashAndPassword
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		app.logger.Error(err.Error())
		return
	}

	// Generate a cryptographically random token with crypto/rand
	token := rand.Text()
	// Hash the token with sha256
	hash := sha256.New()
	hash.Write([]byte(token))
	tokenHash := string(hash.Sum(nil))
	// Call CreateSession with the hash
	if err = app.queries.CreateSession(r.Context(), database.CreateSessionParams{
		UserID:    user.ID,
		TokenHash: string(hash.Sum(nil)),
	}); err != nil {
		app.logger.Error(err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenHash,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   3600,
	}
	http.SetCookie(w, cookie)
	// Call UpdateLastLogin
	if err = app.queries.UpdateLastLogin(r.Context(), user.ID); err != nil {
		app.logger.Error(err.Error())
		return
	}
	// Redirect to the dashboard
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
