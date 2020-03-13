package slackknocker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Knocker struct {
	Config *Config
}

func NewKnocker(config *Config) *Knocker {
	return &Knocker{Config: config}
}

func (x *Knocker) BuildMessage(args []string) string {
	if x.Config.TextFormat == "" {
		return strings.Join(args, " ")
	} else {
		s := make([]interface{}, len(args))
		for idx, i := range args {
			s[idx] = i
		}
		return fmt.Sprintf(x.Config.TextFormat, s...)
	}
}

func (x *Knocker) BuildPayload(args []string) *Payload {
	return &Payload{
		Text:        x.BuildMessage(args),
		PayloadBase: x.Config.PayloadBase,
	}
}

func (x *Knocker) Post(payload *Payload) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	params := url.Values{}
	params.Add("payload", string(b))

	resp, err := http.PostForm(x.Config.WebhookURL, params)
	if err != nil {
		return err
	}

	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		return nil
	} else {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("%s Status returned")
		}
		return fmt.Errorf("%s Status returned with body %s", string(b))
	}
}
