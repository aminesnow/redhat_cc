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


Table: bucket
[ 0] id                                             TEXT                 null: false  primary: true   isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "QguPGwulUQWwxcyShbJKOsJlT"}



*/

// Bucket struct is a row record of the bucket table in the bucket_store database
type Bucket struct {
	//[ 0] id                                             TEXT                 null: false  primary: true   isArray: false  auto: false  col: TEXT            len: -1      default: []
	ID string `gorm:"primary_key;column:id;type:TEXT;"`
}

var bucketTableInfo = &TableInfo{
	Name: "bucket",
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
	},
}

// TableName sets the insert table name for this struct type
func (b *Bucket) TableName() string {
	return "bucket"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Bucket) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Bucket) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Bucket) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *Bucket) TableInfo() *TableInfo {
	return bucketTableInfo
}
