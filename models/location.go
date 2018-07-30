package models

import (
	_ "fmt"
	_ "github.com/google/uuid"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "time"
)

type LocaltionMarkers struct {
	Markers []struct {
		Height    int     `json:"height"`
		IconPath  string  `json:"iconPath"`
		ID        int     `json:"id"`
		Label     string  `json:"label"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Title     string  `json:"title"`
		Width     int     `json:"width"`
	} `json:"markers"`
	Circles []Circle `json:"circles"`
	// 	Latitude    float64 `json:"latitude"`
	// 	Longitude   float64 `json:"longitude"`
	// 	Color       string  `json:"color"`
	// 	Radius      int     `json:"radius"`
	// 	StorkeWidth int     `json:"storkeWidth"`
	// } `json:"circles"`
}

type Circle struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Color       string  `json:"color"`
	Radius      int     `json:"radius"`
	StorkeWidth int     `json:"storkeWidth"`
}

type Markers struct {
	Height    int     `json:"height"`
	IconPath  string  `json:"iconPath"`
	ID        int     `json:"id"`
	Label     string  `json:"label"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`
	Width     int     `json:"width"`
}
