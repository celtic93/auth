package user

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Delete(ctx context.Context, id int64) error {
	log.Printf("repository.User.Delete started. Delete user with id: %d", id)
	builderDelete := sq.Delete(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Print(err)
		return err
	}

	if _, err = r.db.Exec(ctx, query, args...); err != nil {
		log.Print(err)
		return err
	}

	log.Printf("repository.User.Delete ended. Deleted user with id: %d", id)

	return nil
}
