package grpc

import (
	"api-gateway/internal/infra/proto/actor"
	"api-gateway/internal/infra/proto/director"
	"api-gateway/internal/infra/proto/genre"
	"api-gateway/internal/infra/proto/movie"
	"fmt"
	"sync"

	"github.com/bugsnag/bugsnag-go/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"api-gateway/cmd/config"
)

var once sync.Once
var err error
var con *grpc.ClientConn

var publisherMovie movie.MovieCrudClient
var publisherDirector director.DirectorCrudClient
var publisherActor actor.ActorCrudClient
var publisherGenre genre.GenreCrudClient

func NewGrpcMovieConnection() {
	once.Do(func() {
		con, err = grpc.Dial(config.GetConfig().Domain.GrpcMovie, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			bugsnag.Notify(fmt.Errorf("error al conectarse al grpc: %s", err))
			panic("error al conectarse al grpc")
		}
		publisherMovie = movie.NewMovieCrudClient(con)
		publisherDirector = director.NewDirectorCrudClient(con)
		publisherActor = actor.NewActorCrudClient(con)
		publisherGenre = genre.NewGenreCrudClient(con)
	})
}

func GetPublisherMovie() movie.MovieCrudClient {
	return publisherMovie
}

func GetPublisherDirector() director.DirectorCrudClient {
	return publisherDirector
}

func GetPublisherActor() actor.ActorCrudClient {
	return publisherActor
}

func GetPublisherGenre() genre.GenreCrudClient {
	return publisherGenre
}
