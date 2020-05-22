package sessions

import (
	"errors"
	"github.com/go_todo_sample/helpers/crypto"
	"net/http"
)

type DummyStore struct {
	database map[string]interface{}
}

var kvs DummyStore

func init() {
	kvs.database = map[string]interface{}{}
}

func NewDummyStore() *DummyStore {
	return &kvs
}

func (s *DummyStore) New(r *http.Request, cookieName string) (*DummySession, error) {
	cookie, err := r.Cookie(cookieName)
	if err == nil && s.Exists(cookie.Value) {
		return nil, errors.New("sessionID already exists")
	}

	session := NewDummySession(s, cookieName)
	session.ID = s.NewSessionID()
	session.request = r

	return session, nil
}

func (s *DummyStore) NewSessionID() string {
	return crypto.LongSecureRandomBase64()
}

func (s *DummyStore) Exists(sessionID string) bool {
	_, r := s.database[sessionID]
	return r
}

func (s *DummyStore) Save(r *http.Request, w http.ResponseWriter, session *DummySession) error {
	s.database[session.ID] = session

	c := &http.Cookie{
		Name:  session.Name(),
		Value: session.ID,
		Path:  "/",
	}

	http.SetCookie(session.writer, c)
	return nil
}

func (s *DummyStore) Delete(sessionID string) {
	delete(s.database, sessionID)
}

func (s *DummyStore) Get(r *http.Request, cookieName string) (*DummySession, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	sessionID := cookie.Value
	buffer, exists := s.database[sessionID]
	if !exists {
		return nil, errors.New("Invalid sessionID")
	}

	session := buffer.(*DummySession)
	session.request = r
	return session, nil
}
