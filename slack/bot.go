package slack

import (
	"fmt"
	"freefrom.space/slackBot/conf"
	"reflect"
	"strings"

	slackApi "freefrom.space/slackBot/api"
	"github.com/slack-go/slack"
)

var (
	SlackToken        = conf.GetConf().Slack.Token
	SpecificChannelID = conf.GetConf().Slack.Channel
)

func RunBot() {
	fmt.Println("RunBot function started")
	api := slack.New(SlackToken)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Println("Received an event") //
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			fmt.Printf("Received message: %s\n", ev.Text) //
			if ev.Channel == SpecificChannelID && strings.HasPrefix(ev.Text, ".add ") {
				pubkey := strings.TrimPrefix(ev.Text, ".add ")
				slackApi.SendPostRequest(pubkey)
			}
		default:
			fmt.Printf("Received an event of a different type: %s\n", reflect.TypeOf(msg.Data)) //
		}
	}
	fmt.Println("RunBot function ended") //
}
