package core

import (
	"strings"
	//"fmt"
	"io/ioutil"
	//"os"
	"net/http"
	"encoding/json"
)

const endpoint string = "https://api.mollie.nl"
const apiVersion string = "v1"

type Core struct {
	ApiKey string
}

type message struct {
	TotalCount int `json:"totalCount"`
	Offset     int
	Count      int
	Data       interface{}
}

func getUri(action string) string {
	return endpoint + "/" + apiVersion + "/" + action
}

func (c Core) Request(action string, d interface{}) {
	reader := strings.NewReader("bots")
	req, _ := http.NewRequest("GET", getUri(action), reader)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "go-mollie-api/v0.0.0-dev")
	req.Header.Set("Authorization", "Bearer " + c.ApiKey)

	// TODO: check err
	client := &http.Client{}
	resp, _ := client.Do(req)

	//fmt.Println(resp)
	//fmt.Println("Body:")
	//_, _ = io.Copy(os.Stdout, resp.Body)
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	m := message{Data: d}
	json.Unmarshal(data, &m)
	//fmt.Println(m.Data)
}
