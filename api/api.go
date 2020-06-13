package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	football "github.com/sebasm/footballgo"
)

func GetLeague(code string) (*football.LeagueDTO, error) {
	url := "http://api.football-data.org/v2/competitions/WC"
	token := ""
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Auth-Token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var responseObject football.LeagueDTO
	json.NewDecoder(resp.Body).Decode(&responseObject)
	fmt.Println(responseObject)
	return &responseObject, nil
}
