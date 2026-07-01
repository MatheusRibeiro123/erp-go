package apperrors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func TranslatePostgresError(err error) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {

		switch pgErr.Code {

		}

	}

	return err
}
