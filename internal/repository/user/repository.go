package user

import (
	"context"
	"log"
	"time"

	"github.com/celtic93/auth/internal/model"
	"github.com/celtic93/auth/internal/repository"
	"github.com/celtic93/auth/internal/repository/user/converter"
	"github.com/jackc/pgx/v4/pgxpool"

	sq "github.com/Masterminds/squirrel"
	modelRepo "github.com/celtic93/auth/internal/repository/user/model"
)

const (
	usersTable string = "users"

	idColumn        string = "id"
	nameColumn      string = "name"
	emailColumn     string = "email"
	passwordColumn  string = "password"
	roleColumn      string = "role"
	createdAtColumn string = "created_at"
	updatedAtColumn string = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	log.Printf("repository.User.Get User id: %d", id)
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

	log.Printf("get user id: %d, name: %s, email: %s, role: %d, created_at: %v, updated_at: %v\n",
		user.ID, user.Name, user.Email, user.Role, user.CreatedAt, user.UpdatedAt)

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	log.Printf("repository.User.Create User email: %s", user.Email)

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

	log.Printf("inserted user with id: %d", userID)

	return userID, nil
}

func (r *repo) Update(ctx context.Context, user *model.User) error {
	log.Printf("repository.User.Update User id: %d", user.ID)

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

	log.Printf("updated user with id: %d", user.ID)

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	log.Printf("server.Delete User id: %d", id)
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

	log.Printf("deleted user with id: %d", id)

	return nil
}
