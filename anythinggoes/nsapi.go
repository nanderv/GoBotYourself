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
	ReisStop []JourneyStop
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
func nsapi(update tgbotapi.Update, _ *tgbotapi.BotAPI) (string, bool) {

	var from string
	var to string
	_, err := fmt.Sscanf(update.Message.Text, "/travel %v %v", &from, &to)
	if err != nil {
		return err.Error(), false
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

	boop := JourneyOptions{}

	err = xml.Unmarshal(b, &boop)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(boop)
	for _, loc := range boop.ReisMogelijkheid {
		str := "Journey possible:\n\n"

		for _, leg := range loc.ReisDeel {
			str += "\n"
			for _, station := range leg.ReisStop {
				str += station.Naam + " " + station.Tijd.Format("15:04") + "  " + station.Spoor + "\n"

			}

		}

		return str, true
	}
	return " ", false
}

var NSApi Module = Module{"nsapi",
	"store value using /store, retrieve using /retreive",
	gofuckyourself.RunAlways,
	gofuckyourself.IFReplier(nsapi)}

