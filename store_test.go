package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"testing"
)

type StoreSuit struct {
	suite.Suite
	store *dbStore
	db    *sql.DB
}

func (s *StoreSuit) SetupSuite() {
	connString := "dbname=bird_encyclopedia_test sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		s.T().Fatal(err)
	}

	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuit) SetupTest() {
	_, err := s.db.Query("DELETE FROM birds")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuit) TearDownSuite() {
	s.db.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuit)
	suite.Run(t, s)
}

func (s *StoreSuit) TestCreateBird() {
	_ = s.store.CreateBird(&Bird{
		Description: "test description",
		Species:     "test species",
	})

	res, err := s.db.Query(`SELECT COUNT(*) FROM birds WHERE description='test description' AND species='test species'`)
	if err != nil {
		s.T().Fatal(err)
	}

	var count int
	for res.Next() {
		err := res.Scan(&count)

		if err != nil {
			s.T().Error(err)
		}
	}

	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}

func (s *StoreSuit) TestGetBird() {
	_, err := s.db.Query(`INSERT INTO birds (species, description) VALUES('bird','description')`)
	if err != nil {
		s.T().Fatal(err)
	}

	birds, err := s.store.GetBirds()
	if err != nil {
		s.T().Fatal(err)
	}

	nBirds := len(birds)
	if nBirds != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", nBirds)
	}

	expectedBird := Bird{"bird", "description"}
	if *birds[0] != expectedBird {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedBird, *birds[0])
	}
}
