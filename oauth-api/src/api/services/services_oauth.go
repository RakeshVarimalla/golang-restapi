package services

import (
	"microservices/api/utils/errors"
	"microservices/oauth-api/src/api/domain/oauth"
	"time"
)

type oauthService struct {
}
type oauthServiceInterface interface {
	CreateAccessToken(request *oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiErrors)
	GetAccessToken(accessToken string)(*oauth.AccessToken, errors.ApiErrors)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}
func (s *oauthService) CreateAccessToken(request *oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiErrors) {

	err := request.Validate()
	if err != nil {
		return nil, err
	}

	user, err := oauth.GetUserByUsernameAndPassword(request.Username,request.Password)
	if err != nil{
		return nil, err
	}
	token := oauth.AccessToken{
		UserId: user.Id,
		Expires: time.Now().UTC().Add(24*time.Hour).Unix(),
	}
	err1 := token.Save()
	if err1 != nil{
		return nil, err1
	}
	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string)(*oauth.AccessToken, errors.ApiErrors){
token, err := oauth.GetAccessTokenByToken(accessToken)
if err != nil{
	return nil,err
}
if token.IsExpired(){
	return nil, errors.NewNotFoundError("no access token found for given parameters")
}
return token, err
}