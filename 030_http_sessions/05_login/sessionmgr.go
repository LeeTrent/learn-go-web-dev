package main

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
)

type session struct {
	sessionId, userName string
}

type SessionMgr map[string]string // Session ID, User Name

func (sm SessionMgr) getUserName(sessionId string) string {
	return sm[sessionId]
}

func (sm SessionMgr) getSession(sessionId string) (session, error) {
	if userName, ok := sm[sessionId]; ok {
		return session{sessionId, userName}, nil
	}
	errMsg := fmt.Sprintf("Session for Session ID '%s' NOT FOUND", sessionId)
	return session{}, errors.New(errMsg)
}

func (sm SessionMgr) removeSession(sessionId string) {
	delete(sm, sessionId)
}

func (sm SessionMgr) createSession(userName string) session {
	sessionId := uuid.NewV4().String()
	sm[sessionId] = userName
	return session{sessionId, userName}
}

func (sm SessionMgr) createSessionId() string {
	return uuid.NewV4().String()
}
