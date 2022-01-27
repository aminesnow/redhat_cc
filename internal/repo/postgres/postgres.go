package postgres

import (
	"errors"

	"github.com/aminesnow/redhat_cc/internal/common"
	"github.com/aminesnow/redhat_cc/internal/entity"
	"github.com/aminesnow/redhat_cc/postgresql/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresqlRepo struct {
	db *gorm.DB
}

func NewPostgresqlRepo(db *gorm.DB) PostgresqlRepo {
	return PostgresqlRepo{db}
}

func (pr PostgresqlRepo) ReadObject(bucketID string, objectID string) (*entity.Object, error) {
	var obj models.Object

	res := pr.db.Where(&models.Object{ID: objectID, BucketID: bucketID}).First(&obj)

	if res.Error != nil {
		logrus.Error(res.Error)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, common.NewErrNotFoundErrorMsg("object with id: %s in bucket: %s was not found", objectID, bucketID)
		}

		return nil, common.NewErrInternalError(res.Error,
			"error fetching object with id: %s in bucket: %s from postgresql", objectID, bucketID)
	}

	return &entity.Object{
		ObjectID: obj.ID,
		Content:  obj.Content,
	}, nil
}

func (pr PostgresqlRepo) WriteObject(bucketID string, object entity.Object) error {
	// check if bucket exists
	var bCount int64
	res := pr.db.Model(&models.Bucket{}).Where(&models.Bucket{ID: bucketID}).Count(&bCount)

	if res.Error != nil {
		logrus.Error(res.Error)
		return common.NewErrInternalError(res.Error,
			"error fetching object with id: %s in bucket: %s from postgresql", object.ObjectID, bucketID)
	}

	// if bucket does not exist, create it
	if bCount == 0 {
		res = pr.db.Create(&models.Bucket{ID: bucketID})
		if res.Error != nil {
			logrus.Error(res.Error)
			return common.NewErrInternalError(res.Error,
				"error creating bucket: %s", bucketID)
		}
	}

	res = pr.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&models.Object{
		ID:       object.ObjectID,
		Content:  object.Content,
		BucketID: bucketID,
	})

	if res.Error != nil {
		logrus.Error(res.Error)
		return common.NewErrInternalError(res.Error,
			"error creating object with id: %s in bucket: %s", object.ObjectID, bucketID)
	}

	return nil
}

func (pr PostgresqlRepo) DeleteObject(bucketID string, objectID string) error {
	res := pr.db.Delete(models.Object{ID: objectID, BucketID: bucketID})

	if res.Error != nil {
		return common.NewErrInternalError(res.Error,
			"error fetching object with id: %s in bucket: %s from postgresql", objectID, bucketID)
	}

	if res.RowsAffected == 0 {
		return common.NewErrNotFoundErrorMsg("object with id: %s in bucket: %s was not found", objectID, bucketID)
	}

	return nil
}
