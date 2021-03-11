package handlers

import (
	"ProjectMIRRO/backend/gateway/models/users"
	"ProjectMIRRO/backend/gateway/sessions"
)

// ContextHandler deals with the session id,
// the store of session and user
type ContextHandler struct {
	SessionID    string
	SessionStore sessions.Store
	UserStore    users.Store
}
