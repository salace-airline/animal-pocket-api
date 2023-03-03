package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Fish struct {
	gorm.Model
	Name               string         `json:"name" gorm:"text;not null"`
	MonthNorthern      string         `json:"month_northern" gorm:"text;null"`
	MonthArrayNorthern pq.Int32Array  `json:"month_array_northern" gorm:"type:integer[];not null"`
	MonthSouthern      string         `json:"month-southern" gorm:"text;null"`
	MonthArraySouthern pq.Int32Array  `json:"month_array_southern" gorm:"type:integer[];not null"`
	AvailableTime      string         `json:"available_time" gorm:"text;null"`
	AvailableTimeArray pq.Int32Array  `json:"available_time_array" gorm:"type:integer[];not null"`
	IsAllDay           bool           `json:"isAllDay" gorm:"bool;not null"`
	IsAllYear          bool           `json:"isAllYear" gorm:"bool;not null"`
	Location           string         `json:"location" gorm:"text;not null"`
	Rarity             string         `json:"rarity" gorm:"text;not null"`
	Shadow             string         `json:"shadow" gorm:"text;not null"`
	Price              int64          `json:"price" gorm:"int64;not null"`
	PriceCj            int64          `json:"price_cj" gorm:"int64;not null"`
	CatchPhrase        string         `json:"catch_phrase" gorm:"text;not null"`
	AltCatchPhrase     pq.StringArray `json:"alt_atch_phrase,omitempty" gorm:"type:text[];null"`
	MuseumPhrase       string         `json:"museum_phrase" gorm:"text;not null"`
	ImageURI           string         `json:"image_uri" gorm:"text;not null"`
	IconURI            string         `json:"icon_uri" gorm:"text;not null"`
}
