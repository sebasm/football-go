package football

type League struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Code     string `json:"Code"`
	AreaName string `json:"AreaName"`
}

type LeagueService interface {
	League(id int) (*League, error)
}
