package main

import (
	"github.com/google/uuid"
	"net/http"
)

type Session struct {
	ID          string
	State       string
	Username    string
	AccessToken string
}

var sessionStore map[string]Session

func getSession(req *http.Request) Session {
	SessionCookie, err := req.Cookie("sessionid")
	if err != nil || SessionCookie.Value == "" { // if cookie not present or no sessionid in cookie
		return Session{ID: uuid.New().String()}
	}

	session, sessionFound := sessionStore[SessionCookie.Value]
	if !sessionFound { // if session details not found in store
		return Session{ID: SessionCookie.Value}
	}

	return session
}

func putSession(res http.ResponseWriter, session Session) {

	// save to store
	sessionStore[session.ID] = session

	// set cookie
	http.SetCookie(res, &http.Cookie{
		Name:  "sessionid",
		Value: session.ID,
	})
}
