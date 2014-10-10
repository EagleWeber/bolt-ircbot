package main

import (
	"github.com/thoj/go-ircevent"
	"regexp"
)

func AddAction(ircproj *irc.Connection, hash string, response string) error {
	x := regexp.MustCompile(hash)
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		matches := x.FindAllStringSubmatch(event.Message(), -1)
		if len(matches) > 0 {
			ircproj.Action(event.Arguments[0], response)
		}
	})

	return nil
}

func AddActionf(ircproj *irc.Connection, hash string, response string) error {
	x := regexp.MustCompile(hash)
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		matches := x.FindAllStringSubmatch(event.Message(), -1)
		if len(matches) > 0 {
			ircproj.Actionf(event.Arguments[0], response, event.Nick)
		}
	})

	return nil
}

func AddPrivmsgRules(ircproj *irc.Connection) error {
    x := regexp.MustCompile(`#rules`)
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		matches := x.FindAllStringSubmatch(event.Message(), -1)
		if len(matches) > 0 {
			ircproj.Privmsg(event.Arguments[0], "1. A robot may not injure a human being or, through inaction, allow a human being to come to harm.")
			ircproj.Privmsg(event.Arguments[0], "2. A robot must obey the orders given to it by human beings, except where such orders would conflict with the First Law.")
			ircproj.Privmsg(event.Arguments[0], "3. A robot must protect its own existence as long as such protection does not conflict with the First or Second Law.")
		}
	})

	return nil
}