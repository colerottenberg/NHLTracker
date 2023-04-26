package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "encoding/json"
)
type Team struct {
    ID           int    `json:"id"`
    Name         string `json:"name"`
    Link         string `json:"link"`
    Abbreviation string `json:"abbreviation"`
    TeamName     string `json:"teamName"`
    LocationName string `json:"locationName"`
    FirstYearOfPlay string `json:"firstYearOfPlay"`
    Division     struct {
        ID            int    `json:"id"`
        Name          string `json:"name"`
        NameShort     string `json:"nameShort"`
        Link          string `json:"link"`
        Abbreviation  string `json:"abbreviation"`
    } `json:"division"`
    Conference struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
        Link string `json:"link"`
    } `json:"conference"`
    Venue struct {
        ID       int `json:"id"`
        Name     string `json:"name"`
        Link     string `json:"link"`
        City     string `json:"city"`
        TimeZone struct {
            ID     string `json:"id"`
            Offset int    `json:"offset"`
            TZ     string `json:"tz"`
        } `json:"timeZone"`
    } `json:"venue"`
    Franchise struct {
        FranchiseId int    `json:"franchiseId"`
        TeamName    string `json:"teamName"`
        Link        string `json:"link"`
    } `json:"franchise"`
    ShortName      string `json:"shortName"`
    OfficialSiteUrl string `json:"officialSiteUrl"`
    FranchiseId    int    `json:"franchiseId"`
    Active         bool   `json:"active"`
}

type Response struct {
    Copyright string `json:"copyright"`
    Teams     []Team `json:"teams"`
}
type TeamRecords struct {
  tName string `json:"teamName"`
  gPlayed int `json:"gamesPlayed"`
  wins int `json:"Wins"`
}
func main() {
//  standingsUrl := "https://statsapi.web.nhl.com/api/v1/standings"
  habsUrl := "https://statsapi.web.nhl.com/api/v1/teams/8"
  //gets JSON data
  teamResp, fetchErr := http.Get(habsUrl)
  if fetchErr != nil {
    fmt.Println("Error fetching standings!")
    return
  }
  //defer standingsResp.Body.close()

  body, err := ioutil.ReadAll(teamResp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
  var habs Response
  if err := json.Unmarshal(body, &habs); err != nil {
    fmt.Println("error unmarshalling!")
    return
  }
  //return variable inside habs
  fmt.Println(habs.Teams[0].TeamName)
  fmt.Println(habs.Teams[0].Abbreviation)
  return
}
