package sessions

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
	"net/http"
)

const (
	DefaultSessionName = "sample-sessions-default"
	DefaultCookieName  = "sample-session"
)

type DummySession struct {
	cookieName string
	ID         string
	store      *DummyStore
	request    *http.Request
	writer     http.ResponseWriter
	Values     map[string]interface{}
}

func NewDummySession(store *DummyStore, cookieName string) *DummySession {
	return &DummySession{
		cookieName: cookieName,
		store:      store,
		Values:     map[string]interface{}{},
	}
}

func StartSession(sessionName, cookieName string, store *DummyStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var session *DummySession
		var err error
		// リクエストからクッキーを受け取ってstoreからsessionを取り出す
		session, err = store.Get(ctx.Request, cookieName)
		if err != nil {
			// storeになければ作成
			session, err = store.New(ctx.Request, cookieName)
			if err != nil {
				println("Abort: " + err.Error())
				ctx.Abort()
			}
		}
		session.writer = ctx.Writer
		ctx.Set(sessionName, session)
		defer context.Clear(ctx.Request) // 遅延実行（関数処理の最後にリソース解放）
		ctx.Next()
	}
}

func StartDefaultSession(store *DummyStore) gin.HandlerFunc {
	return StartSession(DefaultSessionName, DefaultCookieName, store)
}

func GetSession(c *gin.Context, sessionName string) *DummySession {
	return c.MustGet(sessionName).(*DummySession)
}

func GetDefaultSession(c *gin.Context) *DummySession {
	return GetSession(c, DefaultSessionName)
}

func (s *DummySession) Name() string {
	return s.cookieName
}

func (s *DummySession) Get(key string) (interface{}, bool) {
	ret, exists := s.Values[key]
	return ret, exists
}

func (s *DummySession) Set(key string, val interface{}) {
	s.Values[key] = val
}

func (s *DummySession) Delete(key string) {
	delete(s.Values, key)
}

func (s *DummySession) Save() error {
	return s.store.Save(s.request, s.writer, s)
}
