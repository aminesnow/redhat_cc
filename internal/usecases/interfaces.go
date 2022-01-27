package usecases

import "github.com/aminesnow/redhat_cc/internal/entity"

type ManageObjects interface {
	GetObject(bucketID string, objectID string) (*entity.Object, error)
	UploadObject(bucketID string, object entity.Object) error
	DeleteObject(bucketID string, objectID string) error
}
