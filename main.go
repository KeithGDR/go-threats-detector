package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"drixevel.dev/threats-detector/internal/config"
	"drixevel.dev/threats-detector/internal/services"
	"github.com/gorilla/mux"
)

type Config struct {
	API_Key string `json:"api_key"`
	Format  string `json:"format"`
}

func main() {
	fmt.Println("Threats Detector Initialized...")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", onMainPage)
	myRouter.HandleFunc("/threat/{url}", onCheckForThreat)

	http.ListenAndServe(":8080", myRouter)
}

func onMainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please use '/threat' to check for threats.")
}

func onCheckForThreat(w http.ResponseWriter, r *http.Request) {
	var cfg Config

	//Parse the config.json file.
	if err := config.ParseConfig(&cfg); err != nil {
		log.Fatal(err)
	}

	vars := mux.Vars(r)
	url_check := vars["url"]

	Threats := services.CheckURL(cfg, url_check)

	json.NewEncoder(w).Encode(Threats)
}
