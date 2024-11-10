package utils

import "github.com/jackc/pgx/v5/pgtype"

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
