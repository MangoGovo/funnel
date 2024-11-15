package libraryService

import (
	"bytes"
	"funnel/app/apis/library"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"net/http"
)

// OAuthLogin 统一登陆
func OAuthLogin(username string, password string) ([]*http.Cookie, error) {
	client := resty.New()
	// 1. 初始化请求
	if _, err := client.R().
		EnableTrace().
		Get(library.OAuthChaoXingInit); err != nil {
		return nil, err
	}
	// 2. 登陆参数生成
	resp, err := client.R().
		EnableTrace().
		Get(library.LoginFromOAuth)
	if err != nil {
		return nil, err
	}

	// 解析execution
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return nil, err
	}
	execution := doc.
		Find("input[type=hidden][name=execution]").
		AttrOr("value", "")

	// 密码加密
	encPwd, err := GetEncryptedPwd(client, password)

	loginParams := map[string]string{
		"username":   username,
		"mobileCode": "",
		"password":   encPwd,
		"authcode":   "",
		"execution":  execution,
		"_eventId":   "submit",
	}
	// 3. 发送登陆请求
	resp, err = client.R().
		EnableTrace().
		SetFormData(loginParams).
		Post(library.LoginFromOAuth)
	if err != nil {
		return nil, err
	}

	// 4. 处理重定向
	// 这里我们需要手动的去处理位于js中的重定向
	// resty只能自动处理header.Location中的重定向
	redirect := GetRedirectLocation(resp.String())

	resp, err = client.R().Get(redirect)
	if err != nil {
		return nil, err
	}
	cookies := resp.Cookies()
	if !CheckCookie(cookies) {
		return nil, err
	}
	return resp.Cookies(), nil
}
