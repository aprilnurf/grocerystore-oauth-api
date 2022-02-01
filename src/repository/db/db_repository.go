package db

import (
	"github.com/aprilnurf/grocerystore-oauth-api/clients/cassandra"
	"github.com/aprilnurf/grocerystore-oauth-api/src/domain/access_token"
	"github.com/aprilnurf/grocerystore-oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	getAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?;"
	insertAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	updateAccessToken = "UPDATE access_tokens SET expires = ? WHERE access_token = ?;"
)

func New() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpiredTime(access_token.AccessToken) *errors.RestError
}

type dbRepository struct {
}

func (r *dbRepository) UpdateExpiredTime(at access_token.AccessToken) *errors.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	if err := session.Query(updateAccessToken, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	if err := session.Query(insertAccessToken, token.AccessToken, token.UserId, token.ClientId, token.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	//TODO: implement cassandra db
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	var result access_token.AccessToken
	if err := session.Query(getAccessToken, id).Scan(
		&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires)
		err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotExistError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}
