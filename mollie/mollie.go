package mollie

import (
	"strings"
	"fmt"
	"io"
	"os"
	"net/http"
)

const endpoint string = "https://api.mollie.nl"
const apiVersion string = "v1"

const apiKey string = "test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU"

func getUri(action string) string {
	return endpoint + "/" + apiVersion + "/" + action
}

func Request() {
	reader := strings.NewReader("bots")
	req, _ := http.NewRequest("GET", getUri("issuers"), reader)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "go-mollie-api/v0.0.0-dev")
	req.Header.Set("Authorization", "Bearer " + apiKey)

	// TODO: check err
	client := &http.Client{}
	resp, _ := client.Do(req)

	fmt.Println(resp)
	fmt.Println("Body:")
	_, _ = io.Copy(os.Stdout, resp.Body)
	fmt.Println("")

	resp.Body.Close()
}
