package session

import "fmt"

// Get 根据key获取值
func (s *SessionData) Get(key string) (value interface{}, err error) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	value, ok := s.Data[key]
	if !ok {
		err = fmt.Errorf("invalid key")
		return
	}
	return
}

// Set 创建数据
func (s *SessionData) Set(key string, value interface{}) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.Data[key] = value
}

// Del 删除key
func (s *SessionData) Del(key string) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	delete(s.Data, key)
}
