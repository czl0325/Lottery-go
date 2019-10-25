package comm

import (
	"Lottery-go/conf"
	"Lottery-go/models"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

// 得到客户端IP地址
func ClientIP(request *http.Request) string {
	host, _, _ := net.SplitHostPort(request.RemoteAddr)
	return host
}

// 跳转URL
func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location", url)
	writer.WriteHeader(http.StatusFound)
}

// 从cookie中得到当前登录的用户
func GetLoginUser(request *http.Request) *models.ObjLoginUser {
	c, err := request.Cookie("lottery_loginUser")
	if err != nil {
		return nil
	}
	params, err := url.ParseQuery(c.Value)
	if err != nil {
		return nil
	}
	uid, err := strconv.Atoi(params.Get("uid"))
	if err != nil || uid < 1 {
		return nil
	}
	// Cookie最长使用时长
	now, err := strconv.Atoi(params.Get("now"))
	if err != nil || NowUnix()-now > 86400*30 {
		return nil
	}
	//// IP修改了是不是要重新登录
	//ip := params.Get("ip")
	//if ip != ClientIP(request) {
	//	return nil
	//}
	// 登录信息
	loginUser := &models.ObjLoginUser{}
	loginUser.Uid = uid
	loginUser.Username = params.Get("username")
	loginUser.Now = now
	loginUser.Ip = ClientIP(request)
	loginUser.Sign = params.Get("sign")
	if err != nil {
		log.Println("fuc_web GetLoginUser Unmarshal ", err)
		return nil
	}
	sign := createLoginUserSign(loginUser)
	if sign != loginUser.Sign {
		log.Println("签名校验失败！", sign, loginUser.Sign)
		return nil
	}

	return loginUser
}

// 将登录的用户信息设置到cookie中
func SetLoginUser(writer http.ResponseWriter, loginUser *models.ObjLoginUser) {
	if loginUser == nil || loginUser.Uid < 1 {
		c := &http.Cookie{
			Name:   "lottery_loginUser",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(writer, c)
		return
	}
	if loginUser.Sign == "" {
		loginUser.Sign = createLoginUserSign(loginUser)
	}
	params := url.Values{}
	params.Add("uid", strconv.Itoa(loginUser.Uid))
	params.Add("username", loginUser.Username)
	params.Add("now", strconv.Itoa(loginUser.Now))
	params.Add("ip", loginUser.Ip)
	params.Add("sign", loginUser.Sign)
	c := &http.Cookie{
		Name:  "lottery_loginUser",
		Value: params.Encode(),
		Path:  "/",
	}
	http.SetCookie(writer, c)
}

// 根据登录用户信息生成加密字符串
func createLoginUserSign(loginUser *models.ObjLoginUser) string {
	str := fmt.Sprintf("uid=%d&username=%s&secret=%s", loginUser.Uid, loginUser.Username, conf.CookieSecret)
	return CreateSign(str)
}
