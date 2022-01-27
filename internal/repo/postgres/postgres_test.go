package postgres_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aminesnow/redhat_cc/internal/common"
	"github.com/aminesnow/redhat_cc/internal/entity"
	"github.com/aminesnow/redhat_cc/internal/repo/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	gorm_pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	id     = "42"
	bucket = "r3"
	expObj = entity.Object{
		ObjectID: id,
		Content:  "foo bar",
	}
)

/*////////////////////////////////////
////////// ReadObject Tests //////////
////////////////////////////////////*/

func (s *TestSuite) TestReadObject_Nomnial() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "object" WHERE "object"."id" = $1 AND "object"."bucket_id" = $2 ORDER BY "object"."id" LIMIT 1`)).
		WithArgs(expObj.ObjectID, bucket).
		WillReturnRows(sqlmock.NewRows([]string{"id", "bucket_id", "content"}).
			AddRow(expObj.ObjectID, bucket, expObj.Content))

	obj, err := s.repo.ReadObject(bucket, expObj.ObjectID)

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), obj)
	assert.Equal(s.T(), expObj, *obj)
}

func (s *TestSuite) TestReadObject_NotFound() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "object" WHERE "object"."id" = $1 AND "object"."bucket_id" = $2 ORDER BY "object"."id" LIMIT 1`)).
		WithArgs(id, bucket).
		WillReturnError(gorm.ErrRecordNotFound)

	obj, err := s.repo.ReadObject(bucket, id)

	assert.Nil(s.T(), obj)
	assert.NotNil(s.T(), err)
	_, ok := err.(common.ErrNotFoundError)
	assert.True(s.T(), ok)
}

func (s *TestSuite) TestReadObject_InternalError() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "object" WHERE "object"."id" = $1 AND "object"."bucket_id" = $2 ORDER BY "object"."id" LIMIT 1`)).
		WithArgs(id, bucket).
		WillReturnError(gorm.ErrInvalidData)

	obj, err := s.repo.ReadObject(bucket, id)

	assert.Nil(s.T(), obj)
	assert.NotNil(s.T(), err)
	_, ok := err.(common.ErrInternalError)
	assert.True(s.T(), ok)
}

/*//////////////////////////////////////
////////// DeleteObject Tests //////////
//////////////////////////////////////*/

func (s *TestSuite) TestDeleteObject_Nomnial() {
	s.mock.ExpectExec(regexp.QuoteMeta(
		`DELETE FROM "object" WHERE ("object"."id","object"."bucket_id") IN (($1,$2))`)).
		WithArgs(id, bucket).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := s.repo.DeleteObject(bucket, id)
	assert.Nil(s.T(), err)
}

func (s *TestSuite) TestDeleteObject_NotFound() {
	s.mock.ExpectExec(regexp.QuoteMeta(
		`DELETE FROM "object" WHERE ("object"."id","object"."bucket_id") IN (($1,$2))`)).
		WithArgs(id, bucket).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := s.repo.DeleteObject(bucket, id)

	assert.NotNil(s.T(), err)
	_, ok := err.(common.ErrNotFoundError)
	assert.True(s.T(), ok)
}

func (s *TestSuite) TestDeleteObject_InternalError() {
	s.mock.ExpectExec(regexp.QuoteMeta(
		`DELETE FROM "object" WHERE ("object"."id","object"."bucket_id") IN (($1,$2))`)).
		WithArgs(id, bucket).
		WillReturnError(gorm.ErrInvalidData)

	err := s.repo.DeleteObject(bucket, id)

	assert.NotNil(s.T(), err)
	_, ok := err.(common.ErrInternalError)
	assert.True(s.T(), ok)
}

/*//////////////////////////////////////
////////// WriteObject Tests //////////
//////////////////////////////////////*/

func (s *TestSuite) TestWriteObject_Nomnial() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "bucket" WHERE "bucket"."id" = $1`)).
		WithArgs(bucket).WillReturnRows(sqlmock.NewRows([]string{"count"}).
		AddRow(1))

	s.mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "object" ("id","content","bucket_id") VALUES ($1,$2,$3)`)).
		WithArgs(expObj.ObjectID, expObj.Content, bucket).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repo.WriteObject(bucket, expObj)
	assert.Nil(s.T(), err)

	// insert second time to override exisitng
	newContent := "new content"
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "bucket" WHERE "bucket"."id" = $1`)).
		WithArgs(bucket).WillReturnRows(sqlmock.NewRows([]string{"count"}).
		AddRow(1))

	s.mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "object" ("id","content","bucket_id") VALUES ($1,$2,$3) ON CONFLICT ("id","bucket_id") DO UPDATE SET "content"="excluded"."content"`)).
		WithArgs(expObj.ObjectID, newContent, bucket).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = s.repo.WriteObject(bucket, entity.Object{
		ObjectID: expObj.ObjectID,
		Content:  newContent,
	})
	assert.Nil(s.T(), err)
}

func (s *TestSuite) TestWriteObject_NewBucket() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "bucket" WHERE "bucket"."id" = $1`)).
		WithArgs(bucket).WillReturnRows(sqlmock.NewRows([]string{"count"}).
		AddRow(0))

	s.mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "bucket" ("id") VALUES ($1)`)).
		WithArgs(bucket).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "object" ("id","content","bucket_id") VALUES ($1,$2,$3)`)).
		WithArgs(expObj.ObjectID, expObj.Content, bucket).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repo.WriteObject(bucket, expObj)

	assert.Nil(s.T(), err)
}

func (s *TestSuite) TestWriteObject_InternalError() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT count(*) FROM "bucket" WHERE "bucket"."id" = $1`)).
		WithArgs(bucket).
		WillReturnError(gorm.ErrInvalidData)

	err := s.repo.WriteObject(bucket, expObj)

	assert.NotNil(s.T(), err)
	_, ok := err.(common.ErrInternalError)
	assert.True(s.T(), ok)
}

/*///////////////////////////////
////////// Suite Setup //////////
///////////////////////////////*/

type TestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock

	repo postgres.PostgresqlRepo
}

func (s *TestSuite) SetupSuite() {
	assert := assert.New(s.T())
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	assert.Nil(err)

	s.db, err = gorm.Open(gorm_pg.New(gorm_pg.Config{
		Conn: db,
	}), &gorm.Config{SkipDefaultTransaction: true})
	assert.Nil(err)

	s.repo = postgres.NewPostgresqlRepo(s.db)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
