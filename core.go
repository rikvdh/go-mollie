// Copyright 2017 The Go-Mollie Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mollie

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const endpoint string = "https://api.mollie.nl"
const apiVersion string = "v1"

type core struct {
	apiKey string
}

func getURI(action string) string {
	return endpoint + "/" + apiVersion + "/" + action
}

func (c core) Get(action string, d interface{}) error {
	return c.request("GET", action, d, nil)
}

func (c core) Post(action string, d interface{}, postData interface{}) error {
	postStr, err := json.Marshal(postData)
	if err != nil {
		return err
	}

	reader := strings.NewReader(string(postStr))
	return c.request("POST", action, d, reader)
}

func (c core) request(method, action string, d interface{}, reader io.Reader) error {
	req, err := http.NewRequest(method, getURI(action), reader)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "go-mollie-api/v0.0.0-dev")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := resp.Body.Close(); err != nil {
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
