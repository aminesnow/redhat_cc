package usecases

import (
	"github.com/aminesnow/redhat_cc/internal/entity"
	"github.com/aminesnow/redhat_cc/internal/repo"
)

type ObjectManager struct {
	repo repo.ObjectStore
}

func NewObjectManager(repo repo.ObjectStore) ObjectManager {
	return ObjectManager{repo}
}

func (om ObjectManager) GetObject(bucket string, objectID string) (*entity.Object, error) {
	return om.repo.ReadObject(bucket, objectID)
}

func (om ObjectManager) UploadObject(bucket string, object entity.Object) error {
	return om.repo.WriteObject(bucket, object)
}

func (om ObjectManager) DeleteObject(bucket string, objectID string) error {
	return om.repo.DeleteObject(bucket, objectID)
}
