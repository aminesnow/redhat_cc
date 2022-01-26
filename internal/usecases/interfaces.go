package usecases

import "github.com/aminesnow/redhat_cc/internal/entity"

type ManageObjects interface {
	GetObject(bucket string, objectID string) (*entity.Object, error)
	UploadObject(bucket string, object entity.Object) error
	DeleteObject(bucket string, objectID string) error
}
