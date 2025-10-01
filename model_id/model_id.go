package modelid

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/mbilarusdev/fool_base/log"
)

type ModelId struct {
	uuid.UUID
}

func (uu *ModelId) Scan(src any) error {
	value, ok := src.(string)
	if !ok || len(value) == 0 {
		return log.Err(fmt.Errorf("Invalid UUID format"), "")
	}

	parsedUUID, err := uuid.Parse(value)
	if err != nil {
		return log.Err(err, "Failed to parse UUID")
	}

	uu.UUID = parsedUUID
	return nil
}

func (uu ModelId) Value() (driver.Value, error) {
	return uu.String(), nil
}
