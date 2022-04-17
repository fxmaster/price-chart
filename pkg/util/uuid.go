package util

import uuid "github.com/satori/go.uuid"

func IsValidUUID(values ...string) bool {
	for _, v := range values {
		_, err := uuid.FromString(v)
		if err != nil {
			return false
		}
	}

	return true
}

func GenerateUUID() uuid.UUID {
	return uuid.NewV4()
}
