package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/keybase/go-keybase-chat-bot/kbchat"
	"github.com/prometheus/alertmanager/notify/webhook"
)

func handleWebhook(kbc *kbchat.API, user string, tmpl *template.Template) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading webhook post: %+v", err)
		}

		wh := &webhook.Message{}
		err = json.Unmarshal(buf, wh)
		if err != nil {
			log.Printf("Error parsing webhook post: %+v", err)
		}

		log.Printf("Received and parsed incoming webook: %+v", wh)

		tlfName := fmt.Sprintf("%s,%s", kbc.GetUsername(), user)
		log.Printf("tlfName: %s", tlfName)

		writer := bytes.NewBufferString("")

		tmpl.Execute(writer, *wh)
		log.Printf("%s", writer.String())

		if _, err = kbc.SendMessageByTlfName(tlfName, writer.String()); err != nil {
			log.Printf("Error sending message: %+v", err)
		}
	}

}

func main() {
	var kbLoc string
	var kbc *kbchat.API
	var listenPort int
	var user string
	var err error

	flag.StringVar(&kbLoc, "keybase", "keybase", "the location of the Keybase app")
	flag.IntVar(&listenPort, "port", 3000, "Port to listen for webhooks")
	flag.StringVar(&user, "user", "", "Keybase user to send message to")
	flag.Parse()

	log.Printf("%+v", os.Environ())

	tmpl, err := template.New("default.tmpl").Funcs(template.FuncMap{
		"toUpper": strings.ToUpper,
		"join":    strings.Join,
	}).ParseFiles("default.tmpl")

	if err != nil {
		log.Panicf("Unable to parse template: %+v", err)
	}

	if kbc, err = kbchat.Start(kbchat.RunOptions{
		KeybaseLocation:    kbLoc,
		StartService:       true,
		DisableBotLiteMode: true,
		Oneshot: &kbchat.OneshotOptions{
			PaperKey: os.Getenv("KEYBASE_PAPERKEY"),
			Username: os.Getenv("KEYBASE_USERNAME"),
		},
	}); err != nil {
		log.Fatalf("Error creating API: %+v", err)
	}

	http.HandleFunc("/webhook", handleWebhook(kbc, user, tmpl))

	log.Printf("Listening on port %d", listenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil))

}
