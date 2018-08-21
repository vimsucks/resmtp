package model

import (
	sq "github.com/Masterminds/squirrel"
)

type Dialer struct {
	ID          uint `hash:"ignore" json:"id"`
	Host        string
	Port        int
	Username    string
	Password    string
	FromName    string `json:"from_name"`
	HashSalt    byte   `db:"hash_salt"`
	AccessToken string `hash:"ignore" db:"access_token" json:"access_token"`
}

func CreateDialer(dl *Dialer) error {
	sql, args, err := sq.Insert("dialer").
		Columns("host", "port", "username", "password", "hash_salt", "access_token").
		Values(dl.Host, dl.Port, dl.Username, dl.Password, dl.HashSalt, dl.AccessToken).
		ToSql()
	if err != nil {
		return err
	}

	result, err := DB.Exec(sql, args...)
	if err != nil {
		return err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	dl.ID = uint(lastId)
	return nil
}

func GetDialerByAccessToken(accessToken string) (*Dialer, error) {
	sql, args, err := sq.Select("id", "host", "port", "username", "password", "hash_salt", "access_token").
		From("dialer").
		Where(sq.Eq{"access_token": accessToken}).
		ToSql()
	if err != nil {
		return nil, err
	}

	dl := Dialer{}
	err = DB.Get(&dl, sql, args...)
	if err != nil {
		return nil, err
	}

	return &dl, nil
}

func IfDialerAccessTokenExists(accessToken string) (bool, error) {
	sql, args, err := sq.Select("COUNT(access_token)").
		From("dialer").
		Where(sq.Eq{"access_token": accessToken}).
		ToSql()
	if err != nil {
		return false, err
	}

	var count []int
	err = DB.Select(&count, sql, args...)
	if err != nil {
		return false, err
	}

	return count[0] != 0, nil
}
