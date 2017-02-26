package db

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"

	"errors"
)

const (
	Name       = "testing"
	Collection = "sanitize"
)

// Connects to MongoDB. If a session is not supplied as an arg this creates a copy of a session and returns that
func Connect(s ...*mgo.Session) (*mgo.Session, error) {
	switch {
	// No session was supplied, create a new primary session
	case len(s) == 0:
		session, err := mgo.Dial("localhost")
		if err != nil {
			return nil, errors.New("Can not connect to database")
		}
		return session, nil
	// Session was supplied so make a copy and return that
	case len(s) == 1:
		return s[0].Copy(), nil
	// There are too many sessions supplied so we can't handle this at the moment
	default:
		return nil, errors.New("db.Connect can only create or copy a single session")
	}
}
