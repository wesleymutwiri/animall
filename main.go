package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type CsvLine struct {
	Name    string
	Episode string
	Status  string
}

func main() {
	db, err := sql.Open("sqlite3", "blog.db")
	checkErr(err)
	_, err = db.Exec(`CREATE TABLE if not exists shows(id integer, name TEXT, episode TEXT, status TEXT, PRIMARY KEY("id" AUTOINCREMENT)) `)
	checkErr(err)
	fmt.Println("DB created and opened successfully")
	defer db.Close()
	GetAnimeList(db)
	// fName := "latestanime.csv"
	// file, err := os.create(fName)
	// if err != nil {
	// 	log.Fatalf("Cannot create file: %q: %s\n", fName, err)
	// }
	// defer file.Close()
	// writer := csv.NewWriter(file)
	// defer writer.Flush()

	// writer.Write([]string{""})
	// for reading the csv file provided
	// handleRequests()
}

func handleRequests() {
	http.HandleFunc("/schedule", getLatestAnime)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func getLatestAnime(w http.ResponseWriter, r *http.Request) {

}

func ReadCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func GetAnimeList(db *sql.DB) {
	// currentShows := make([]*CsvLine, 0)
	lines, err := ReadCsv("AnimeList.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		data := CsvLine{
			Name:    line[0],
			Episode: line[1],
			Status:  line[3],
		}
		saveShows(db, data.Name, data.Episode, data.Status)
		fmt.Println("Schedule saved", " ", data.Name, " ", data.Episode, " ", data.Status)
		// currentShows = append(currentShows, data)
		// fmt.Printf("Final List: %d", len(currentShows))
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func saveShows(db *sql.DB, name string, episode string, status string) {
	tx, err := db.Begin()
	checkErr(err)
	stmt, _ := tx.Prepare("insert into shows(name, episode, status) values (?,?,?)")
	checkErr(err)
	_, err = stmt.Exec(name, episode, status)
	checkErr(err)
	fmt.Println("saved shows successfully")
	tx.Commit()
}

func getAllShows(db *sql.DB) {
	rows, err := db.Query("select * from schedule")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var episode string
		var status string
		rows.Scan(&id, &name, &episode, &status)
		fmt.Println("Schedule: ", id, " ", name, " ", episode, " ", status)
	}
}
