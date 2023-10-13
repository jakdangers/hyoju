package entity

import (
	"database/sql/driver"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        BinaryUUID `gorm:"type:binary(16);primary_key;default:(UUID_TO_BIN(UUID(),1));"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func BinaryUUIDNew() BinaryUUID {
	return BinaryUUID(uuid.New())
}

type BinaryUUID uuid.UUID

func ParseUUID(id string) (BinaryUUID, error) {
	result, err := uuid.Parse(id)
	if err != nil {
		return BinaryUUID{}, err
	}

	return BinaryUUID(result), nil
}

func (b BinaryUUID) String() string {
	return uuid.UUID(b).String()
}

func (BinaryUUID) GormDataType() string {
	return "binary(16)"
}

// Scan -> scan value into BinaryUUID
func (b *BinaryUUID) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	data, err := uuid.FromBytes(bytes)
	*b = BinaryUUID(data)
	return err
}

// Value -> return BinaryUUID to []bytes binary(16)
func (b BinaryUUID) Value() (driver.Value, error) {
	return uuid.UUID(b).MarshalBinary()
}
