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

		if elapsed < 300*time.Second {
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
	about := regexp.MustCompile(`#(about|introduction)`)
	cheatsheet := regexp.MustCompile(`#cheatsheet`)
	codequality := regexp.MustCompile(`#(quality|codequality)`)
	contenttypes := regexp.MustCompile(`#contenttype`)
	contributing := regexp.MustCompile(`#(contribute|contributing)`)
	credits := regexp.MustCompile(`#credits`)
	docs := regexp.MustCompile(`#docs`)
	extend := regexp.MustCompile(`#(extend|themes)`)
	extensions := regexp.MustCompile(`#extensions`)
	github := regexp.MustCompile(`#github`)
	htaccess_override := regexp.MustCompile(`\/users\/edit\/ was not found on this server`)
	htaccess := regexp.MustCompile(`#(rewrite|htaccess|apache)`)
	install := regexp.MustCompile(`#install`)
	internals := regexp.MustCompile(`#internal`)
	issues := regexp.MustCompile(`#issue`)
	maintenance := regexp.MustCompile(`#maintenance`)
	manifesto := regexp.MustCompile(`#manifesto`)
	menus := regexp.MustCompile(`#menu`)
	nest := regexp.MustCompile(`#nest`)
	nut := regexp.MustCompile(`#nut`)
	paging := regexp.MustCompile(`#(page|paging)`)
	permissions := regexp.MustCompile(`#permissions`)
	records := regexp.MustCompile(`#record`)
	relationships := regexp.MustCompile(`#relationship`)
	requirements := regexp.MustCompile(`#requirements`)
	resources := regexp.MustCompile(`#resources`)
	roadmap := regexp.MustCompile(`#roadmap`)
	routes := regexp.MustCompile(`#(routes|routing)`)
	screenshots := regexp.MustCompile(`#screenshot`)
	search := regexp.MustCompile(`#search`)
	taxonomies := regexp.MustCompile(`#(taxonomy|taxonomies)`)
	templates := regexp.MustCompile(`#template`)
	temptags := regexp.MustCompile(`#(tags|templatetags)`)
	updating := regexp.MustCompile(`#(updates|updating)`)
	webroot := regexp.MustCompile(`#webroot`)
	workmap := regexp.MustCompile(`#workmap`)

	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		if len(about.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Read an introduction to Bolt https://docs.bolt.cm/getting-started/about")
		}
		if len(cheatsheet.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "A cheatsheet for Bolt is available at https://docs.bolt.cm/cheatsheet")
		}
		if len(codequality.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt's Code Quality Guidelines are available at https://docs.bolt.cm/other/code-quality")
		}
		if len(contenttypes.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "ContentTypes documentation is available at https://docs.bolt.cm/contenttypes")
		}
		if len(contributing.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Information on contributing to Bolt is available at https://docs.bolt.cm/other/contributing")
		}
		if len(credits.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Credits for code used in Bolt are available at https://docs.bolt.cm/other/credits")
		}
		if len(docs.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt documentation is available at https://docs.bolt.cm")
		}
		if len(extend.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Extensions and themes are available at https://extensions.bolt.cm")
		}
		if len(extensions.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Extensions documentation is available at https://docs.bolt.cm/extensions")
		}
		if len(github.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt source code is available at https://github.com/bolt/bolt")
		}
		if len(htaccess.FindAllStringSubmatch(event.Message(), -1)) > 0 || len(htaccess_override.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Having Apache rewrite issues? Have a look at https://docs.bolt.cm/installation/webserver/apache")
		}
		if len(install.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Full documentation on installing Bolt is available at https://docs.bolt.cm/installation/installation")
		}
		if len(internals.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Documentation on Bolt's internals is available at https://docs.bolt.cm/internals")
		}
		if len(issues.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt issue tracker available at https://github.com/bolt/bolt/issues")
		}
		if len(maintenance.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Information on maintenance mode is available at https://docs.bolt.cm/other/maintenance-mode")
		}
		if len(manifesto.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt manifesto available at https://docs.bolt.cm/other/manifesto")
		}
		if len(menus.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Menus documentation is available at https://docs.bolt.cm/configuration/menus")
		}
		if len(nest.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Maximum function nesting level of '100' reached? Set xdebug.max_nesting_level=1000 in your php.ini file. For more information see https://adayinthelifeof.nl/2015/11/17/symfony-xdebug-and-maximum-nesting-level-issues/")
		}
		if len(nut.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Documentation on Nut is available at https://docs.bolt.cm/other/nut")
		}
		if len(paging.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Pagers and Paging documentation is available at https://docs.bolt.cm/templating/content-paging")
		}
		if len(permissions.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Permissions documentation is available at https://docs.bolt.cm/configuration/permissions")
		}
		if len(records.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Record and Records documentation is available at https://docs.bolt.cm/templating/record-and-records")
		}
		if len(relationships.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Relationships documentation is available at https://docs.bolt.cm/contenttypes/relationships")
		}
		if len(requirements.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt requirements are listed at https://docs.bolt.cm/getting-started/requirements")
		}
		if len(resources.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "A list of Bolt resources is available at https://docs.bolt.cm/other")
		}
		if len(roadmap.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Bolt 2.x Roadmap available at https://docs.bolt.cm/3.0/other/roadmap")
		}
		if len(routes.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Routes documentation is available at https://docs.bolt.cm/templating/templates-routes")
		}
		if len(screenshots.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Screenshots of Bolt are available at https://docs.bolt.cm/getting-started/screenshots")
		}
		if len(search.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Search documentation is available at https://docs.bolt.cm/templating/content-search")
		}
		if len(taxonomies.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Taxonomies documentation is available at https://docs.bolt.cm/contenttypes/taxonomies")
		}
		if len(templates.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Templates documentation is available at https://docs.bolt.cm/templating/building-templates")
		}
		if len(temptags.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Template Tags information is available at https://docs.bolt.cm/templating/templatetags")
		}
		if len(updating.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Information on updates and updating is at https://docs.bolt.cm/upgrading/updating")
		}
		if len(webroot.FindAllStringSubmatch(event.Message(), -1)) > 0 || len(htaccess_override.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Having trouble with using Bolt outside of the web root. Have a look at https://docs.bolt.cm/3.0/howto/troubleshooting-outside-webroot")
		}
		if len(workmap.FindAllStringSubmatch(event.Message(), -1)) > 0 {
			ircproj.Privmsg(event.Arguments[0], "Core Development Workmap: https://github.com/bolt/bolt/wiki/Bolt-Core-Development-Workmap")
		}
	})

	return nil
}
