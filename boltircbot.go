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
		}
	})

	r := regexp.MustCompile(`#(\d+)`)
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		matches := r.FindAllStringSubmatch(event.Message(), -1)
		for _, match := range matches {
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
			    ircproj.Notice(event.Arguments[0], "#1 Port Bolt to Go to keep the bot happy https://github.com/bolt/bolt/issues/1")
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
	
	// Get a list of users and remove the "@" sign for chanops
	ircproj.AddCallback("353", func(event *irc.Event) {
	        s := strings.Replace(event.Arguments[3], "@", "", -1)
	        ChannelUsers = strings.Fields(s)
	})

	// Just for Bopp, for now
	ircproj.AddCallback("JOIN", func(event *irc.Event) {
		if event.Nick == "Bopp" {
			ircproj.Privmsgf(event.Arguments[0], RandomMessage(), event.Nick)
		}
	})

	AddActionf(ircproj, `#(kitten|cat)`,  "starts to meow at %v... *purr* *purr*")
	AddActionf(ircproj, `#dog`,  "rolls over, and wants its tummy scratched by %v")
	AddActionf(ircproj, `#beer`,  "return $this->app['beer']->serve('everyone')->sendBillTo('%v');")
	AddActionf(ircproj, `#coffee`,  "turns on the espresso machine for %v")
	AddActionf(ircproj, `#hotchocolate`,  "believes in miracles, %v, you sexy thing!")
	AddActionf(ircproj, `#tea`,  "has boiled some water, and begins to brew %v a nice cup of tea.")
	AddActionf(ircproj, `#wine`,  "opens a bottle of Château Lafite at %v's request!")
	AddActionf(ircproj, `#whisky`,  "pours a nip of Glenavon Special for %v.")
	AddActionf(ircproj, `#whiskey`,  "takes a swig of Jameson, hands the bottle to %v, and sings - \"Whack fol de daddy-o, There's whiskey in the jar.\"")
	AddActionf(ircproj, `#shiraz`,  "wonders if %v has ever had a Heathcote Estate Shiraz?")
	
	AddAction(ircproj, `#tequila`,  "drinks one Tequila, two Tequilas, three Tequilas... floor!")
	
	AddActionKarma(c, ircproj)

	ircproj.Loop()
}
