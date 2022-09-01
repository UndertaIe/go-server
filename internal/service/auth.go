package service

type AuthRequest struct {
	// AppKey    string `form:"app_key" binding:"required"` // 原始post方式
	AppKey    string `json:"app_key" binding:"required"` // json格式
	AppSecret string `json:"app_secret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	return nil
}
