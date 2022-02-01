package access_token

import (
	"github.com/aprilnurf/grocerystore-oauth-api/src/utils/errors"
	_ "github.com/aprilnurf/grocerystore-oauth-api/src/utils/errors"
	"strings"
	"time"
)

const expiredTime = 24

type AccessToken struct {
	AccessToken string `json:"accessToken"`
	UserId      int64  `json:"userId"`
	ClientId    int64  `json:"clientId"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)

	if at.AccessToken == "" {
		return errors.NewBadRequestError("Access token can't be empty")
	}
	if at.UserId == 0 {
		return errors.NewBadRequestError("User id can't be empty")
	}
	if at.ClientId == 0 {
		return errors.NewBadRequestError("Client id can't be empty")
	}
	if at.Expires == 0 {
		return errors.NewBadRequestError("Expires id can't be empty")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expiredTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) isExpired() bool {
	//now := time.Now().UTC()
	//expiredTime := time.Unix(at.Expires, 0)
	//return now.After(expiredTime)
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
