package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"drixevel.dev/threatschecker/internal/config"
	"drixevel.dev/threatschecker/internal/services"
	"github.com/gorilla/mux"
)

type Threat struct {
	URL        string `json:"url"`
	Suspicious bool   `json:"suspicious"`
}

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

	req, err := http.NewRequest(http.MethodPost, "https://checkurl.phishtank.com/checkurl/index.php", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "phishtank/Drixevel")

	url_check = fmt.Sprintf("https://%s", url_check)

	q := req.URL.Query()
	q.Add("url", url_check)
	q.Add("format", cfg.Format)
	q.Add("app_key", cfg.API_Key)

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(resp.Status)
	// fmt.Println(string(responseBody))

	var response services.Response
	if xml.Unmarshal([]byte(string(responseBody)), &response) == nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}

	var Threats = []Threat{
		{URL: url_check, Suspicious: response.Results.Url0.Valid},
	}

	json.NewEncoder(w).Encode(Threats)
}
