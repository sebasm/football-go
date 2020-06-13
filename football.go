package football

type League struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Code     string `json:"Code"`
	AreaName string `json:"AreaName"`
}

type Team struct {
}

type Player struct {
}

type LeagueDTO struct {
	ID   int     `json:"ID"`
	Name string  `json:"Name"`
	Code string  `json:"Code"`
	Area AreaDTO `json:"area"`
}

type TeamDTO struct {
}

type AreaDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
}

type LeagueService interface {
	// League(id int) (*League, error)
	ImportLeague(league *LeagueDTO) (string, error)
}

type LeagueAPI interface {
	GetLeague(code string) (*LeagueDTO, error)
	GetTeams(league string) (*TeamDTO, error)
	GetTeam(teamID int) (*TeamDTO, error)
}
