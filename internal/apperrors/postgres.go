package apperrors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func TranslatePostgresError(err error) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return ErrDuplicateKey
		case "23503":
			return ErrForeignKey
		}
	}
	return err
}
