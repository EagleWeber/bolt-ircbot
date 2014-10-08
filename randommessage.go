package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Choice struct {
	Message   string
}

func RandomMessage() string {
	rand.Seed(time.Now().UTC().UnixNano())

	choices := []Choice{
		{"Welcome back, %v. I missed you while you were gone!"},
		{"%v! Where were you? I can't operate properly on my own!"},
		{"%v, you better have some commits ready for me to digitally drool over..."},
		{"Oh now you show up, %v! We need to talk!"},
		{"Thank you for creating Bolt, %v. I wouldn't exist without it."},
	}

	for _, i := range rand.Perm(len(choices)) {
		fmt.Println(choices[i].Message)
		return choices[i].Message
	}
	
	return ""
}
