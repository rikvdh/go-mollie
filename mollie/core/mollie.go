package core

import (
	"errors"
	"strings"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

const endpoint string = "https://api.mollie.nl"
const apiVersion string = "v1"

type Core struct {
	ApiKey string
}

func getUri(action string) string {
	return endpoint + "/" + apiVersion + "/" + action
}

func (c Core) Request(action string, d interface{}) error {
	reader := strings.NewReader("bots")
	req, err := http.NewRequest("GET", getUri(action), reader)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "go-mollie-api/v0.0.0-dev")
	req.Header.Set("Authorization", "Bearer " + c.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return errors.New(resp.Status)
	}
	err = json.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	return nil
}
