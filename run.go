package main

import (
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/plugins/wxweb/config"
	"github.com/songtianyi/wechat-go/plugins/wxweb/switcher"
	"github.com/songtianyi/wechat-go/plugins/wxweb/system"
	"github.com/songtianyi/wechat-go/wxweb"
	"time"
	"wechat-super-go/src/plugins/demo"
	"wechat-super-go/src/plugins/love"
	"wechat-super-go/src/plugins/revoker"
)

func main() {
	logs.Info("start bot...")
	// create session
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}

	logs.Info("session:数据:",*session)
	system.Register(session)
	switcher.Register(session)
	config.Register(session)
	// load plugins for this session
	demo.Register(session)
	//管理群
	love.Register(session)
	revoker.Register(session)

	// disable by type example
	if err := session.HandlerRegister.DisableByType(wxweb.MSG_SYS); err != nil {
		logs.Error(err)
		return
	}
	if err := session.HandlerRegister.DisableByType(wxweb.MSG_TEXT); err != nil {
		logs.Error(err)
		return
	}
	if err := session.HandlerRegister.DisableByType(wxweb.MSG_IMG); err != nil {
		logs.Error(err)
		return
	}

	session.HandlerRegister.EnableByName("switcher")
	session.HandlerRegister.EnableByName("love")
	session.HandlerRegister.EnableByName("revoker")

	for {
		if err := session.LoginAndServe(false); err != nil {
			logs.Error("session exit, %s", err)
			for i := 0; i < 3; i++ {
				logs.Info("trying re-login with cache")
				if err := session.LoginAndServe(true); err != nil {
					logs.Error("re-login error, %s", err)
				}
				time.Sleep(3 * time.Second)
			}
			if session, err = wxweb.CreateSession(nil, session.HandlerRegister, wxweb.TERMINAL_MODE); err != nil {
				logs.Error("create new sesion failed, %s", err)
				break
			}
		} else {
			logs.Info("closed by user")
			break
		}
	}
}
