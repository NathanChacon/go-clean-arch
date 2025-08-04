package uuidGenerator

import "github.com/google/uuid"

type UuidGenerator struct {
}

func NewUuidGenerator() *UuidGenerator {
	return &UuidGenerator{}
}

func (generator *UuidGenerator) NewUuid() string {
	return uuid.New().String()
}
