package wechat

import "github.com/medivhzhan/weapp/v2"

type Service struct {
	AppID  string
	Secret string
}

func (s *Service) Resolve(code string) (string, error) {
	login, err := weapp.Login(s.AppID, s.Secret, code)
	if err != nil {
		return "", err
	}
	if login.GetResponseError() != nil {
		return "", err
	}
	return login.OpenID, nil
}
