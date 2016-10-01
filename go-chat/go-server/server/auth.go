package server

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"
	"time"
)

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

func basicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		if pair[0] == "" || pair[1] != os.Getenv("secret") {
			http.Error(w, "Not authorized", 401)
			return
		}

		expire := time.Now().Add(5 * time.Minute)
		cookie := &http.Cookie{Name: "friend-circle-user", Expires: expire, Value: pair[0]}
		r.AddCookie(cookie)

		h.ServeHTTP(w, r)
	}
}
