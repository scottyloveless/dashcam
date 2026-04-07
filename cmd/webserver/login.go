package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/scottyloveless/dashcam/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) loginPostHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := app.queries.GetUserByEmail(r.Context(), email)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		app.logger.Error(err.Error())
		http.Redirect(w, r, "/login?error=invalid+email+or+password", http.StatusSeeOther)
		return
	}

	tokenBytes := make([]byte, 32)
	_, err = rand.Read(tokenBytes)
	if err != nil {
		app.logger.Error(err.Error())
		http.Redirect(w, r, "/login?error=internal+server+error", http.StatusSeeOther)
		return
	}

	rawToken := hex.EncodeToString(tokenBytes)
	hashBytes := sha256.Sum256([]byte(rawToken))
	tokenHash := hex.EncodeToString(hashBytes[:])

	if err = app.queries.CreateSession(r.Context(), database.CreateSessionParams{
		UserID:    user.ID,
		TokenHash: tokenHash,
	}); err != nil {
		app.logger.Error(err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    rawToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   3600,
	}
	http.SetCookie(w, cookie)

	if err = app.queries.UpdateLastLogin(r.Context(), user.ID); err != nil {
		app.logger.Error(err.Error())
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

type loginPageData struct {
	Error string
}

func (app *application) loginGetHandler(w http.ResponseWriter, r *http.Request) {
	data := loginPageData{
		Error: r.URL.Query().Get("error"),
	}
	err := app.templates["login"].Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		app.logger.Error(err.Error())
		return
	}
	rawToken := cookie.Value
	hashBytes := sha256.Sum256([]byte(rawToken))
	tokenHash := hex.EncodeToString(hashBytes[:])

	if err = app.queries.DeleteSession(r.Context(), tokenHash); err != nil {
		app.logger.Error(err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   -1,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
