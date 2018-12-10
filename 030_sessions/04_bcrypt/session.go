package main

import (
	"github.com/google/uuid"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewUUID()
		c = &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
	}

	// assign user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	return u
}

func isLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	_, ok := dbSessions[c.Value]
	return ok

}

