package Utils

import (
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("something-very-secret"))
