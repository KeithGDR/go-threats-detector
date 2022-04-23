package services

import "encoding/xml"

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
