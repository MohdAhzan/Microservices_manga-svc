package helper

import (
	"manga-svc/pkg/models"
	"manga-svc/pkg/pb"
)

func ConvertToPbManga(manga models.Manga) *pb.Manga {
	return &pb.Manga{
		Id:            int64(manga.ID),
		Name:          manga.Name,
		Author:        manga.Author,
		PublishedDate: manga.PublishedDate,
		Genre:         manga.Genre,
		Price:         manga.Price,
		Stock:         manga.Stock,
		Description:   manga.Description,
		Rating:        manga.Rating,
	}
}
