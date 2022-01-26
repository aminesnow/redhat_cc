package repo

import (
	"github.com/aminesnow/redhat_cc/internal/entity"
)

type ObjectStore interface {
	ReadObject(bucket string, objectID string) (*entity.Object, error)
	WriteObject(bucket string, object entity.Object) error
	DeleteObject(bucket string, objectID string) error
}
