package user

import (
	"context"
	"log"
	"time"

	"github.com/celtic93/auth/internal/model"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	log.Printf("repository.User.Create started. User email: %s", user.Email)

	timeNow := time.Now()
	builderInsert := sq.Insert(usersTable).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn, createdAtColumn, updatedAtColumn).
		Values(user.Name, user.Email, user.Password, user.Role, timeNow, timeNow).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Print(err)
		return 0, err
	}

	var userID int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	log.Printf("repository.User.Create ended. Inserted user with id: %d", userID)

	return userID, nil
}
