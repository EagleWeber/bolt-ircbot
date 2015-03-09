package main

import (
	"github.com/thoj/go-ircevent"
	"time"
)

var LastTobiasComment = time.Now()
var TobiasCommentCount = 0

func TobiasCommentTimeTrack() time.Duration {
	elapsed := time.Since(LastTobiasComment)

	return elapsed
}

func AddTobias(ircproj *irc.Connection) error {
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		if event.Nick == "tdammers" {
			elapsed := TobiasCommentTimeTrack()

			if elapsed < 30*time.Second {
				TobiasCommentCount++
			} else {
				TobiasCommentCount = 1
			}

			if TobiasCommentCount > 5 {
				ircproj.Notice(event.Arguments[0], "~~~ WARNING! TOBIAS RANT DETECTED! ~~~")
				TobiasCommentCount = 0
			}

			LastTobiasComment = time.Now()
		}
	})

	return nil
}
