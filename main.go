package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
)

type CsvLine struct {
	Name    string
	Episode string
	Status  string
}

func main() {
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
	handleRequests()
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

func GetAnimeList() {
	currentShows := make([]*CsvLine, 0)
	lines, err := ReadCsv("AnimeList.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		data := &CsvLine{
			Name:    line[0],
			Episode: line[1],
			Status:  line[3],
		}
		// fmt.Println(data.Name + " " + data.Episode + " " + data.Status)
		currentShows = append(currentShows, data)
		// fmt.Printf("Final List: %d", len(currentShows))
	}
}
