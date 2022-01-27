package memory_test

import (
	"testing"

	"github.com/aminesnow/redhat_cc/internal/common"
	"github.com/aminesnow/redhat_cc/internal/entity"
	"github.com/aminesnow/redhat_cc/internal/repo/memory"
	"github.com/stretchr/testify/assert"
)

func TestWriteObject_Nominal(t *testing.T) {
	assert := assert.New(t)
	store := memory.NewMemoryObjectRepo()
	obj := entity.Object{
		ObjectID: "42",
		Content:  "foo",
	}
	bucket := "r3"

	// write object
	err := store.WriteObject(bucket, obj)
	assert.Nil(err)

	// check object inserted
	sObj, err := store.ReadObject(bucket, obj.ObjectID)
	assert.Nil(err)
	assert.NotNil(sObj)
	assert.Equal(obj.ObjectID, sObj.ObjectID)
	assert.Equal(obj.Content, sObj.Content)

	// delete object
	err = store.DeleteObject(bucket, obj.ObjectID)
	assert.Nil(err)

	// check object deleted
	sObj, err = store.ReadObject(bucket, obj.ObjectID)
	assert.NotNil(err)
	assert.Nil(sObj)
	_, ok := err.(common.ErrNotFoundError)
	assert.True(ok)
}

func TestWriteObject_NotFound(t *testing.T) {
	assert := assert.New(t)
	store := memory.NewMemoryObjectRepo()

	// try to read inexistant object
	sObj, err := store.ReadObject("bucket", "42")
	assert.Nil(sObj)
	assert.NotNil(err)
	_, ok := err.(common.ErrNotFoundError)
	assert.True(ok)

	// try to delete inexistant object
	err = store.DeleteObject("bucket", "42")
	assert.NotNil(err)
	_, ok = err.(common.ErrNotFoundError)
	assert.True(ok)
}
