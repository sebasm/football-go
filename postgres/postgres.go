package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	football "github.com/sebasm/footballgo"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "football"
	password = "football"
	dbname   = "football"
)

type LeagueService struct {
	Path string
	DB   *sql.DB
}

func Open(env string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db, err
}

func (s *LeagueService) League(id int) (*football.League, error) {
	var u football.League
	row := s.DB.QueryRow(`SELECT id, name, code, area_name FROM league WHERE id = $1`, id)
	if err := row.Scan(&u.ID, &u.Name, &u.Code, &u.AreaName); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			fmt.Println(err)
		}
		return nil, err
	}
	return &u, nil
}

func (s *LeagueService) ImportLeague(league *football.LeagueDTO) (string, error) {
	sqlStatement := `
	INSERT INTO league (id, name, code, area_name)
	VALUES ($1, $2, $3, $4)`
	_, err := s.DB.Exec(sqlStatement, league.ID, league.Name, league.Code, league.Area.Name)
	if err != nil {
		panic(err)
	}
	return "", nil
}
