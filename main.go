package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

/*
type ScientificName struct {
	Genus   string `json:"genus"`
	Species string `json:"species"`
}

type SpeciesLevelTaxon struct {
	ScientificName ScientificName `json:"scientificName"`
	EnglishName    string         `json:"englishName"`
}
*/

type Variety struct {
	ID                 string           `json:"id"`
	Pedigree           string           `json:"pedigree"`
	CommonName         string           `json:"name"`
	AllSynonyms        []string         `json:"allSynonyms"`
	SkinColor          string           `json:"skinColor"`
	CountryOfOrigin    *CountryOfOrigin `json:"countryOfOrigin"`
	YearOfIntroduction int              `json:"yearOfIntroduction"`
}

type CountryOfOrigin struct {
	Name string `json:"countryName"`
	Code string `json:"countryCode"`
}

var grapevarieties []Variety

func getVarieties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(grapevarieties)
	if err != nil {
		fmt.Printf("json encoding error")
		return
	}
}

func getVariety(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range grapevarieties {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				fmt.Printf("json encoding error")
				return
			}
			return
		}
	}
}

func createVariety(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var variety Variety
	_ = json.NewDecoder(r.Body).Decode(&variety)
	variety.ID = strconv.Itoa(rand.Intn(10e6))
	grapevarieties = append(grapevarieties, variety)
	err := json.NewEncoder(w).Encode(variety)
	if err != nil {
		fmt.Printf("json encoding error")
		return
	}

}
func updateVariety(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range grapevarieties {
		if item.ID == params["id"] {
			grapevarieties = append(grapevarieties[:index], grapevarieties[index+1:]...)
			var variety Variety
			_ = json.NewDecoder(r.Body).Decode(&variety)
			variety.ID = strconv.Itoa(rand.Intn(10e6))
			grapevarieties = append(grapevarieties, variety)
			err := json.NewEncoder(w).Encode(variety)
			if err != nil {
				fmt.Printf("json encoding error")
				return
			}
		}
	}
}

func deleteVariety(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range grapevarieties {

		if item.ID == params["id"] {
			grapevarieties = append(grapevarieties[:index], grapevarieties[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(grapevarieties)
	if err != nil {
		fmt.Printf("json encoding error")
		return
	}
}

func main() {
	r := mux.NewRouter()

	grapevarieties = append(grapevarieties, Variety{ID: "1", Pedigree: "Possible natural Piedirosso x Casavecchia cross", CommonName: "Abbuoto", AllSynonyms: []string{"Aboto", "Cecubo"}, SkinColor: "Black", CountryOfOrigin: &CountryOfOrigin{Name: "Italy", Code: "IT"}, YearOfIntroduction: 1984})
	grapevarieties = append(grapevarieties, Variety{ID: "2", Pedigree: "Irsai Olivér × Mátrai muskotály", CommonName: "Csaba gyöngye", AllSynonyms: []string{"Csabagyöngye", "Pearl of Csaba", "Perle di Csaba"}, SkinColor: "Blanc", CountryOfOrigin: &CountryOfOrigin{Name: "Hungary", Code: "HU"}, YearOfIntroduction: 1905})

	r.HandleFunc("/varieties", getVarieties).Methods("GET")
	r.HandleFunc("/varieties/{id}", getVariety).Methods("GET")
	r.HandleFunc("/varieties", createVariety).Methods("POST")
	r.HandleFunc("/varieties/{id}", updateVariety).Methods("PUT")
	r.HandleFunc("/varieties/{id}", deleteVariety).Methods("DELETE")

	fmt.Printf("Starting server at port: 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
