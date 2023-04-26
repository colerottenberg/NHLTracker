package models

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
