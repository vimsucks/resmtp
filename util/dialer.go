package util

import (
	"github.com/mitchellh/hashstructure"
	"strconv"
	"github.com/vimsucks/resmtp/model"
	"math/rand"
)

func GetUniqueDialerAccessToken(dl *model.Dialer) (accessToken string, err error) {
	hashUInt, err := hashstructure.Hash(*dl, nil)
	if err != nil {
		return "", err
	}
	exists, err := model.IfDialerAccessTokenExists(strconv.FormatUint(hashUInt, 10))
	for exists == true && err == nil {
		dl.HashSalt = byte(rand.Int())
		hashUInt, err = hashstructure.Hash(dl, nil)
		if err != nil {
			return "", err
		}
		exists, err = model.IfDialerAccessTokenExists(strconv.FormatUint(hashUInt, 10))
	}
	return strconv.FormatUint(hashUInt, 10), nil
}
