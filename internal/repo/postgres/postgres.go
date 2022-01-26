package postgres

import (
	"gorm.io/gorm"

	"github.com/aminesnow/redhat_cc/internal/entity"
)

type PostgresqlRepo struct {
	db *gorm.DB
}

func NewPostgresqlRepo(db *gorm.DB) PostgresqlRepo {
	return PostgresqlRepo{db}
}

func (pr PostgresqlRepo) ReadObject(bucket string, objectID string) (*entity.Object, error) {
	return nil, nil
}

func (pr PostgresqlRepo) WriteObject(bucket string, object entity.Object) error {
	return nil
}

func (pr PostgresqlRepo) DeleteObject(bucket string, objectID string) error {
	return nil
}
