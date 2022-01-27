package usecases_test

import (
	"testing"

	"github.com/aminesnow/redhat_cc/internal/common"
	"github.com/aminesnow/redhat_cc/internal/entity"
	mock_repo "github.com/aminesnow/redhat_cc/internal/repo/mock"
	"github.com/aminesnow/redhat_cc/internal/usecases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestWriteObject_Nominal(t *testing.T) {
	assert := assert.New(t)
	obj := entity.Object{
		ObjectID: "42",
		Content:  "foo",
	}
	bucket := "r3"

	ctrlr := gomock.NewController(t)
	defer ctrlr.Finish()

	repo := mock_repo.NewMockObjectStore(ctrlr)
	uc := usecases.NewObjectManager(repo)

	// write object
	repo.EXPECT().WriteObject(bucket, obj).Return(nil)
	err := uc.UploadObject(bucket, obj)
	assert.Nil(err)

	// check object inserted
	repo.EXPECT().ReadObject(bucket, obj.ObjectID).Return(&obj, nil)
	sObj, err := uc.GetObject(bucket, obj.ObjectID)
	assert.Nil(err)
	assert.NotNil(sObj)
	assert.Equal(obj.ObjectID, sObj.ObjectID)
	assert.Equal(obj.Content, sObj.Content)

	// delete object
	repo.EXPECT().DeleteObject(bucket, obj.ObjectID).Return(nil)
	err = uc.DeleteObject(bucket, obj.ObjectID)
	assert.Nil(err)

	// check object deleted
	repo.EXPECT().ReadObject(bucket, obj.ObjectID).Return(nil, common.ErrNotFoundError{})
	sObj, err = uc.GetObject(bucket, obj.ObjectID)
	assert.NotNil(err)
	assert.Nil(sObj)
	_, ok := err.(common.ErrNotFoundError)
	assert.True(ok)
}

func TestWriteObject_NotFound(t *testing.T) {
	assert := assert.New(t)
	ctrlr := gomock.NewController(t)
	defer ctrlr.Finish()

	repo := mock_repo.NewMockObjectStore(ctrlr)
	uc := usecases.NewObjectManager(repo)
	bucket := "bucket"
	id := "42"

	// try to read inexistant object
	repo.EXPECT().ReadObject(bucket, id).Return(nil, common.ErrNotFoundError{})
	sObj, err := uc.GetObject(bucket, id)
	assert.Nil(sObj)
	assert.NotNil(err)
	_, ok := err.(common.ErrNotFoundError)
	assert.True(ok)

	// try to delete inexistant object
	repo.EXPECT().DeleteObject(bucket, id).Return(common.ErrNotFoundError{})
	err = uc.DeleteObject(bucket, id)
	assert.NotNil(err)
	_, ok = err.(common.ErrNotFoundError)
	assert.True(ok)
}
