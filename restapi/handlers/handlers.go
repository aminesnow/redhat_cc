package handlers

import (
	"github.com/aminesnow/redhat_cc/internal/common"
	"github.com/aminesnow/redhat_cc/internal/entity"
	"github.com/aminesnow/redhat_cc/internal/usecases"
	"github.com/aminesnow/redhat_cc/restapi/models"
	"github.com/aminesnow/redhat_cc/restapi/server/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

func UploadObjectHandler(uc usecases.ManageObjects) func(params operations.UploadObjectParams) middleware.Responder {
	return func(params operations.UploadObjectParams) middleware.Responder {
		obj := entity.Object{
			ObjectID: params.ObjectID,
			Content:  params.Object.Content,
		}

		logrus.Infof("storing object: %+v in bucket: %s", obj, params.Bucket)

		err := uc.UploadObject(params.Bucket, obj)
		if err != nil {
			logrus.Error(err.Error())

			return operations.NewDeleteObjectInternalServerError()
		}

		return operations.NewUploadObjectCreated().WithPayload(&models.ObjectID{ID: params.ObjectID})
	}
}

func GetObjectHandler(uc usecases.ManageObjects) func(params operations.GetObjectParams) middleware.Responder {
	return func(params operations.GetObjectParams) middleware.Responder {
		logrus.Infof("fetching object with id: %s in bucket: %s", params.ObjectID, params.Bucket)

		storedObj, err := uc.GetObject(params.Bucket, params.ObjectID)

		if err != nil {
			logrus.Error(err.Error())

			switch err.(type) {
			case common.ErrNotFoundError:
				return operations.NewGetObjectNotFound()
			case common.ErrInternalError:
				return operations.NewDeleteObjectInternalServerError()
			default:
				return operations.NewDeleteObjectInternalServerError()
			}
		}

		return operations.NewGetObjectOK().WithPayload(&models.Object{Content: storedObj.Content})
	}
}

func DeleteObjectHandler(uc usecases.ManageObjects) func(params operations.DeleteObjectParams) middleware.Responder {
	return func(params operations.DeleteObjectParams) middleware.Responder {
		logrus.Infof("deleteing object with id: %s in bucket: %s", params.ObjectID, params.Bucket)

		err := uc.DeleteObject(params.Bucket, params.ObjectID)
		if err != nil {
			logrus.Error(err.Error())

			switch err.(type) {
			case common.ErrNotFoundError:
				return operations.NewGetObjectNotFound()
			case common.ErrInternalError:
				return operations.NewDeleteObjectInternalServerError()
			default:
				return operations.NewDeleteObjectInternalServerError()
			}
		}

		return operations.NewDeleteObjectOK()
	}
}
