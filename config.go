package slackknocker

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	WebhookURL string `json:"webhook_url"`
	TextFormat string `json:"text_format,omitempty"`
	PayloadBase
}

func LoadConfigFile(path string) (rconfig *Config, rerr error) {
	rerr = func() error {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		var c Config
		if err := json.Unmarshal(b, &c); err != nil {
			return err
		}

		rconfig = &c
		return nil
	}()
	return
}
