package grpc

import (
	"fmt"
	"sync"

	"github.com/bugsnag/bugsnag-go/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"api-gateway/cmd/config"
	"api-gateway/internal/infra/proto/user"
)

var onceUser sync.Once
var connectionUser *grpc.ClientConn

var publisherUser user.UserCrudClient

func NewGrpcUserConnection() {
	onceUser.Do(func() {
		connectionUser, err = grpc.Dial(config.GetConfig().Domain.GrpcUser, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			bugsnag.Notify(fmt.Errorf("error al conectarse al grpc: %s", err))
			panic("error al conectarse al grpc")
		}

		publisherUser = user.NewUserCrudClient(connectionUser)
	})
}

func GetPublisherUser() user.UserCrudClient {
	return publisherUser
}
