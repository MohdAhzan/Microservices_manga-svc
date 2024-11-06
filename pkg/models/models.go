package models

import "time"

type Manga struct{

  ID             int64 `json:"id" gorm:"primaryKey"` 
  Name           string  `json:"name"`
  Author         string    `json:"author"`
  PublishedDate  *time.Time `json:"published_date"` 
  Genre          string  `json:"genre"`
  Price          float64  `json:"price"`
  Stock          int64  `json:"stock"`
  Description    string  `json:"description"`
  Rating         float64 `json:"rating"`
}

type UpdateManga struct {
  Name           string  
  Author         string 
  PublishedDate  string
  Genre          string
  Price          float64 
  Stock          int64  
  Description    string  
}
