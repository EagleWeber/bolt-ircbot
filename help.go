package main

import (
	"github.com/thoj/go-ircevent"
	"regexp"
	"time"
)

func helpMsg(ircproj *irc.Connection, event *irc.Event, msg string) error {
	ircproj.Privmsg(event.Nick, msg)
	time.Sleep(100 * time.Millisecond)

	return nil
}

func AddHelp(ircproj *irc.Connection) error {
	x := regexp.MustCompile(`#help`)
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		matches := x.FindAllStringSubmatch(event.Message(), -1)
		if len(matches) > 0 {
			helpMsg(ircproj, event, "I can give you helpful links if you use the following # commands in channel:")
			helpMsg(ircproj, event, "\t#about")
			helpMsg(ircproj, event, "\t#cheatsheet")
			helpMsg(ircproj, event, "\t#codequality")
			helpMsg(ircproj, event, "\t#contenttypes")
			helpMsg(ircproj, event, "\t#contribute")
			helpMsg(ircproj, event, "\t#credit")
			helpMsg(ircproj, event, "\t#docs")
			helpMsg(ircproj, event, "\t#extend")
			helpMsg(ircproj, event, "\t#extensions")
			helpMsg(ircproj, event, "\t#github")
			helpMsg(ircproj, event, "\t#install")
			helpMsg(ircproj, event, "\t#internals")
			helpMsg(ircproj, event, "\t#introduction")
			helpMsg(ircproj, event, "\t#issue")
			helpMsg(ircproj, event, "\t#maintenance")
			helpMsg(ircproj, event, "\t#menu")
			helpMsg(ircproj, event, "\t#nut")
			helpMsg(ircproj, event, "\t#paging")
			helpMsg(ircproj, event, "\t#permissions")
			helpMsg(ircproj, event, "\t#record")
			helpMsg(ircproj, event, "\t#relationship")
			helpMsg(ircproj, event, "\t#requirements")
			helpMsg(ircproj, event, "\t#resources")
			helpMsg(ircproj, event, "\t#roadmap")
			helpMsg(ircproj, event, "\t#routes")
			helpMsg(ircproj, event, "\t#rules")
			helpMsg(ircproj, event, "\t#screenshot")
			helpMsg(ircproj, event, "\t#search")
			helpMsg(ircproj, event, "\t#tags")
			helpMsg(ircproj, event, "\t#taxonomies")
			helpMsg(ircproj, event, "\t#taxonomy")
			helpMsg(ircproj, event, "\t#template")
			helpMsg(ircproj, event, "\t#templatetags")
			helpMsg(ircproj, event, "\t#themes")
			helpMsg(ircproj, event, "\t#updates")
			helpMsg(ircproj, event, "\t#updating")
		}
	})

	return nil
}
