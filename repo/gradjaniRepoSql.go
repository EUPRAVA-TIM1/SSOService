package repo

import (
	"EuprvaSsoService/data"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type GradjaninRepoSql struct {
	pass string
	host string
	port string
}

const schema = "eupravaSSO"

/*
NewGrdjaninRepoSql
generates new GradjaniRepo struct and accepts port , pass and Host strings in that order
*/
func NewGrdjaninRepoSql(port, pass, host string) data.GradjaniRepo {
	return &GradjaninRepoSql{
		pass: pass,
		host: host,
		port: port,
	}
}

func (g GradjaninRepoSql) GetByJmbg(jmbg string) (*data.Gradjanin, error) {
	db, err := g.OpenConnection()
	if err != nil {
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()

	query := "select JMBG,Ime,Prezime,Adresa,BrojTelefona,Email,Lozinka,PTT,Naziv from Gradjani g inner join Opstina o on g.OpstinaID = o.PTT where g.JMBG = ?;"
	rows, err := db.Query(query, jmbg)
	if err != nil {
		panic(err)
		return nil, errors.New("There has been problem with reading from db")
	}

	var JMBG string
	var Ime string
	var Prezime string
	var Adresa string
	var BrojTelefona string
	var Email string
	var Lozinka string
	var PTT string
	var Naziv string

	for rows.Next() {
		err = rows.Scan(&JMBG, &Ime, &Prezime, &Adresa, &BrojTelefona, &Email, &Lozinka, &PTT, &Naziv)
		if err != nil {
			panic(err.Error())
		}
		return &data.Gradjanin{
			Ime:          Ime,
			Prezime:      Prezime,
			JMBG:         JMBG,
			Adresa:       Adresa,
			BrojTelefona: BrojTelefona,
			Email:        Email,
			Lozinka:      Lozinka,
			Opstina: data.Opstina{
				PTT:   PTT,
				Naziv: Naziv,
			},
		}, nil
	}

	return nil, errors.New("No gradjanin with that JMBG found")
}

func (g GradjaninRepoSql) OpenConnection() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s:%s)/%s", g.pass, g.host, g.port, schema))
}
