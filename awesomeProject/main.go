package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*
		db, err := sql.Open("mysql", "root:@/feladat")
		if err != nil {
			panic(err)
		}
		// See "Important settings" section.
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	*/
	db, err := sql.Open("mysql", "root:@/feladat")
	if err != nil {
		fmt.Print("Error")
		//panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		fmt.Print("Connection error")
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	var nyeremeny []Nyeremeny

	nyeremeny, err = getAllNyeremenyek(db)
	fmt.Printf("Nyeremenyek found: %v\n", nyeremeny)
}

func getAllNyeremenyek(db *sql.DB) ([]Nyeremeny, error) {
	var nyeremenyek []Nyeremeny
	rows, err := db.Query("SELECT * FROM nyeremeny")
	if err != nil {
		fmt.Errorf("allNyeremenyek %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var nyeremeny Nyeremeny
		if err := rows.Scan(&nyeremeny.ID, &nyeremeny.HuzasId, &nyeremeny.Talalat, &nyeremeny.Darab, &nyeremeny.Ertek); err != nil {
			return nil, fmt.Errorf("allNyeremenyek %v", err)
		}
		nyeremenyek = append(nyeremenyek, nyeremeny)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("allNyeremenyek %v", err)
	}
	return nyeremenyek, nil
}

type Nyeremeny struct {
	ID      int
	HuzasId int
	Talalat int
	Darab   int
	Ertek   int
}

/*
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}
*/
