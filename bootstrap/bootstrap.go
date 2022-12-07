package bootstrap

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/wechatgpt/wechatbot/handler"
	"log"
)

func Run() {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.MessageHandler = handler.Handler
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	reloadStorage := openwechat.NewJsonFileHotReloadStorage("token.json")
	err := bot.HotLogin(reloadStorage)
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Fatal(err)
			return
		}
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Fatal(err)
		return
	}

	friends, err := self.Friends()
	log.Println(friends, err)
	groups, err := self.Groups()
	log.Println(groups, err)

	err = bot.Block()
	if err != nil {
		log.Fatal(err)
		return
	}

}