package user

import (
	"context"
	"log"
	"time"

	"github.com/celtic93/auth/internal/model"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Update(ctx context.Context, user *model.User) error {
	log.Printf("repository.User.Update started. Update user with id: %d", user.ID)

	builderUpdate := sq.Update(usersTable).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, user.Name).
		Set(emailColumn, user.Email).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: user.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Print(err)
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Printf("repository.User.Update ended. Updated user with id: %d", user.ID)

	return nil
}
