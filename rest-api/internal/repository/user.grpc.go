package repository

import (
	"context"
	"devstream-rest-api/internal/model"
	"devstream-rest-api/internal/util/apperror"
	"devstream-rest-api/internal/util/idgen"
	"devstream-rest-api/pkg/userpb"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryGrpc struct {
	db     *sqlx.DB
	client userpb.UserServiceRpcClient
	idgen  idgen.IdGen
}

func NewUserRepositoryGrpc(DB *sqlx.DB, client userpb.UserServiceRpcClient, idgen idgen.IdGen) UserRepositoryGrpc {
	return UserRepositoryGrpc{
		db:     DB,
		client: client,
		idgen:  idgen,
	}
}

func (r UserRepositoryGrpc) Create(user *model.User) *apperror.AppError {
	user.Id = r.idgen.New()

	err := user.HashPassword()

	if err != nil {
		return apperror.Parse(err)
	}

	response, err := r.client.Create(context.Background(), &userpb.UserRequest{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})

	if !response.Status {
		return apperror.Parse(err)
	}

	if err != nil {
		return apperror.Parse(err)
	}

	user.Id = *response.Id

	return nil
}

func (r UserRepositoryGrpc) GetByEmail(email string) (*model.User, *apperror.AppError) {

	var user model.User

	if err := r.db.QueryRowx("SELECT * from users WHERE email=$1;", email).StructScan(&user); err != nil {
		return nil, apperror.Parse(err)
	}

	return &user, nil
}

func (r UserRepositoryGrpc) GetByUsername(username string) (*model.User, *apperror.AppError) {

	var user model.User

	if err := r.db.QueryRowx("SELECT * from users WHERE username=$1;", username).StructScan(&user); err != nil {
		return nil, apperror.Parse(err)
	}

	return &user, nil
}
