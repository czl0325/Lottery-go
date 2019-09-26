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
	sign := createLoginuserSign(loginUser)
	if sign != loginUser.Sign {
		log.Println("fuc_web GetLoginUser createLoginuserSign not sign", sign, loginUser.Sign)
		return nil
	}

	return loginUser
}

// 将登录的用户信息设置到cookie中
func SetLoginuser(writer http.ResponseWriter, loginuser *models.ObjLoginUser) {
	if loginuser == nil || loginuser.Uid < 1 {
		c := &http.Cookie{
			Name:   "lottery_loginUser",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(writer, c)
		return
	}
	if loginuser.Sign == "" {
		loginuser.Sign = createLoginuserSign(loginuser)
	}
	params := url.Values{}
	params.Add("uid", strconv.Itoa(loginuser.Uid))
	params.Add("username", loginuser.Username)
	params.Add("now", strconv.Itoa(loginuser.Now))
	params.Add("ip", loginuser.Ip)
	params.Add("sign", loginuser.Sign)
	c := &http.Cookie{
		Name:  "lottery_loginUser",
		Value: params.Encode(),
		Path:  "/",
	}
	http.SetCookie(writer, c)
}

// 根据登录用户信息生成加密字符串
func createLoginuserSign(loginUser *models.ObjLoginUser) string {
	str := fmt.Sprintf("uid=%d&username=%s&secret=%s", loginUser.Uid, loginUser.Username, conf.CookieSecret)
	return CreateSign(str)
}
