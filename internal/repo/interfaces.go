package repo

import (
	"github.com/aminesnow/redhat_cc/internal/entity"
)

type ObjectStore interface {
	ReadObject(bucketID string, objectID string) (*entity.Object, error)
	WriteObject(bucketID string, object entity.Object) error
	DeleteObject(bucketID string, objectID string) error
}
