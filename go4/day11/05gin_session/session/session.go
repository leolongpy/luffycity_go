package session

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// SessionData 表示一个具体的用户Session数据
type SessionData struct {
	ID     string
	Data   map[string]interface{}
	rwLock sync.RWMutex
}

// NewSessionData 构造函数
func NewSessionData(id string) *SessionData {
	return &SessionData{
		ID:   id,
		Data: make(map[string]interface{}, 8),
	}
}

type Mgr struct {
	Session map[string]SessionData
	rwLock  sync.RWMutex
}

// GetSessionData 根据传进来的SessionId找到对应的SessionData
func (m *Mgr) GetSessionData(sessionId string) (sd SessionData, err error) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	sd, ok := m.Session[sessionId]
	if !ok {
		err = fmt.Errorf("invalid session id")
		return
	}
	return
}

func (m *Mgr) CreateSession() (sd *SessionData, err error) {
	//造一个sessionId
	timeStamp := time.Now().UnixNano()
	if err != nil {
		return
	}
	sd = NewSessionData(strconv.FormatInt(timeStamp, 10))
	return
}
