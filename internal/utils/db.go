package utils

import "github.com/jackc/pgx/v5/pgtype"

// ParseUUID parses the given text and transforms it into a pgtype.UUID.
// If the given text is emtpy or the text is not a valid pgx UUID, ParseUUID returns and empty pgtype.UUID and false.
func ParseUUID(text string) (pgtype.UUID, bool) {
	if text == "" {
		return pgtype.UUID{}, false
	}
	var id pgtype.UUID
	if err := id.Scan(text); err != nil {
		return pgtype.UUID{}, false
	}
	return id, true
}
