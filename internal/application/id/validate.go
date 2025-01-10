package id

import (
	"fmt"
	"github.com/google/uuid"
)

func Coalesce(id uuid.UUID) (uuid.UUID, error) {
	if id == uuid.Nil {
		return uuid.NewV7()
	}

	if err := Validate(id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func Validate(id uuid.UUID) error {
	if id.Version().String() == "VERSION_7" {
		return nil
	}
	return fmt.Errorf("invalid uuid version: %s", id.Version().String())
}
