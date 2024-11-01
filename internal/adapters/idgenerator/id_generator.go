package idgenerator

import (
	"account-service/internal/adapters/uuid"
)

func NewUUIDGenerator() *uuid.Generator {
	return uuid.NewGenerator()
}
