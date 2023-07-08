package models

import "time"

type User struct {
    ID        int       `gorm:"primaryKey"`
    Name      string
    Mobile    string
    Latitude  float64
    Longitude float64
    CreatedAt time.Time
    UpdatedAt time.Time
}
