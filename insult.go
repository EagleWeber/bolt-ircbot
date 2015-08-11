package main

import (
	"math/rand"
	"time"
	"github.com/thoj/go-ircevent"
	"regexp"
	"strings"
)


var insults1 = [...]Choice {{"artless"},{"bawdy"},{"beslubbering"},{"bootless"},{"churlish"},{"cockered"},{"clouted"},{"craven"},{"currish"},{"dankish"},{"dissembling"},{"droning"},{"errant"},{"fawning"},{"fobbing"},{"froward"},{"frothy"},{"gleeking"},{"goatish"},{"gorbellied"},{"impertinent"},{"infectious"},{"jarring"},{"loggerheaded"},{"lumpish"},{"mammering"},{"mangled"},{"mewling"},{"paunchy"},{"pribbling"},{"puking"},{"puny"},{"qualling"},{"rank"},{"reeky"},{"roguish"},{"ruttish"},{"saucy"},{"spleeny"},{"spongy"},{"surly"},{"tottering"},{"unmuzzled"},{"vain"},{"venomed"},{"villainous"},{"warped"},{"wayward"},{"weedy"},{"yeasty"},{"cullionly"},{"fusty"},{"caluminous"},{"wimpled"},{"burly-boned"},{"misbegotten"},{"odiferous"},{"poisonous"},{"fishified"},{"Wart-necked"}}
var insults2 = [...]Choice {{"base-court"},{"bat-fowling"},{"beef-witted"},{"beetle-headed"},{"boil-brained"},{"clapper-clawed"},{"clay-brained"},{"common-kissing"},{"crook-pated"},{"dismal-dreaming"},{"dizzy-eyed"},{"doghearted"},{"dread-bolted"},{"earth-vexing"},{"elf-skinned"},{"fat-kidneyed"},{"fen-sucked"},{"flap-mouthed"},{"fly-bitten"},{"folly-fallen"},{"fool-born"},{"full-gorged"},{"guts-griping"},{"half-faced"},{"hasty-witted"},{"hedge-born"},{"hell-hated"},{"idle-headed"},{"ill-breeding"},{"ill-nurtured"},{"knotty-pated"},{"milk-livered"},{"motley-minded"},{"onion-eyed"},{"plume-plucked"},{"pottle-deep"},{"pox-marked"},{"reeling-ripe"},{"rough-hewn"},{"rude-growing"},{"rump-fed"},{"shard-borne"},{"sheep-biting"},{"spur-galled"},{"swag-bellied"},{"tardy-gaited"},{"tickle-brained"},{"toad-spotted"},{"unchin-snouted"},{"weather-bitten"},{"whoreson"},{"malmsey-nosed"},{"rampallian"},{"lily-livered"},{"scurvy-valiant"},{"brazen-faced"},{"unwash'd"},{"bunch-back'd"},{"leaden-footed"},{"muddy-mettled"},{"pigeon-liver'd"},{"scale-sided"}}
var insults3 = [...]Choice {{"apple-john"},{"baggage"},{"barnacle"},{"bladder"},{"boar-pig"},{"bugbear"},{"bum-bailey"},{"canker-blossom"},{"clack-dish"},{"clotpole"},{"coxcomb"},{"codpiece"},{"death-token"},{"dewberry"},{"flap-dragon"},{"flax-wench"},{"flirt-gill"},{"foot-licker"},{"fustilarian"},{"giglet"},{"gudgeon"},{"haggard"},{"harpy"},{"hedge-pig"},{"horn-beast"},{"hugger-mugger"},{"joithead"},{"lewdster"},{"lout"},{"maggot-pie"},{"malt-worm"},{"mammet"},{"measle"},{"minnow"},{"miscreant"},{"moldwarp"},{"mumble-news"},{"nut-hook"},{"pigeon-egg"},{"pignut"},{"puttock"},{"pumpion"},{"ratsbane"},{"scut"},{"skainsmate"},{"strumpet"},{"varlot"},{"vassal"},{"whey-face"},{"wagtail"},{"knave"},{"blind-worm"},{"popinjay"},{"scullian"},{"jolt-head"},{"malcontent"},{"devil-monk"},{"toad"},{"rascal"},{"Basket-Cockle"},{"facade"}}

func RandomInsult() string {
	rand.Seed(time.Now().UTC().UnixNano())
	
	insult1 := ""
	insult2 := ""
	insult3 := ""
	
	for _, i := range rand.Perm(len(insults1)) {
		insult1 = insults1[i].Message
	}
	for _, i := range rand.Perm(len(insults2)) {
		insult2 = insults2[i].Message
	}
	for _, i := range rand.Perm(len(insults3)) {
		insult3 = insults3[i].Message
	}
	
	return "calls out %v, the " + insult1 + " " + insult2 + " " + insult3 + "!"
}


func AddActionInsult(c *Config, ircproj *irc.Connection) error {

	hash := `#insult`

	x := regexp.MustCompile(hash)
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
    
		matches := x.FindAllStringSubmatch(event.Message(), -1)
		if len(matches) > 0 {
			msg := strings.Trim(event.Arguments[1], " ")
			tokens := strings.Split(msg, " ")

			for _, element := range tokens {
				// Don't react to the '#insult' hash
				if strings.HasPrefix(element, "#") {
					continue
				}
				time.Sleep(100)
				ircproj.Actionf(event.Arguments[0], RandomInsult(), element)
			}
		}
	})

	return nil
}

