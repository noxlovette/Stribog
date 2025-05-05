package types

import "github.com/jackc/pgx/v5/pgtype"

func ToPgText(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: s, Valid: true}
}

func PgTextToString(t pgtype.Text) string {
	if t.Valid {
		return t.String
	}
	return ""
}
