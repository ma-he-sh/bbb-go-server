package functions

import (
	"net/http"

	env "github.com/devmarka/bbb-go-server/env"
	"github.com/gorilla/sessions"
)

var CookieStore = sessions.NewCookieStore([]byte(env.COOKIEHash())) // TODO use site sessiontoken

// SessionRedirect
func SessionRedirect(w http.ResponseWriter, r *http.Request, redirect string) {
	// redirect user if session exists
	session, _ := CookieStore.Get(r, "user_cookie")
	if auth, ok := session.Values["authed"].(bool); ok || auth {
		http.Redirect(w, r, redirect, http.StatusSeeOther)
	}
}

// SessionAuthCheck
func SessionAuthCheck(w http.ResponseWriter, r *http.Request, redirect string) {
	// redirect if user auth not available
	session, _ := CookieStore.Get(r, "user_cookie")
	if auth, ok := session.Values["authed"].(bool); !ok || !auth {
		http.Redirect(w, r, redirect, http.StatusSeeOther)
	}
}

// SessionAdminCheck
func SessionAdminCheck(w http.ResponseWriter, r *http.Request) bool {
	// session exist check
	session, _ := CookieStore.Get(r, "user_cookie")
	if auth, ok := session.Values["authed"].(bool); ok || auth {
		return true
	}
	return false
}
