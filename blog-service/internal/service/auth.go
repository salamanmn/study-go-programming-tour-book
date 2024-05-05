package service

import "errors"

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.Model != nil && auth.ID > 0 {
		return nil
	}
	//错误情况：返回auth中的model为空时出错
	//if auth.ID > 0 {
	//	return nil
	//}
	return errors.New("auth info does not exist.")
}
