package functions

import (
	"github.com/gorilla/sessions"
	env "github.com/devmarka/bbb-go-server/env"
	"net/http"
)

var CookieStore = sessions.NewCookieStore([]byte(env.COOKIEHash())) // TODO use site sessiontoken

func SessionRedirect( w http.ResponseWriter, r *http.Request, redirect string ) {
	// redirect user if session exists
	session, _ := CookieStore.Get( r, "user_cookie" )
	if auth, ok := session.Values["authed"].(bool); ok || auth {
		http.Redirect( w, r, redirect, http.StatusSeeOther )
	}
}

func SessionAuthCheck( w http.ResponseWriter, r *http.Request, redirect string ) {
	// redirect if user auth not available
	session, _ := CookieStore.Get(r, "user_cookie")
	if auth, ok := session.Values["authed"].(bool); !ok || !auth {
		http.Redirect(w, r, redirect, http.StatusSeeOther)
	}
}