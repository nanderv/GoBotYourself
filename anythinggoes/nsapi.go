package anythinggoes

import (
	"gopkg.in/telegram-bot-api.v4"
	"bot/gofuckyourself"
	"fmt"
	"net/http"
	"encoding/base64"
	"bot/settings"
	"io/ioutil"
	"encoding/xml"
	"time"
	"strings"
)

type customTime struct {
	time.Time
}

func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse("2006-01-02T15:04:05-0700", v)
	if err != nil {
		return err
	}
	*c = customTime{parse}
	return nil
}

type JourneyOptions struct {
	ReisMogelijkheid []JourneyOption
}
type JourneyStop struct {
	Naam  string
	Tijd  customTime
	Spoor string
}
type JourneyPart struct {
	ReisStop  []JourneyStop
	RitNummer string
}
type JourneyOption struct {
	AantalOverstappen int
	ActueleReisTijd   string
	ReisDeel          []JourneyPart
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
func redirectPolicyFunc(req *http.Request, _ []*http.Request) error {
	req.Header.Add("Authorization", "Basic " + basicAuth(settings.GetSettings().NSAPIUSER, settings.GetSettings().NSAPIPASS))
	return nil
}
func nsapi(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var from string
	var to string
	_, err := fmt.Sscanf(update.Message.Text, "/travel %v %v", &from, &to)
	if err != nil {
		return
	}
	client := &http.Client{
		Jar: nil,
		CheckRedirect: redirectPolicyFunc,
	}
	req, err := http.NewRequest("GET", "http://webservices.ns.nl/ns-api-treinplanner?fromStation=" + from + "&toStation=" + to, nil)
	req.Header.Add("Authorization", "Basic " + basicAuth(settings.GetSettings().NSAPIUSER, settings.GetSettings().NSAPIPASS))
	resp, err := client.Do(req)
	body := resp.Body
	b, err := ioutil.ReadAll(body)

	journey := JourneyOptions{}
	err = xml.Unmarshal(b, &journey)
	if err != nil {
		fmt.Println(err)
	}
	timeLocation, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		panic(err)
	}
	for _, loc := range journey.ReisMogelijkheid {
		str := "Journey possible:\n\n```"
		var stationLength int
		stationLength = 0
		for _, leg := range loc.ReisDeel {

			for _, station := range leg.ReisStop {
				if len(station.Naam) > stationLength {
					stationLength = len(station.Naam)
				}
			}
		}
		if loc.ReisDeel[0].ReisStop[0].Tijd.Time.In(timeLocation).Sub(time.Now()) > 0 {

			for _, leg := range loc.ReisDeel {
				str += "\n" + leg.RitNummer + "\n"

				for _, station := range leg.ReisStop {
					str += station.Naam
					str += strings.Repeat(" ", stationLength + 1 - len(station.Naam))
					str += station.Tijd.In(timeLocation).Format("15:04") + "	" + station.Spoor + "\n"

				}

			}
			str += "```"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, str)
			msg.ParseMode = "MarkDown"
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
			return
		} else {
			fmt.Println(loc)
		}

	}

}

var NSApi Module = Module{"nsapi",
	"store value using /store, retrieve using /retreive",
	gofuckyourself.RunAlways,
	nsapi}

