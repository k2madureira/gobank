// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package dbbank

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID          int64              `json:"id"`
	Owner       string             `json:"owner"`
	Balance     int64              `json:"balance"`
	Currency    string             `json:"currency"`
	CreateAt    pgtype.Timestamptz `json:"create_at"`
	CountryCode pgtype.Int4        `json:"country_code"`
}

type Entry struct {
	ID        int64              `json:"id"`
	AccountID pgtype.Int8        `json:"account_id"`
	Amount    int64              `json:"amount"`
	CreateAt  pgtype.Timestamptz `json:"create_at"`
}

type Transfer struct {
	ID            int64              `json:"id"`
	FromAccountID int64              `json:"from_account_id"`
	ToAccountID   int64              `json:"to_account_id"`
	Amount        pgtype.Int8        `json:"amount"`
	CreateAt      pgtype.Timestamptz `json:"create_at"`
}
