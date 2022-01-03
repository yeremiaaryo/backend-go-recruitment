package main

import (
	movieApplication "backend-go-recruitment/application/movie"
	"backend-go-recruitment/common/movie/movie_grpc"
	"backend-go-recruitment/controller/movie"
	movieInterface "backend-go-recruitment/interfaces/movie"
	movieRepository "backend-go-recruitment/repository/movie"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	RestPort = ":9000"
	GRPCPort = ":9001"
)

func main() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/local")
	if err != nil {
		log.Fatal("could not connect to database", err.Error())
	}
	defer db.Close()

	r := echo.New()

	mr := movieRepository.NewMovieRepository(db)
	ma := movieApplication.NewMovieApplication(mr)
	mi := movieInterface.NewMovieInterface(ma)
	movieController := movie.NewMovieController(mi)
	r.GET("/movie/search", movieController.FindMovieByKeyword)
	go func() {
		r.Start(RestPort)
	}()

	srv := grpc.NewServer()
	movieSrv := movie.NewMovieServer(mi)
	movie_grpc.RegisterMovieServer(srv, movieSrv)
	reflection.Register(srv)

	log.Println("Starting RPC server at", GRPCPort)
	l, err := net.Listen("tcp", GRPCPort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", GRPCPort, err)
	}
	srv.Serve(l)
}
