package utils

import (
	"github.com/lib/pq"
)

func IsUniqueConstraintViolation(err error) bool {
    pgErr, ok := err.(*pq.Error)
    if !ok {
        return false
    }
    return pgErr.Code == "23505"
}