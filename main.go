package main

import (
	"net/http"
	"log"
	"strings"
	"fmt"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"mysite/controllers"
)

//主程序
func main()  {
	http.HandleFunc("/check", controllers.CheckSignature)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Listenandserver: ", err)
	}

}

//获取accessToken
func accessToken(appid, secret string) (string, error) {
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
		fmt.Println("微信返回出错")
		return "", nil
	}
}

