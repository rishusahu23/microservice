//go:build wireinject
// +build wireinject

package wire

import (
	"crypto/tls"
	"github.com/google/wire"
	"github.com/rishu/microservice/config"
	"github.com/rishu/microservice/external/ohttp"
	"github.com/rishu/microservice/external/post"
	mongo2 "github.com/rishu/microservice/pkg/transaction/mongo"
	"github.com/rishu/microservice/user"
	mongoDao "github.com/rishu/microservice/user/dao/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func GetPostClientProvider(provider *post.ClientImpl) post.Client {
	return provider
}

func getHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func InitialiseUserService(conf *config.Config, mongoClient *mongo.Client) *user.Service {
	wire.Build(
		user.NewService,
		mongoDao.UserDaoWireSet,
		mongo2.DefaultTxnManagerWireSet,
		GetPostClientProvider,
		post.NewPostClientImpl,
		ohttp.NewHttpRequestHandler,
		getHttpClient,
	)
	return &user.Service{}
}
