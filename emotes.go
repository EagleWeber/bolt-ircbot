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

func AddPrivmsgDocs(ircproj *irc.Connection) error {
	cheatsheet := regexp.MustCompile(`#cheatsheet`)
	docs := regexp.MustCompile(`#docs`)
	github := regexp.MustCompile(`#github`)
	install := regexp.MustCompile(`#install`)
	routes := regexp.MustCompile(`#(routes|routing)`)
	contenttypes := regexp.MustCompile(`#contenttype`)
	extend := regexp.MustCompile(`#(extend|themes)`)
	extensions := regexp.MustCompile(`#extensions`)
	permissions := regexp.MustCompile(`#permissions`)
	requirements := regexp.MustCompile(`#requirements`)
	updating := regexp.MustCompile(`#(updates|updating)`)
	screenshots := regexp.MustCompile(`#screenshot`)
	taxonomies := regexp.MustCompile(`#(taxonomy|taxonomies)`)
	menus := regexp.MustCompile(`#menu`)
	relationships := regexp.MustCompile(`#relationship`)
	templates := regexp.MustCompile(`#template`)
	records := regexp.MustCompile(`#record`)
	paging := regexp.MustCompile(`#(page|paging)`)
	search := regexp.MustCompile(`#search`)
	temptags := regexp.MustCompile(`#(tags|templatetags)`)
	internals := regexp.MustCompile(`#internal`)
	nut := regexp.MustCompile(`#nut`)
	contributing := regexp.MustCompile(`#(contribute|contributing)`)
	maintenance := regexp.MustCompile(`#maintenance`)
	roadmap := regexp.MustCompile(`#roadmap`)
	resources := regexp.MustCompile(`#resources`)
	about := regexp.MustCompile(`#(about|introduction)`)
	codequality := regexp.MustCompile(`#(quality|codequality)`)
	credits := regexp.MustCompile(`#credits`)
	issues := regexp.MustCompile(`#issue`)

	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		if len(cheatsheet.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "A cheatsheet for bolt is available at https://cheatsheet.bolt.cm")
		}
		if len(docs.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt documentation is available at https://docs.bolt.cm")
		}
		if len(github.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt source code is available at https://github.com/bolt/bolt")
		}
		if len(install.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Full documentation on installing bolt is available at https://docs.bolt.cm/installation")
		}
		if len(routes.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Routes documentation is available at https://docs.bolt.cm/templates-routes")
		}
		if len(contenttypes.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Templates and Routes documentation is available at https://docs.bolt.cm/content")
		}
		if len(extend.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Extensions and themes are available at https://extensions.bolt.cm")
		}
		if len(permissions.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Permissions documentation is available at https://docs.bolt.cm/permissions")
		}
		if len(requirements.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt requirements are listed at https://docs.bolt.cm/requirements")
		}
		if len(updating.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Information on updates and updating is at https://docs.bolt.cm/updating")
		}
		if len(screenshots.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Screenshots of bolt are available at https://docs.bolt.cm/screenshots")
		}
		if len(taxonomies.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Taxonomies documentation is available at https://docs.bolt.cm/taxonomies")
		}
		if len(menus.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Menus documentation is available at https://docs.bolt.cm/menus")
		}
		if len(relationships.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Relationships documentation is available at https://docs.bolt.cm/relationships")
		}
		if len(templates.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Templates documentation is available at https://docs.bolt.cm/templates")
		}
		if len(records.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Record and Records documentation is available at https://docs.bolt.cm/record-and-records")
		}
		if len(paging.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Pagers and Paging documentation is available at https://docs.bolt.cm/content-paging")
		}
		if len(search.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Search documentation is available at https://docs.bolt.cm/content-search")
		}
		if len(temptags.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Template Tags information is available at https://docs.bolt.cm/templatetags")
		}
		if len(extensions.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Extensions documentation is available at https://docs.bolt.cm/extensions")
		}
		if len(internals.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Documentation on bolts internals is available at https://docs.bolt.cm/internals")
		}
		if len(nut.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Documentation on Nut is available at https://docs.bolt.cm/nut")
		}
		if len(contributing.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Information on contributing to bolt is available at https://docs.bolt.cm/contributing")
		}
		if len(maintenance.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Information on maintenance mode is available at https://docs.bolt.cm/maintenancemode")
		}
		if len(roadmap.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "A roadmap for bolt is available at https://docs.bolt.cm/roadmap")
		}
		if len(resources.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "A list of bolt resources is available at https://docs.bolt.cm/resources")
		}
		if len(about.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Read an introduction to bolt https://docs.bolt.cm/resources")
		}
		if len(codequality.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Code quality guidelines are available at https://docs.bolt.cm/code-quality")
		}
		if len(credits.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Credits for code used in bolt are available at https://docs.bolt.cm/credits")
		}
		if len(issues.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt issue tracker available at https://github.com/bolt/bolt/issues")
		}

	})

	return nil
}
