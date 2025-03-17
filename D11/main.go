package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	es "github.com/elastic/go-elasticsearch/v8"
)

type Movie struct {
	Name        string
	Description string
	Rating      float64
}

func main() {

	fmt.Println("WOI")
	client, err := es.NewClient(es.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "elastic",
		Password: "Qbkfzlvf",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("COBA")

	moviesIndex := "movies"

	// Create index
	client.Indices.Create(moviesIndex)

	movie := Movie{
		Name:        "Interstellar",
		Description: "The Best",
		Rating:      10,
	}

	data, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	// Insert data
	client.Index(moviesIndex, bytes.NewReader(data))

	// GET by ID

	res, err := client.Get(moviesIndex, "cW1TpJUB6HCjoWL3LF0q")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	r, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Result get by ID: %v\n\n", string(r))

	// Get by serach

	query := `{
		"query": {
			"match_all": {}
		}
	}`
	res, err = client.Search(
		client.Search.WithIndex(moviesIndex),
		client.Search.WithBody(bytes.NewReader([]byte(query))),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	r, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Result get by search: %v\n\n", string(r))

	// Update
	movieUpdate := Movie{
		Name:        "Interstillir",
		Description: "The Best Of The Best",
		Rating:      100,
	}

	movieUpdateDoc := map[string]interface{}{
		"doc": movieUpdate,
	}

	dataUpdate, err := json.Marshal(movieUpdateDoc)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	_, err = client.Update(moviesIndex, "cm1bpJUB6HCjoWL3m10x", bytes.NewReader(dataUpdate), client.Update.WithRefresh("true"))
	if err != nil {
		fmt.Println(err)
		return
	}

	client.Delete(moviesIndex, "c21bpJUB6HCjoWL34V1I")
}
