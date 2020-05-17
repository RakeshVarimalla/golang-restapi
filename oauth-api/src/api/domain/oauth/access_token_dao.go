package oauth

import (
	"fmt"
	"microservices/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() errors.ApiErrors {
	//at.AccessToken = "234234"
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserId)
	tokens[at.AccessToken] = at
	return nil
}
func GetAccessTokenByToken(accesstoken string) (*AccessToken, errors.ApiErrors) {
	token := tokens[accesstoken]
	if token == nil {
		return nil, errors.NewNotFoundError("no access token found for given parameters")
	}
	return token, nil
}
