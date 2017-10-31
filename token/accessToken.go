package token

import (
	"strings"
	"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"database/sql"
	_"github.com/Go-SQL-Driver/MySQL"
)

const appid = "wxccc67a998a00e936"
const secret = "rmE9eObzaupBLM5aQ8xkGYpCVsuFYZA6IhmG9IEQ7sE"


//获取accessToken
func AccessToken(appid string, secret string) (string, error) {

	type AccessTokenResponse struct{
		AccessToken string   `json:"access_token"`
		ExpireIn    float64  `json:"expire_in"`
	}

	//appid := "wxccc67a998a00e936"
	//secret := "rmE9eObzaupBLM5aQ8xkGYpCVsuFYZA6IhmG9IEQ7sE"
	url := strings.Join([]string{"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=",
		appid,
			"&secret=",
				secret},"")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取accessToken错误", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if bytes.Contains(body, []byte("access_token")){
		atr := AccessTokenResponse{}
		err = json.Unmarshal(body, &atr)

		if err != nil {
			fmt.Println("解析返回json失败")
			return "", err
		}
		return atr.AccessToken, nil
	}else{
		fmt.Println("微信返回错误")
		return "", nil
	}
}

//存储token
func SaveToken() {
	fmt.Println("开始获取token")
	token, err := AccessToken(appid, secret)
	if err != nil {
		fmt.Println(err)
	}

	db ,err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mysite?charset=utf8")

	if err != nil {
		fmt.Println(err)
	}

	result, err := db.Exec("INSERT INTO accessToken (accessToken) VALUES (?)", token)

	if err != nil {
		fmt.Println(err)
	}
	_ = result

	db.Close()
}
