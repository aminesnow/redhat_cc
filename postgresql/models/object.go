package models

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


Table: object
[ 0] id                                             TEXT                 null: false  primary: true   isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 1] content                                        TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] bucket_id                                      TEXT                 null: false  primary: true   isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "rdfUHluAklSNRAcsuEavZPvEk",    "content": "rGbKflCnrXcFNIgkkjNSwtYke",    "bucket_id": "cHQdpsCpAdrMuYTjxWZhrcyYi"}



*/

// Object struct is a row record of the object table in the bucket_store database
type Object struct {
	//[ 0] id                                             TEXT                 null: false  primary: true   isArray: false  auto: false  col: TEXT            len: -1      default: []
	ID string `gorm:"primary_key;column:id;type:TEXT;"`
	//[ 1] content                                        TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Content string `gorm:"column:content;type:TEXT;"`
	//[ 2] bucket_id                                      TEXT                 null: false  primary: true   isArray: false  auto: false  col: TEXT            len: -1      default: []
	BucketID string `gorm:"primary_key;column:bucket_id;type:TEXT;"`
}

var objectTableInfo = &TableInfo{
	Name: "object",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "content",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Content",
			GoFieldType:        "string",
			JSONFieldName:      "content",
			ProtobufFieldName:  "content",
			ProtobufType:       "",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "bucket_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "BucketID",
			GoFieldType:        "string",
			JSONFieldName:      "bucket_id",
			ProtobufFieldName:  "bucket_id",
			ProtobufType:       "",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *Object) TableName() string {
	return "object"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *Object) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *Object) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *Object) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *Object) TableInfo() *TableInfo {
	return objectTableInfo
}
