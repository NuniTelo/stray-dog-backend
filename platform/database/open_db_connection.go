package database

import "stray-dogs/app/queries"

type Queries struct {
	*queries.AnimalQueries
	*queries.UserQueries
}

func OpenDBConnection() (*Queries, error) {
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		AnimalQueries: &queries.AnimalQueries{DB: db},
		UserQueries:   &queries.UserQueries{DB: db},
	}, nil
}
