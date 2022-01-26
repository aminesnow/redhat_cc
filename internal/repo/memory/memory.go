package memory

import (
	"sync"

	"github.com/aminesnow/redhat_cc/internal/common"
	"github.com/aminesnow/redhat_cc/internal/entity"
	"github.com/sirupsen/logrus"
)

type sotredObject struct {
	ID      string
	Content string
}

type MemoryObjectRepo struct {
	store map[string]map[string]sotredObject
	lock  *sync.RWMutex
}

func NewMemoryObjectRepo() *MemoryObjectRepo {
	store := make(map[string]map[string]sotredObject)
	lock := sync.RWMutex{}
	return &MemoryObjectRepo{
		store: store,
		lock:  &lock,
	}
}

func (mo *MemoryObjectRepo) ReadObject(bucket string, objectID string) (*entity.Object, error) {
	mo.lock.RLock()
	defer mo.lock.RUnlock()

	logrus.Debug(mo.store)

	if bStore, okB := mo.store[bucket]; okB {
		if obj, okO := bStore[objectID]; okO {
			return &entity.Object{
				ObjectID: obj.ID,
				Content:  obj.Content,
			}, nil
		}

		return nil, common.NewErrNotFoundError("object", objectID)
	}

	return nil, common.NewErrNotFoundError("bucket", bucket)
}

func (mo *MemoryObjectRepo) WriteObject(bucket string, object entity.Object) error {
	mo.lock.Lock()
	defer mo.lock.Unlock()

	sObj := sotredObject{
		ID:      object.ObjectID,
		Content: object.Content,
	}

	// create bucket if it doesn't exist
	if mo.store[bucket] == nil {
		mo.store[bucket] = make(map[string]sotredObject)
	}

	mo.store[bucket][sObj.ID] = sObj

	return nil
}

func (mo *MemoryObjectRepo) DeleteObject(bucket string, objectID string) error {
	mo.lock.Lock()
	defer mo.lock.Unlock()

	if bStore, okB := mo.store[bucket]; okB {
		if _, okO := bStore[objectID]; okO {
			delete(bStore, objectID)
			return nil
		}

		return common.NewErrNotFoundError("object", objectID)
	}

	return common.NewErrNotFoundError("bucket", bucket)
}
