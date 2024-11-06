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

func (s *MangaService)CreateManga(ctx context.Context,req *pb.CreateMangaRequest)(*pb.CreateMangaResponse,error){


  var manga models.Manga


  if err := s.Repo.DB.First(&manga, req.Manga.Id).Error; err == nil {
    return &pb.CreateMangaResponse{
      Status: http.StatusConflict,
      Error:  "Manga with this ID already exists",
    }, nil
  }

  manga.Name = req.Manga.Name
  manga.Author=req.Manga.Author
  publishedDate:=  helper.ConvertFromProtoTimestamp(req.Manga.PublishedDate)
  manga.PublishedDate = &publishedDate
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

func (s *MangaService)GetManga(ctx context.Context, req *pb.GetMangaRequest)(*pb.GetMangaResponse,error){


  var responseData models.Manga

  if result := s.Repo.DB.Find(&responseData, req.Id); result.Error != nil {
    return &pb.GetMangaResponse{
      Status: http.StatusNotFound,
      Error:  result.Error.Error(),
    }, nil
  }

  pbResponseData:=helper.ConvertToPbManga(responseData)

  return &pb.GetMangaResponse{
    Status: http.StatusOK,
    Manga: pbResponseData,
  },nil

}


func (s *MangaService) ListManga(ctx context.Context,req *pb.ListMangaRequest) (*pb.ListMangaResponse, error){

  var responseDatas []models.Manga


  if result:=  s.Repo.DB.Find(&responseDatas); result.Error !=nil{

    return &pb.ListMangaResponse{
      Status: http.StatusNotFound,
      Error: result.Error.Error(),
    },nil

  }

  var pbResponseDatas []*pb.Manga
  for _,data:=range  responseDatas{
    pbData:=helper.ConvertToPbManga(data)
    pbResponseDatas=append(pbResponseDatas, pbData)
  }

  return &pb.ListMangaResponse{
    Status: http.StatusOK,
    Manga: pbResponseDatas,
  },nil


}


func (s *MangaService) UpdateManga(ctx context.Context, req *pb.UpdateMangaRequest) (*pb.UpdateMangaResponse, error) {
  var manga models.Manga

  if err := s.Repo.DB.First(&manga, req.Updatemodel.Id).Error; err != nil {
    return &pb.UpdateMangaResponse{
      Status: http.StatusNotFound,
      Error:  "Manga not found with this ID",
    }, nil
  }

  updateFields := make(map[string]any)

  if req.Updatemodel.Name != "" {
    updateFields["name"] = req.Updatemodel.Name
  }
  if req.Updatemodel.Author != "" {
    updateFields["author"] = req.Updatemodel.Author
  }
  if req.Updatemodel.Genre != "" {
    updateFields["genre"] = req.Updatemodel.Genre
  }
  if req.Updatemodel.Price > 0 {
    updateFields["price"] = req.Updatemodel.Price
  }
  if req.Updatemodel.Stock > 0 {
    updateFields["stock"] = req.Updatemodel.Stock
  }
  if req.Updatemodel.Description != "" {
    updateFields["description"] = req.Updatemodel.Description
  }
  if req.Updatemodel.PublishedDate != nil {
    updateFields["published_date"] = req.Updatemodel.PublishedDate.AsTime() 
  }

  if err := s.Repo.DB.Model(&manga).Updates(updateFields).Error; err != nil {
    return &pb.UpdateMangaResponse{
      Status: http.StatusInternalServerError,
      Error:  err.Error(),
    }, nil
  }

  return &pb.UpdateMangaResponse{
    Status: http.StatusOK,
  }, nil
}


func (s *MangaService) DeleteManga(ctx context.Context, req *pb.DeleteMangaRequest) (*pb.DeleteMangaResponse, error) {
  var manga models.Manga

  if err := s.Repo.DB.First(&manga, req.Id).Error; err != nil {
    return &pb.DeleteMangaResponse{

      Status: http.StatusNotFound,
      Success: false,
      Error:  "Manga not found with the provided ID",
    }, nil
  }

  if err := s.Repo.DB.Delete(&manga).Error; err != nil {
    return &pb.DeleteMangaResponse{
      Status: http.StatusInternalServerError,
      Success: false,
      Error:  err.Error(),
    }, nil
  }

  return &pb.DeleteMangaResponse{
    Status:  http.StatusOK,
    Success: true,
  }, nil
}
