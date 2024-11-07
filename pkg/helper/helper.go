package helper

import (
	"manga-svc/pkg/models"
	"manga-svc/pkg/pb"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertToPbManga(manga models.Manga) *pb.Manga {


  return &pb.Manga{
		Id:            int64(manga.ID),
		Name:          manga.Name,
		Author:        manga.Author,
		PublishedDate: ConvertToProtoTimestamp(*manga.PublishedDate),
		Genre:         manga.Genre,
		Price:         manga.Price,
		Stock:         manga.Stock,
		Description:   manga.Description,
		Rating:        manga.Rating,
	}
}


func ConvertToProtoTimestamp(t time.Time) *timestamppb.Timestamp {
    return timestamppb.New(t)
}

func ConvertFromProtoTimestamp(ts *timestamppb.Timestamp) time.Time {
    return ts.AsTime()
}
