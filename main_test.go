package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/prometheus/alertmanager/notify/webhook"
	"github.com/prometheus/alertmanager/template"
)

func TestWebhookFormat(t *testing.T) {

	msg := webhook.Message{
		Data: &template.Data{
			Receiver: "foo-receiver",
			Status:   "Firing",
			Alerts: template.Alerts{
				template.Alert{
					Status: "Firing",
					Labels: template.KV{
						"foo": "bar",
						"biz": "baz",
					},
				},
			},
		},
		Version:  "v1",
		GroupKey: "dunno",
	}

	buf, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%s", string(buf))

}
