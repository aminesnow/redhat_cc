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

func (om ObjectManager) GetObject(bucketID string, objectID string) (*entity.Object, error) {
	return om.repo.ReadObject(bucketID, objectID)
}

func (om ObjectManager) UploadObject(bucketID string, object entity.Object) error {
	return om.repo.WriteObject(bucketID, object)
}

func (om ObjectManager) DeleteObject(bucketID string, objectID string) error {
	return om.repo.DeleteObject(bucketID, objectID)
}
