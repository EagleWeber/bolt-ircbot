package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/thoj/go-ircevent"
	"log"
	"os"
	"regexp"
	"strings"
)

const dbname = "var/bolt-ircbot/karma.db"
const table = "karma"

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateDb() (*sql.DB, error) {

	// sql.Open will create a new database file if one does not exist
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Printf("Error in CreateDb()\n	%q\n", err)
	}

	// Create our table
	sqlStmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (name text PRIMARY KEY, score integer);", table)
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	} else {
		log.Printf("Creating table:\n	%s\n", sqlStmt)
	}

	return db, nil
}

func GetKarma(name string) (int, error) {

	var score int

	db, err := getDb()
	defer db.Close()

	if err != nil {
		return 0, err
	}

	sqlStmt := fmt.Sprintf("SELECT score from %v WHERE name = '%v' LIMIT 1", table, name)

	rows := db.QueryRow(sqlStmt).Scan(&score)

	if rows == sql.ErrNoRows {
		log.Printf("GetKarma() no user with that ID.")
		return score, rows
	} else if rows != nil {
		log.Printf(fmt.Sprintf("GetKarma query failed:\n	%s\n	%s", rows, sqlStmt))
		return score, rows
	}

	return score, nil
}

func AddKarma(name string) (int, error) {

	var score int
	var err error

	db, err := getDb()
	defer db.Close()

	if err != nil {
		return 0, nil
	}

	// Get the current score, if not found, a value of 0 is returned
	sqlStmt := fmt.Sprintf("SELECT score from %v WHERE name = '%v' LIMIT 1", table, name)
	rows := db.QueryRow(sqlStmt).Scan(&score)
	if rows != nil {
		log.Println(fmt.Sprintf("AddKarma() query failed: %q - %s\n", rows, sqlStmt))
	}

	// Increment the score value
	score++

	// Insert or update the users record
	if rows == sql.ErrNoRows {
		sqlStmt := fmt.Sprintf("INSERT INTO %v (name, score) VALUES ('%v', '%d');", table, name, score)
		//log.Println(fmt.Sprintf("INSERT INTO %v (name, score) VALUES ('%v', '%d');", table, name, score))
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Println(fmt.Sprintf("INSERT failed: %s", err))
		}
	} else {
		sqlStmt := fmt.Sprintf("UPDATE %v SET score = '%d' WHERE name = '%s';", table, score, name)
		//log.Println(fmt.Sprintf("UPDATE %v SET score = '%d' WHERE name = '%s';", table, score, name))
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Println(fmt.Sprintf("UPDATE failed: %s", err))
		}
	}

	return score, nil
}

func getDb() (*sql.DB, error) {
	var db *sql.DB
	var err error
	//var err error

	if Exists(dbname) {
		db, err = sql.Open("sqlite3", dbname)
	} else {
		db, err = CreateDb()
	}

	if err != nil {
		log.Println(err)
		return db, err
	}

	return db, nil
}

func AddActionKarma(ircproj *irc.Connection) error {

	hash := `#karma`
	
	x := regexp.MustCompile(hash)
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {

		matches := x.FindAllStringSubmatch(event.Message(), -1)
		if len(matches) > 0 {
			msg := strings.Trim(event.Arguments[1], " ")
			tokens := strings.Split(msg, " ")
			
			// TODO Get the list of users in channel now
			//ircproj.SendRawf("NAMES %v", event.Arguments[0])

			for _, element := range tokens {
				if strings.HasPrefix(element, "#") {
					continue
				}
			
				// Catch some using their own name
				if event.Nick == element {
					karma, err := GetKarma(event.Nick)
					
					if err != nil {
						ircproj.Privmsgf(event.Arguments[0], "BoltKarma for %s is currently zero", event.Nick)
					} else {
						ircproj.Privmsgf(event.Arguments[0], "BoltKarma for %s is currently %d", event.Nick, karma)
					}
				} else {
					karma, err := AddKarma(element)
					
					if err != nil {
						// log an error
						log.Println(fmt.Sprintf("Ooopsy %s", err))
					} else {
						ircproj.Privmsgf(event.Arguments[0], "BoltKarma for %s is now %d", element, karma)
					}
				}
			}
		}
	})
	
	ircproj.AddCallback("NAMES", func(event *irc.Event) {
		log.Println("Got some names")
		log.Println(event.Arguments[1])
	})

	return nil
}