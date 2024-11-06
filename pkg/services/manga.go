package services

import (
	"context"
	"manga-svc/pkg/db"
	"manga-svc/pkg/helper"
	"manga-svc/pkg/models"
	"manga-svc/pkg/pb"
	"net/http"
)

type MangaService struct{

  Repo db.Repository
  pb.UnimplementedMangaServiceServer

}

func (s *MangaService)CreateMange(ctx context.Context,req *pb.CreateMangaRequest)(*pb.CreateMangaResponse,error){


  var manga models.Manga
  

  manga.Name = req.Manga.Name
  manga.Author=req.Manga.Author
  manga.PublishedDate = req.Manga.PublishedDate
  manga.Genre = req.Manga.Genre
  manga.Price = req.Manga.Price
  manga.Stock = req.Manga.Stock
  manga.Description=req.Manga.Description
  manga.Rating=req.Manga.Rating


  if result := s.Repo.DB.Create(&manga); result.Error != nil {
    return &pb.CreateMangaResponse{
      Status: http.StatusConflict,
      Error:  result.Error.Error(),
    }, nil
  }

  pbManga:=helper.ConvertToPbManga(manga)

  return &pb.CreateMangaResponse{
    Status: http.StatusCreated,
    Manga: pbManga,
  }, nil 





}


