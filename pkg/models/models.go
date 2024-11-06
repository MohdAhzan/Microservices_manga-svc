package models

type Manga struct{

  ID             int64 `json:"id" gorm:"primaryKey"` 
  Name           string  `json:"name"`
  Author         string    `json:"author"`
  PublishedDate  string `json:"published_date"` 
  Genre          string  `json:"genre"`
  Price          float64  `json:"price"`
  Stock          int64  `json:"stock"`
  Description    string  `json:"description"`
  Rating         float64 `json:"rating"`
}


