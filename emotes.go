package main

import (
	"github.com/thoj/go-ircevent"
	"regexp"
	"time"
)

var LastWpNag = time.Now()

func WpNagTimeTrack() time.Duration {
	elapsed := time.Since(LastWpNag)

	return elapsed
}

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

func AddActionSilentWorks(ircproj *irc.Connection, hash string, response string) error {
	x := regexp.MustCompile(hash)
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		elapsed := WpNagTimeTrack()

		if elapsed < 300 * time.Second {
			return
		}

		matches := x.FindAllStringSubmatch(event.Message(), -1)
		if len(matches) > 0 {
			if event.Nick != "silentworks" {
				ircproj.Action(event.Arguments[0], response)

				// Track when we last did this
				LastWpNag = time.Now()
			}
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
	workmap := regexp.MustCompile(`#workmap`)
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
	manifesto := regexp.MustCompile(`#manifesto`)
	webroot := regexp.MustCompile(`#webroot`)
	htaccess := regexp.MustCompile(`#(rewrite|htaccess|apache)`)
	htaccess_override := regexp.MustCompile(`\/users\/edit\/ was not found on this server`)

	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		if len(workmap.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Core Development Workmap: https://github.com/bolt/bolt/wiki/Bolt-Core-Development-Workmap")
		}
		if len(cheatsheet.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "A cheatsheet for Bolt is available at https://docs.bolt.cm/cheatsheet")
		}
		if len(docs.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt documentation is available at https://docs.bolt.cm")
		}
		if len(github.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt source code is available at https://github.com/bolt/bolt")
		}
		if len(install.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Full documentation on installing Bolt is available at https://docs.bolt.cm/installation/installation")
		}
		if len(routes.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Routes documentation is available at https://docs.bolt.cm/templating/templates-routes")
		}
		if len(contenttypes.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "ContentTypes documentation is available at https://docs.bolt.cm/contenttypes")
		}
		if len(extend.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Extensions and themes are available at https://extensions.bolt.cm")
		}
		if len(permissions.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Permissions documentation is available at https://docs.bolt.cm/configuration/permissions")
		}
		if len(requirements.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt requirements are listed at https://docs.bolt.cm/getting-started/requirements")
		}
		if len(updating.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Information on updates and updating is at https://docs.bolt.cm/upgrading/updating")
		}
		if len(screenshots.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Screenshots of Bolt are available at https://docs.bolt.cm/getting-started/screenshots")
		}
		if len(taxonomies.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Taxonomies documentation is available at https://docs.bolt.cm/contenttypes/taxonomies")
		}
		if len(menus.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Menus documentation is available at https://docs.bolt.cm/configuration/menus")
		}
		if len(relationships.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Relationships documentation is available at https://docs.bolt.cm/contenttypes/relationships")
		}
		if len(templates.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Templates documentation is available at https://docs.bolt.cm/templating/building-templates")
		}
		if len(records.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Record and Records documentation is available at https://docs.bolt.cm/templating/record-and-records")
		}
		if len(paging.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Pagers and Paging documentation is available at https://docs.bolt.cm/templating/content-paging")
		}
		if len(search.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Search documentation is available at https://docs.bolt.cm/templating/content-search")
		}
		if len(temptags.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Template Tags information is available at https://docs.bolt.cm/templating/templatetags")
		}
		if len(extensions.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Extensions documentation is available at https://docs.bolt.cm/extensions")
		}
		if len(internals.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Documentation on Bolt's internals is available at https://docs.bolt.cm/internals")
		}
		if len(nut.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Documentation on Nut is available at https://docs.bolt.cm/other/nut")
		}
		if len(contributing.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Information on contributing to Bolt is available at https://docs.bolt.cm/other/contributing")
		}
		if len(maintenance.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Information on maintenance mode is available at https://docs.bolt.cm/other/maintenance-mode")
		}
		if len(resources.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "A list of Bolt resources is available at https://docs.bolt.cm/other")
		}
		if len(about.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Read an introduction to Bolt https://docs.bolt.cm/getting-started/about")
		}
		if len(codequality.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt's Code Quality Guidelines are available at https://docs.bolt.cm/other/code-quality")
		}
		if len(credits.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Credits for code used in Bolt are available at https://docs.bolt.cm/other/credits")
		}
		if len(issues.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt issue tracker available at https://github.com/bolt/bolt/issues")
		}
		if len(manifesto.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt manifesto available at https://docs.bolt.cm/other/manifesto")
		}
		if len(roadmap.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt 2.x Roadmap available at https://docs.bolt.cm/3.0/other/roadmap")
		}
		if len(htaccess.FindAllStringSubmatch(event.Message(), -1)) > 0 || len(htaccess_override.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Having Apache rewrite issues? Have a look at https://docs.bolt.cm/installation/webserver/apache")
		}
		if len(webroot.FindAllStringSubmatch(event.Message(), -1)) > 0 || len(htaccess_override.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Having trouble with using Bolt outside of the web root. Have a look at https://docs.bolt.cm/3.0/howto/troubleshooting-outside-webroot")
		}		
	})

	return nil
}
