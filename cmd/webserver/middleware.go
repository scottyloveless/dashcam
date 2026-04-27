package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"
)

func (app *application) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			app.logger.Error(err.Error())
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		rawToken := cookie.Value
		hashBytes := sha256.Sum256([]byte(rawToken))
		tokenHash := hex.EncodeToString(hashBytes[:])

		session, err := app.queries.GetSessionByTokenHash(r.Context(), tokenHash)
		if err != nil {
			app.logger.Error(err.Error())
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if session.ExpiresAt.Time.Before(time.Now()) {
			app.logger.Error("session expired")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
