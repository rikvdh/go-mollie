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

type errorReply struct {
	Error struct {
		Message string
	}
}

func getURI(action string) string {
	return endpoint + "/" + apiVersion + "/" + action
}

func (c core) Get(action string, d interface{}) error {
	return c.request(http.MethodGet, action, d, nil)
}

func (c core) Post(action string, d interface{}, postData interface{}) error {
	postStr, err := json.Marshal(postData)
	if err != nil {
		return err
	}

	reader := strings.NewReader(string(postStr))
	return c.request(http.MethodPost, action, d, reader)
}

func (c core) request(method, action string, d interface{}, reader io.Reader) error {
	req, err := http.NewRequest(method, getURI(action), reader)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "go-mollie-api/v1")
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
		if len(data) > 0 {
			r := &errorReply{}
			err := json.Unmarshal(data, r)
			if err == nil && len(r.Error.Message) > 0 {
				return errors.New(r.Error.Message)
			}
		}
		return errors.New(resp.Status)
	}

	return json.Unmarshal(data, &d)
}
