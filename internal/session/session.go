package session

import (
	"fmt"
	"sync"

	"github.com/charmbracelet/ssh"
)

type SessionManager struct {
	sync.RWMutex
	sessions map[string]struct {
		sess ssh.Session
		nick string
	}
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]struct {
			sess ssh.Session
			nick string
		}),
	}
}

func (sm *SessionManager) Add(id string, sess ssh.Session, nick string) {
	sm.Lock()
	defer sm.Unlock()
	sm.sessions[id] = struct {
		sess ssh.Session
		nick string
	}{sess, nick}
}

func (sm *SessionManager) Remove(id string) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.sessions, id)
}

func (sm *SessionManager) Broadcast(msg string, excludeID string) {
	sm.RLock()
	defer sm.RUnlock()
	for id, entry := range sm.sessions {
		if id != excludeID {
			fmt.Fprintln(entry.sess, msg)
		}
	}
}
