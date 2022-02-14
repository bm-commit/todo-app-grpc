package storage

import (
	"strings"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
)

// UUID is used to wrap other unique identifiers.
type UUID struct {
	mssql.UniqueIdentifier
}

// NewUUID returns a parsed uuid with the provided string.
func NewUUID(id string) UUID {
	var u UUID

	// mssql.UniqueIdentifier has some problems when uuid is invalid
	// and causes panics or errors. To handle this problem,
	// we use the Google uuid library to check if an uuid is valid.
	// If the uuid is not valid it returns an empty uuid.
	if _, err := uuid.Parse(id); err != nil {
		return u
	}
	if err := u.Scan(id); err != nil {
		return u
	}

	return u
}

// String returns an uuid with lower case.
func (u UUID) String() string {
	return strings.ToLower(u.UniqueIdentifier.String())
}
