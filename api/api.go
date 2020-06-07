package api

import (
	"fmt"

	football "github.com/sebasm/footballgo"
)

func GetLeague(code string) (*football.League, error) {
	url := "http://api.football-data.org/v2"
	token := ""
	fmt.Println(url)
	fmt.Println(token)
	return nil, nil
}
