package love

import (
	"fmt"
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/wxweb"
)

func Register(session *wxweb.Session)  {
	session.HandlerRegister.Add(wxweb.MSG_TEXT,wxweb.Handler(manageLove),"love")
	if err := session.HandlerRegister.EnableByName("love"); err != nil {
		logs.Error(err)
	}
}

/**
管理群
 */
func manageLove(session *wxweb.Session, msg *wxweb.ReceivedMessage)  {

	logs.Warn("msg.IsGroup::",msg.IsGroup)
	logs.Warn("session.Bot::",session.Bot)
	logs.Warn("msg::",msg)
	for _, v := range session.Cm.GetGroupContacts() {
		fmt.Println(v.NickName)
		logs.Warn(v.NickName)
	}
	//session.SendText(msg.Content, session.Bot.UserName, wxweb.RealTargetUserName(session, msg))
}
