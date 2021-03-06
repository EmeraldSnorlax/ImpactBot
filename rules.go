package main

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"strings"
)

const rulesChannel = "667494326372139008"
const rulesMessage = "667497572264312832"

var rules = []string{
	"Moderators have the final say. Do not argue with them.",
	"Use the correct channels (ask questions in <#" + help + ">, report bugs on github, etc)",
	"Channel specific rules or topics can be found in the channel description",
	"No trolling, unnecessary tagging, spamming, NSFW content, bullying, or blatant rudeness",
	"No advertising",
	"You will not be able to speak until you verify yourself! Click [here.](https://modulobot.xyz/verify/208753003996512258)",
}

const note = "All staff, including Support, Moderators, and Developers are volunteers. " +
	"They are under _no obligation_ to help you, but are likely to if you are polite."

func rulesHandler(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	reply := discordgo.MessageEmbed{
		Color: prettyembedcolor,
	}

	switch len(args) {
	case 1:
		reply.Title = "Rules"
		reply.Description = buildRules()
	case 2:
		index, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		index-- // Rule numbers are one higher than index
		if index >= len(rules) {
			return errors.New("There are only " + strconv.Itoa(len(rules)) + " rules, " + args[1] + " is too high.")
		}
		if index < 0 {
			return errors.New("Rules are counted from 1, " + args[1] + " is too low")
		}
		reply.Title = "Rule " + strconv.Itoa(index+1)
		reply.Description = rules[index]
	default:
		return errors.New("incorrect number of arguments")
	}

	_, err := discord.ChannelMessageSendEmbed(msg.ChannelID, &reply)
	return err
}

func updateRules() {
	_, err := discord.ChannelMessageEditEmbed(rulesChannel, rulesMessage, &discordgo.MessageEmbed{
		Title:       "Rules",
		Description: buildRules(),
		Color:       prettyembedcolor,
	})
	if err != nil {
		log.Println("Unable to edit rules message with id " + rulesMessage)
	}
}

func buildRules() string {
	var r strings.Builder
	for index, rule := range rules {
		r.WriteString(strconv.Itoa(index + 1))
		r.WriteString(". ")
		r.WriteString(rule)
		r.WriteString("\n")
	}
	r.WriteString("\n")
	r.WriteString(note)

	return r.String()
}
