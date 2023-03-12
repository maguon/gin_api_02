package response

type SysCaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	Img       string `json:"img"`
}
