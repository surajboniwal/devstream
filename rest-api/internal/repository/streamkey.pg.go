package repository

import (
	"devstream-rest-api/internal/model"
	"devstream-rest-api/internal/util/apperror"
	"devstream-rest-api/internal/util/idgen"

	"github.com/jmoiron/sqlx"
)

type StreamKeyRepositoryPG struct {
	db    *sqlx.DB
	idgen idgen.IdGen
}

func NewStreamKeyRepositoryPG(DB *sqlx.DB, idgen idgen.IdGen) StreamKeyRepositoryPG {
	return StreamKeyRepositoryPG{
		db:    DB,
		idgen: idgen,
	}
}

func (r StreamKeyRepositoryPG) GetByUserId(id int64) (*[]model.StreamKey, *apperror.AppError) {
	var keys = make([]model.StreamKey, 0)

	err := r.db.Select(&keys, "SELECT * FROM stream_keys WHERE user_id = $1;", id)

	if err != nil {
		return &keys, apperror.Parse(err)
	}

	return &keys, nil
}

func (r StreamKeyRepositoryPG) Create(userId int64, name string) (*model.StreamKey, *apperror.AppError) {
	var streamKey model.StreamKey

	err := r.db.Get(&streamKey, "INSERT INTO stream_keys (id, user_id, name, key) VALUES($1, $2, $3, $4) RETURNING *;", r.idgen.New(), userId, name, r.idgen.New())

	if err != nil {
		return nil, apperror.Parse(err)
	}

	return &streamKey, nil
}

func (r StreamKeyRepositoryPG) Update(id int64, userId int64, name string) (*model.StreamKey, *apperror.AppError) {
	var streamKey model.StreamKey

	err := r.db.Get(&streamKey, "UPDATE stream_keys SET name = $1 WHERE id = $2 AND user_id = $3 RETURNING *;", name, id, userId)

	if err != nil {
		return nil, apperror.Parse(err)
	}

	return &streamKey, nil
}
