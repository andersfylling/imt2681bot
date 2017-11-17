package currencyconversation

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/andersfylling/concurrencyparser"
	"github.com/bwmarrin/discordgo"

	"github.com/s1kx/unison"
	"github.com/s1kx/unison/events"
)

// Hook reacts to a new message written on the server
// Note! the bot needs permissions to read messages and write messages on that
//       channel
var Hook = &unison.EventHook{
	Name:        "currency",
	Description: "Logs messages in given channels to support reviewing them later.",
	Events: []events.EventType{
		events.MessageCreateEvent,
	},
	OnEvent: unison.EventHandlerFunc(chatlogAction),
}

var lookUpSite = "https://frozen-reef-21113.herokuapp.com/latest"

func currenciesToJsonStr(base, target string) string {
	return `{"baseCurrency": "` + base + `", "targetCurrency":"` + target + `"}`
}

func getCurrencyRate(data string) (float64, error) {
	payload := []byte(data)
	logrus.Info("request payload: " + string(payload))
	client := &http.Client{}

	req, err := http.NewRequest("POST", lookUpSite, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return 0.0, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	val, _ := strconv.ParseFloat(string(body), 64)
	return val, nil
}

func findCurrencyRate(exc *concurrencyparser.ExchangeRate) (string, error) {

	if exc.Base == "EUR" {
		logrus.Info("[Chatlog] Base was EURO")
		payload := currenciesToJsonStr("EUR", exc.Target)
		rate, err := getCurrencyRate(payload)
		if err != nil {
			return "", errors.New("Unable to get currency rate")
		}

		return strconv.FormatFloat(rate, 'f', -1, 64) + exc.Target + "/EUR", nil
	}

	payloadA := currenciesToJsonStr("EUR", exc.Base)
	rateA, err := getCurrencyRate(payloadA)
	if err != nil {
		return "", err
	}
	logrus.Info("[Chatlog] rateA: " + strconv.FormatFloat(rateA, 'f', -1, 64))

	var rateB float64 = 1.0
	if exc.Target != "EUR" {
		payloadB := currenciesToJsonStr("EUR", exc.Target)
		rateB, err = getCurrencyRate(payloadB)
		if err != nil {
			return "", err
		}
	}
	logrus.Info("[Chatlog] rateB: " + strconv.FormatFloat(rateB, 'f', -1, 64))

	rate := rateB / rateA
	return strconv.FormatFloat(rate, 'f', -1, 64) + exc.Target + "/" + exc.Base, nil
}

func chatlogAction(ctx *unison.Context, event *events.DiscordEvent) (bool, error) {
	var m *discordgo.Message

	// Check event type
	if event.Type != events.MessageCreateEvent {
		return true, nil
	}

	switch ev := event.Event.(type) {
	default:
		return true, nil
	case *discordgo.MessageCreate:
		m = ev.Message
	}

	// Log request message
	logrus.Infof("[chatlog] <%s>: %s", m.Author.Username, m.ContentWithMentionsReplaced())

	// if message is from the bot itself, don't respond
	if m.Author.ID == ctx.Bot.Discord.State.User.ID {
		return true, nil
	}

	// response
	var response string

	// parse request message
	exc, err := concurrencyparser.ParseStr(m.ContentWithMentionsReplaced())
	if err != nil {
		response = err.Error()
	} else {
		logrus.Info("[Chatlog] Base: " + exc.Base + ", Target: " + exc.Target)
		rate, err := findCurrencyRate(exc)
		if err != nil {
			response = err.Error()
		} else {
			response = "rate: " + rate
		}
	}

	_, err = ctx.Bot.Discord.ChannelMessageSend(m.ChannelID, response)

	return true, err
}
