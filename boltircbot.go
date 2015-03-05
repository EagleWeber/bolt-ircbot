package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/thoj/go-ircevent"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

var ChannelUsers []string
var config = flag.String("config", "", "configuration file")

type Config struct {
	Irc struct {
		Ssl           bool     `json:"ssl"`
		SslVerifySkip bool     `json:"ssl_verify_skip"`
		Port          string   `json:"port"`
		Nickname      string   `json:"nickname"`
		Channels      []string `json:"channels"`
		Host          string   `json:"host"`
		Password      string   `json:"password"`
	} `json:"irc"`
	Github struct {
		Token string `json:"token"`
		Owner string `json:"owner"`
		Repos string `json:"repos"`
	} `json:"github"`
	Database struct {
		Karma string `json:"karma"`
	} `json:"database"`
	Logging struct {
		Location string `json:"location"`
	} `json:"logging"`
}

func (c *Config) Load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}

	if c.Irc.Nickname == "" {
		c.Irc.Nickname = "issuebot"
	}

	if c.Irc.Host == "" {
		return errors.New("host is required.")
	}

	if c.Github.Token == "" {
		return errors.New("token is required.")
	}

	if c.Github.Owner == "" {
		return errors.New("owner is required.")
	}

	if c.Github.Repos == "" {
		return errors.New("repos is required.")
	}

	return nil
}

func main() {
	flag.Parse()
	c := &Config{}
	if err := c.Load(*config); err != nil {
		log.Fatal(err)
	}

	// Logs
	logs := make(map[string]*os.File)

	ircproj := irc.IRC(c.Irc.Nickname, c.Irc.Nickname)
	ircproj.UseTLS = c.Irc.Ssl
	if c.Irc.SslVerifySkip {
		ircproj.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	ircproj.Password = c.Irc.Password

	err := ircproj.Connect(net.JoinHostPort(c.Irc.Host, c.Irc.Port))
	if err != nil {
		log.Fatal(err)
	}

	ircproj.AddCallback("001", func(event *irc.Event) {
		for _, channel := range c.Irc.Channels {
			ircproj.Join(channel)
			log.Println(fmt.Sprintf("Joined %v", channel))

			// Start the logger for this channel
			logs[channel] = StartLogger(c, channel)

			// Set the log to close on exit
			//defer logs[channel].Close()
		}
	})

	// Logging
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		channel := event.Arguments[0]
		WriteLog(c, logs[channel], event.Nick, event.Message())
	})

	r := regexp.MustCompile(`#(\d+)`)
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		matches := r.FindAllStringSubmatch(event.Message(), -1)
		for _, match := range matches {
			// Don't start a bot war
			if event.Nick == "[BoltGitHubBot]" {
				continue
			}
			if len(match) < 2 {
				continue
			}
			u, err := url.Parse(fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", c.Github.Owner, c.Github.Repos, match[1]))
			if err != nil {
				log.Println(err)
				continue
			}
			q := u.Query()
			q.Add("access_token", c.Github.Token)
			u.RawQuery = q.Encode()
			resp, err := http.Get(u.String())
			if err != nil {
				log.Println(err)
				continue
			}
			if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
				log.Println(resp.Status)
				continue
			}
			defer resp.Body.Close()
			m := make(map[string]interface{})
			if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
				log.Println(err)
				continue
			}

			if m["number"].(float64) == 1 {
				// I am a bot, I can have my own rule #1
				ircproj.Noticef(event.Arguments[0], "#1 Port Bolt to Go to keep %v happy https://github.com/bolt/bolt/issues/1", c.Irc.Nickname)
				time.Sleep(5 * time.Second)
				ircproj.Action(event.Arguments[0], "is written in Go, and therefore isn't allowed to like PHP")
			} else if m["number"].(float64) == 1555 {
				// Props to Adrian Guenter
				ircproj.Actionf(event.Arguments[0], "warns %v that #1555 nearly caused the end of the known universe and should never be mentioned again", event.Nick)
			} else {
				ircproj.Noticef(event.Arguments[0], "#%v %v %v", m["number"].(float64), m["title"].(string), m["html_url"].(string))
			}
		}
	})

	// Help
	//AddHelp(ircproj)

	// Get a list of users and remove the "@" sign for chanops
	ircproj.AddCallback("353", func(event *irc.Event) {
		s := strings.Replace(event.Arguments[3], "@", "", -1)
		ChannelUsers = strings.Fields(s)
	})

	// Just for Bopp, for now
	//ircproj.AddCallback("JOIN", func(event *irc.Event) {
	//	if event.Nick == "Bopp" {
	//		time.Sleep(5 * time.Second)
	//		ircproj.Privmsgf(event.Arguments[0], RandomMessage(), event.Nick)
	//	}
	//})

	// Asimov's Laws - Three Laws of Robotics
	AddPrivmsgRules(ircproj)
	// Documentation for bolt
	AddPrivmsgDocs(ircproj)

	AddActionf(ircproj, `#(kitten|cat)`, "starts to meow at %v… *purr* *purr*")
	AddActionf(ircproj, `#dog`, "rolls over, and wants its tummy scratched by %v")
	AddActionf(ircproj, `#champagne`, "opens a nice chilled bottle of Moët & Chandon for %v")
	AddActionf(ircproj, `#beer`, "return $this->app['beer']->serve('everyone')->sendBillTo('%v');")
	AddActionf(ircproj, `#coffee`, "turns on the espresso machine for %v")
	AddActionf(ircproj, `#hotchocolate`, "believes in miracles, %v, you sexy thing!")
	AddActionf(ircproj, `#tea`, "has boiled some water, and begins to brew %v a nice cup of tea.")
	AddActionf(ircproj, `#wine`, "opens a bottle of Château Lafite at %v's request!")
	AddActionf(ircproj, `#whisky`, "pours a nip of Glenavon Special for %v.")
	AddActionf(ircproj, `#whiskey`, "takes a swig of Jameson, hands the bottle to %v, and sings - \"Whack fol de daddy-o, There's whiskey in the jar.\"")
	AddActionf(ircproj, `#shiraz`, "wonders if %v has ever had a Heathcote Estate Shiraz?")
	AddActionf(ircproj, `#water`, "pours water over %v…  That is what they wanted, right?")
	AddActionf(ircproj, `#(PR|pr|Pr|pR)`, "gets the idea that Bopp should take care of %v's pull requests or kittens may cry…")
	AddActionf(ircproj, `#vodka`, "opens a bottle of Billionaire Vodka for %v.  It's good to be the king after all!")
	//AddActionf(ircproj, `bolt`, "calls capital_B_dangit() on %v's behalf")

	AddAction(ircproj, `#popcorn`, "yells: POPCORN! GET YOUR POPCORN!")
	AddAction(ircproj, `#pastebin`, "asks that http://pastebin.com/ be used for more than one-line messages. It makes life easier.")
	AddAction(ircproj, `#(pony|mylittlepony)`, "says \"ZA̡͊͠͝LGΌ ISͮ̂҉̯͈͕̹̘̱ TO͇̹̺ͅƝ̴ȳ̳ TH̘Ë͖́̉ ͠P̯͍̭O̚​N̐Y̡ H̸̡̪̯ͨ͊̽̅̾̎Ȩ̬̩̾͛ͪ̈́̀́͘ ̶̧̨̱̹̭̯ͧ̾ͬC̷̙̲̝͖ͭ̏ͥͮ͟Oͮ͏̮̪̝͍M̲̖͊̒ͪͩͬ̚̚͜Ȇ̴̟̟͙̞ͩ͌͝S̨̥̫͎̭ͯ̿̔̀ͅ\"")
	AddAction(ircproj, `#tequila`, "drinks one Tequila, two Tequilas, three Tequilas… floor!")
	AddActionSilentWorks(ircproj, `(WP|wp|Wordpress|WordPress|wordpress)`, "notes that if code was poetry, WordPress would have been written in Go…  It's more like \"code is pooetry if you ask this bot\"")
	AddAction(ircproj, `#nicotine`, "coughs and opens the windows…")
	AddAction(ircproj, `OCD`, "s/OCD/CDO/ …must be in alphabetical order…")

	AddActionKarma(c, ircproj)

	ircproj.Loop()
}
