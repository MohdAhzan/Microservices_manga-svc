package main

import (
	"log"
	"manga-svc/pkg/config"
	"manga-svc/pkg/db"
	"manga-svc/pkg/pb"
	"manga-svc/pkg/services"
	"net"

	"google.golang.org/grpc"
)


func main(){
  

 cfg, err:=config.LoadConfig()
  if err!=nil{
    log.Fatal("error loading authCOnfig",err)
  }
  dbrepo,err:= db.DBconnect(cfg)
  
  if err!=nil{
    log.Fatal("error loading authCOnfig",err)
  }

  listner,err:=net.Listen("tcp",cfg.Port)
  if err!=nil{
    log.Fatal("error listening to port",cfg.DBPort,err)
  }
  
  manga_svc:=&services.MangaService{
  Repo: dbrepo,
  }
  
  grpcServer:=grpc.NewServer()
  
    pb.RegisterMangaServiceServer(grpcServer,manga_svc)

   log.Println("mangaservice running on port ",cfg.Port) 

  err=grpcServer.Serve(listner)
  if err!=nil{
    log.Fatal(err)
  }
    
}



