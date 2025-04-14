package user

import (
	"context"
	"log"

	"github.com/celtic93/auth/internal/model"
	"github.com/celtic93/auth/internal/repository/user/converter"

	sq "github.com/Masterminds/squirrel"
	modelRepo "github.com/celtic93/auth/internal/repository/user/model"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	log.Printf("repository.User.Get started. Get user with id: %d", id)
	builderSelectOne := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var user modelRepo.User

	err = r.db.QueryRow(ctx, query, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	log.Printf("repository.User.Get ended. Got user id: %d", user.ID)

	return converter.ToUserFromRepo(&user), nil
}
