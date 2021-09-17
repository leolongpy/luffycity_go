package ginsession

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	SessionCookieName  = "session_id"
	SessionContextName = "session"
)

var (
	MgrObj Mgr
)

type Option struct {
	MaxAge   int
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
}

type SessionData interface {
	GetID() string
	Get(key string) (value interface{}, err error)
	Set(key string, value interface{})
	Del(key string)
	Save()
	SetExpire(int)
	Load(sessionID string) (err error)
}

type Mgr interface {
	Init(addr string, options ...string)
	GetSessionData(sessionId string) (sd SessionData, err error)
	CreateSession() (sd SessionData)
}

func InitMgr(name string, addr string, option ...string) {
	switch name {
	case "memory":
		MgrObj = NewMemoryMgr()
	case "redis":
		MgrObj = NewRedisMgr()
	}
	MgrObj.Init(addr, option...)
}

// 实现一个gin框架的中间件
// 所有流经我这个中间件的请求，它的上下文中肯定会有一个session -> session data
func SessionMiddleware(mgrObj Mgr, option *Option) gin.HandlerFunc {
	if mgrObj == nil {
		panic("must call InitMgr before use it.")
	}
	return func(c *gin.Context) {
		// 1. 从请求的Cookie中获取session_id
		var sd SessionData // session data
		sessionID, err := c.Cookie(SessionCookieName)
		fmt.Println(sessionID)
		if err != nil {
			// 1.1 取不到session_id -> 给这个新用户创建一个新的session data，同时分配一个session_id
			sd = mgrObj.CreateSession()
			sessionID = sd.GetID()
			fmt.Println("取不到session_id，创建一个新的", sessionID)
		} else {
			// 1.2 取到session_id
			// 2. 根据session_id去Session大仓库中取到对应的session data
			sd, err = mgrObj.GetSessionData(sessionID)
			if err != nil {
				// 2.1 根据用户传过来的session_id在大仓库中根本取不到session data
				sd = mgrObj.CreateSession() // sd
				// 2.2 更新用户Cookie中保存的那个session_id
				sessionID = sd.GetID()
				fmt.Println("session_id取不到session data,分配一个新的", sessionID)
			}
			fmt.Println("session_id未过期", sessionID)
		}
		sd.SetExpire(option.MaxAge) // 设置session data过期时间
		// 3. 如何实现让后续所有的处理请求的方法都能拿到session data
		// 3. 利用gin的c.Set("session", session data)
		c.Set(SessionContextName, sd) // 保存到上下文
		// 在gin框架中，要回写Cookie必须在处理请求的函数返回之前
		c.SetCookie(SessionCookieName, sessionID, option.MaxAge, option.Path, option.Domain, option.Secure, option.HttpOnly)
		c.Next() // 执行后续的请求处理方法 c.HTML()时已经把响应头写好了
	}
}
