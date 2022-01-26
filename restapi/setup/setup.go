package setup

import (
	"net/http"

	"github.com/aminesnow/redhat_cc/internal/usecases"
	"github.com/aminesnow/redhat_cc/restapi/handlers"
	"github.com/aminesnow/redhat_cc/restapi/server"
	"github.com/aminesnow/redhat_cc/restapi/server/operations"
	openapierrors "github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/sirupsen/logrus"
)

func GetSwaggerAPI() *operations.ObjectStoreServiceAPI {
	swaggerSpec, err := loads.Embedded(server.SwaggerJSON, server.FlatSwaggerJSON)
	if err != nil {
		logrus.Fatalln(err)
	}

	openapierrors.DefaultHTTPCode = 400

	api := operations.NewObjectStoreServiceAPI(swaggerSpec)
	api.Logger = logrus.Infof

	return api
}

func SetHandlers(api *operations.ObjectStoreServiceAPI, uc usecases.ManageObjects) {
	api.UploadObjectHandler = operations.UploadObjectHandlerFunc(handlers.UploadObjectHandler(uc))
	api.GetObjectHandler = operations.GetObjectHandlerFunc(handlers.GetObjectHandler(uc))
	api.DeleteObjectHandler = operations.DeleteObjectHandlerFunc(handlers.DeleteObjectHandler(uc))
}

func SetAPIMiddleware(api *operations.ObjectStoreServiceAPI) http.Handler {
	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}
