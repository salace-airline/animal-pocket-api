package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// User informations and caught resources
type User struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name" validate:"required,min=3,max=20"`
	Email       string         `json:"email" gorm:"unique" validate:"required,email"`
	Password    []byte         `json:"-"` // json:"-" means password will not return with the response
	CaughtItems pq.StringArray `json:"caught_items" gorm:"type:text[];null"`
}

// Resources as fish, bug and sea creatures
type Fish struct {
	gorm.Model
	Name               string         `json:"name" gorm:"text;not null"`
	MonthNorthern      string         `json:"month_northern" gorm:"text;null"`
	MonthNorthernArray pq.Int32Array  `json:"month_northern_array" gorm:"type:integer[];not null"`
	MonthSouthern      string         `json:"month-southern" gorm:"text;null"`
	MonthSouthernArray pq.Int32Array  `json:"month_southern_array" gorm:"type:integer[];not null"`
	AvailableTime      string         `json:"available_time" gorm:"text;null"`
	AvailableTimeArray pq.Int32Array  `json:"available_time_array" gorm:"type:integer[];not null"`
	IsAllDay           bool           `json:"is_all_day" gorm:"bool;not null"`
	IsAllYear          bool           `json:"is_all_year" gorm:"bool;not null"`
	Location           string         `json:"location" gorm:"text;not null"`
	Rarity             string         `json:"rarity" gorm:"text;not null"`
	Shadow             string         `json:"shadow" gorm:"text;not null"`
	Price              int64          `json:"price" gorm:"int64;not null"`
	PriceCj            int64          `json:"price_cj" gorm:"int64;not null"`
	CatchPhrase        string         `json:"catch_phrase" gorm:"text;not null"`
	CatchPhraseAlt     pq.StringArray `json:"catch_phrase_alt,omitempty" gorm:"type:text[];null"`
	MuseumPhrase       string         `json:"museum_phrase" gorm:"text;not null"`
	ImageURI           string         `json:"image_uri" gorm:"text;not null"`
	IconURI            string         `json:"icon_uri" gorm:"text;not null"`
}

type Bug struct {
	gorm.Model
	Name               string         `json:"name" gorm:"text;not null"`
	MonthNorthern      string         `json:"month_northern" gorm:"text;null"`
	MonthNorthernArray pq.Int32Array  `json:"month_northern_array" gorm:"type:integer[];not null"`
	MonthSouthern      string         `json:"month-southern" gorm:"text;null"`
	MonthSouthernArray pq.Int32Array  `json:"month_southern_array" gorm:"type:integer[];not null"`
	AvailableTime      string         `json:"available_time" gorm:"text;null"`
	AvailableTimeArray pq.Int32Array  `json:"available_time_array" gorm:"type:integer[];not null"`
	IsAllDay           bool           `json:"is_all_day" gorm:"bool;not null"`
	IsAllYear          bool           `json:"is_all_year" gorm:"bool;not null"`
	Location           string         `json:"location" gorm:"text;not null"`
	Rarity             string         `json:"rarity" gorm:"text;not null"`
	Price              int64          `json:"price" gorm:"int64;not null"`
	PriceFlick         int64          `json:"price_flick" gorm:"int64;not null"`
	CatchPhrase        string         `json:"catch_phrase" gorm:"text;not null"`
	CatchPhraseAlt     pq.StringArray `json:"catch_phrase_alt,omitempty" gorm:"type:text[];null"`
	MuseumPhrase       string         `json:"museum_phrase" gorm:"text;not null"`
	ImageURI           string         `json:"image_uri" gorm:"text;not null"`
	IconURI            string         `json:"icon_uri" gorm:"text;not null"`
}

type SeaCreature struct {
	gorm.Model
	Name               string         `json:"name" gorm:"text;not null"`
	MonthNorthern      string         `json:"month_northern" gorm:"text;null"`
	MonthNorthernArray pq.Int32Array  `json:"month_northern_array" gorm:"type:integer[];not null"`
	MonthSouthern      string         `json:"month-southern" gorm:"text;null"`
	MonthSouthernArray pq.Int32Array  `json:"month_southern_array" gorm:"type:integer[];not null"`
	AvailableTime      string         `json:"available_time" gorm:"text;null"`
	AvailableTimeArray pq.Int32Array  `json:"available_time_array" gorm:"type:integer[];not null"`
	IsAllDay           bool           `json:"is_all_day" gorm:"bool;not null"`
	IsAllYear          bool           `json:"is_all_year" gorm:"bool;not null"`
	Location           string         `json:"location" gorm:"text;not null"`
	Rarity             string         `json:"rarity" gorm:"text;not null"`
	Shadow             string         `json:"shadow" gorm:"text;not null"`
	Speed              string         `json:"speed" gorm:"text;not null"`
	Price              int64          `json:"price" gorm:"int64;not null"`
	CatchPhrase        string         `json:"catch_phrase" gorm:"text;not null"`
	CatchPhraseAlt     pq.StringArray `json:"catch_phrase_alt,omitempty" gorm:"type:text[];null"`
	MuseumPhrase       string         `json:"museum_phrase" gorm:"text;not null"`
	ImageURI           string         `json:"image_uri" gorm:"text;not null"`
	IconURI            string         `json:"icon_uri" gorm:"text;not null"`
}
