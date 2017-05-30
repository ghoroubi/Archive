package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	distance = 50
)

type HereNow struct {
	Women_Count       int     `json:"count_women"`
	Men_Count         int     `json:"count_men"`
	Age_Average_Women float64 `json:"age_average women"`
	Age_Average_Men   float64 `json:"age_average Men"`
}
type SimpleUsers struct {
	ID         int       `json:"id"`
	Gender     int       `json:"gender"`
	Birth_date time.Time `json:"birth_date"`
	Age        int       `json:"age"`
}
type SimpleVenue struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Lat_Bar          float64  `json:"lat_bar"`
	Lng_Bar          float64  `json:"lng_bar"`
	Phone            string   `json:"phone"`
	Facebook         string   `json:"facebook"`
	FormattedAddress []string `json:"formattedAddress"`
	HereNow          HereNow  `json:"herenow"`
}
type Venues struct {
	Response Response `json:"response"`
}
type Response struct {
	Group []Group `json:"groups"`
}
type Group struct {
	Item []Item `json:"items"`
}
type Item struct {
	ReferralId string `json:"referralId"`
	Venue      Venue  `json:"venue"`
	Stats      Stats  `json:"stats"`
	Verified   bool   `json:"verified"`
}

// type Meta struct {
// 	Code      int    `json:"code"`
// 	RequestId string `json:"requestId"`
// }
type Venue struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Contact    Contact    `json:"contact"`
	Location   Location   `json:"location"`
	Categories []Category `json:"categories"`
	Photos     Photos     `json:"photos"`
	//StoreId    string     `json:"storeid"`
	//Url        string     `json:"url"`
	//Stats Stats `json:"stats"`
}

type Stats struct {
	CheckinsCount int `json:"checkinsCount"`
	UsersCount    int `json:"usersCount"`
	TipCount      int `json:"tipCount"`
	VisitsCount   int `json:"visitsCount"`
}

type Contact struct {
	Phone          string `json:"phone"`
	FormattedPhone string `json:"formattedPhone"`
	Twitter        string `json:"twitter"`
	Facebook       string `json:"facebook"`
}
type Location struct {
	Address          string           `json:"address"`
	CrossStreet      string           `json:"crossStreet"`
	Lat              float64          `json:"lat"`
	Lng              float64          `json:"lng"`
	LabeledLatLngs   []LabeledLatLngs `json:"labeledLatLngs"`
	PostalCode       string           `json:"postalCode"`
	Cc               string           `json:"cc"`
	City             string           `json:"city"`
	State            string           `json:"state"`
	Country          string           `json:"country"`
	FormattedAddress []string         `json:"formattedAddress"`
	IsFuzzed         bool             `json:"isFuzzed"`
	Distance         int              `json:"distance"`
}
type Category struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	PluralName string     `json:"pluralName"`
	ShortName  string     `json:"shortName"`
	Icon       Icon       `json:"icon"`
	Primary    bool       `json:"primary"`
	Categories []Category `json:"categories,omitempty"`
}
type Photos struct {
	Count  int             `json:"count"`
	Groups []PhotoGrouping `json:"groups"`
}
type PhotoGrouping struct {
	Items []Photo `json:"items"`
}
type Photo struct {
	ID        string      `json:"id"`
	CreatedAt int         `json:"createdAt"`
	Source    PhotoSource `json:"source"`
	Prefix    string      `json:"prefix"`
	Suffix    string      `json:"suffix"`
	Demoted   bool        `json:"demoted"`
	Width     int         `json:"width"`
	Height    int         `json:"height"`

	Visibility string `json:"visibility"`
}
type PhotoSource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Icon struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}
type LabeledLatLngs struct {
	Label string  `json:"label"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
}
type BarCollection struct {
}

func (col *BarCollection) GetVenuesWithLatitudeAndLongitude(lat, lng string) []SimpleVenue {
	lt, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		fmt.Println("err", err)
	}
	lg, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		fmt.Println("err", err)
	}
	FSqrUrl := fmt.Sprintf("http://api.foursquare.com/v2/venues/explore?ll=%f,%f&client_id=4LDYRYKCKHFI5IBAHTH1FS4L0DRMLLJCMBGRI1X4NCWEDVLH&client_secret=ZLDFJ2E3SQOXEIMCE2I1C3W4Z5WYBDFONXDDYGOVE0UNCOHM&v=20161107", lt, lg)
	return GetLocations(FSqrUrl)
}
func GetLocations(FSqrUrl string) []SimpleVenue {

	venue := Venues{}

	res, err := http.Get(FSqrUrl)
	if err != nil {
		log.Println("getVenues error: " + err.Error())
	}
	defer res.Body.Close()
	str, _ := ioutil.ReadAll(res.Body)
	body := []byte(str)
	err1 := json.Unmarshal(body, &venue)
	//fmt.Println("venue", venue)
	if err1 != nil {
		fmt.Println("error in numarshal", err1)
	}
	items := venue.Response.Group[0].Item
	//fmt.Println(len(items))
	simple_venue := make([]SimpleVenue, len(items))
	for i := 0; i < len(items); i++ {
		simple_venue[i].ID = items[i].Venue.ID
		simple_venue[i].Name = items[i].Venue.Name
		simple_venue[i].Lat_Bar = items[i].Venue.Location.Lat
		simple_venue[i].Lng_Bar = items[i].Venue.Location.Lng
		simple_venue[i].Phone = items[i].Venue.Contact.Phone
		simple_venue[i].Facebook = items[i].Venue.Contact.Facebook
		// simple_venue[i].City = items[i].Venue.Location.Country
		// simple_venue[i].Country = items[i].Venue.Location.City
		simple_venue[i].FormattedAddress = items[i].Venue.Location.FormattedAddress
		simpleusers := list_user(simple_venue[i].Lat_Bar, simple_venue[i].Lng_Bar)
		herenow := here_now(simpleusers)
		simple_venue[i].HereNow = herenow
	}

	return simple_venue
}

func list_user(lat_bar float64, lng_bar float64) []SimpleUsers {

	var sUsers []SimpleUsers
	var condition, query string
	condition = " where users.is_active =true"
	if lat_bar != 0 && lng_bar != 0 && distance != 0 {
		condition += fmt.Sprintf(" and (point( %f,%f) <@> point(users.lat_current, users.lng_current)) < %d", lat_bar, lng_bar, distance)
	}
	query = "select users.id,users.gender,users.birth_date from users " + condition + " "
	//ORDER BY  random()
	err := db.Select(&sUsers, query)
	if err != nil {
		fmt.Println("error in list_user ", err)

	}
	for i := 0; i < len(sUsers); i++ {
		birth_date := sUsers[i].Birth_date
		age_user := age(birth_date)
		sUsers[i].Age = age_user
	}

	//fmt.Println("total count", total_count)
	//men_avg = float64(men_count) / float64(total_count)
	//fmt.Println("men_avg", men_avg)
	//women_avg = float64(women_count) / float64(total_count)
	//fmt.Println("men_avg", women_avg)
	//sUsers[i].Age_Average_Men = men_avg
	//sUsers[i].Age_Average_Women = women_avg
	//fmt.Println("HereNOW", sUsers)

	//total_count = men_count + women_count

	return sUsers
}
func age(birth_date time.Time) int {
	now := time.Now()
	years := now.Year() - birth_date.Year()
	if now.YearDay() < birth_date.YearDay() {
		years--
	}
	return years
}
func here_now(sUsers []SimpleUsers) HereNow {
	herenow := HereNow{}
	var men_count, women_count, total_count int
	var women_avg, men_avg float64
	total_count = 0
	men_count = 0
	women_count = 0
	for i := 0; i < len(sUsers); i++ {
		if sUsers[i].Gender == 1 {
			men_count++

		} else {
			women_count++

		}

	}
	total_count = men_count + women_count
	women_avg = float64(women_count) / float64(total_count)
	men_avg = float64(men_count) / float64(total_count)
	herenow.Age_Average_Men = men_avg
	herenow.Age_Average_Women = women_avg
	herenow.Men_Count = men_count
	herenow.Women_Count = women_count
	return herenow
}
