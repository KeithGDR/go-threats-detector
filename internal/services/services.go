package services

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Threat struct {
	URL        string `json:"url"`
	Suspicious bool   `json:"suspicious"`
}

type Config struct {
	API_Key string `json:"api_key"`
	Format  string `json:"format"`
}

type Response struct {
	XMLName xml.Name `xml:"response"`
	Meta    Meta     `xml:"meta"`
	Results Results  `xml:"results"`
}

type Meta struct {
	Timestamp string `xml:"timestamp"`
	Serverid  string `xml:"serverid"`
	Requestid string `xml:"requestid"`
}

type Results struct {
	Url0 Url0 `xml:"url0"`
}

type Url0 struct {
	Url               string `xml:"url"`
	In_database       bool   `xml:"in_database"`
	Phish_id          int    `xml:"phish_id"`
	Phish_detail_page string `xml:"phish_detail_page"`
	Verified          bool   `xml:"verified"`
	Verified_at       string `xml:"verified_at"`
	Valid             bool   `xml:"valid"`
}

func CheckURL(cfg Config, url string) []Threat {
	req, err := http.NewRequest(http.MethodPost, "https://checkurl.phishtank.com/checkurl/index.php", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "phishtank/Drixevel")

	url = fmt.Sprintf("https://%s", url)

	q := req.URL.Query()
	q.Add("url", url)
	q.Add("format", cfg.Format)
	q.Add("app_key", cfg.API_Key)

	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request to the server")
		return nil
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(resp.Status)
	// fmt.Println(string(responseBody))

	var response Response
	if xml.Unmarshal([]byte(string(responseBody)), &response) == nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}

	var Threats = []Threat{
		{URL: url, Suspicious: response.Results.Url0.Valid},
	}

	return Threats
}
