// Do not use this file directly, do not attemp to compile this source file directly
// Go To lab/3/shodan/main/main.go

package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Client) MyIP() (string, error) {
	res, err := http.Get(fmt.Sprintf("%s/tools/myip?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return "nil", err
	}
	defer res.Body.Close()

	var ret string
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return "nil", err
	}

	return ret, nil
}