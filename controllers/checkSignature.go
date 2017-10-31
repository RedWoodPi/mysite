package controllers


import (
	"fmt"
	"net/http"
	"log"
	"strings"
	"sort"
	"crypto/sha1"
	"io"
	"io/ioutil"
	"time"
	"encoding/xml"
)
//定义Request结构体
type RequestBody struct {
    ToUserName   string
    FromUserName string
    CreateTime   time.Duration
    MsgType      string
    Content      string
    MsgId        int
}
type ResponseBody struct {
    ToUserName   string
    FromUserName string
    CreateTime   time.Duration
    MsgType      string
    Content      string
}

//转换加密
func str2sha1(data string) string{
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//验证Signature
func CheckSignature(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	//判断请求类型，如果为GET启动验证
	if r.Method == "GET"{
	    token := "qq258000"
	    signature := strings.Join(r.Form["signature"],"")
	    timestamp := strings.Join(r.Form["timestamp"],"")
	    nonce := strings.Join(r.Form["nonce"], "")
	    echostr := strings.Join(r.Form["echostr"], "")
	    tmps := []string{token, timestamp, nonce}
	    sort.Strings(tmps)
	    tmpstr := strings.Join(tmps, "")
	    tmp := str2sha1(tmpstr)
	    if tmp == signature {
		    fmt.Fprintf(w, echostr)
		}else{
			fmt.Fprintf(w,"")
	    	fmt.Println("Signature is not right.")
		}
	}
	if r.Method == "POST"{
		msg, err := ioutil.ReadAll(r.Body)
	    if err != nil {
		    log.Fatal(err)
	    }
	    r.Body.Close()
	    v := RequestBody{}
	    xml.Unmarshal(msg, &v)
	    if v.MsgType == "text"{
		    c := &ResponseBody{v.FromUserName, v.ToUserName,v.CreateTime,
		    v.MsgType,v.Content}
		    output, err := xml.MarshalIndent(c, "", "")
		    if err != nil {
			    fmt.Printf("error:%v\n",err)
		    }
		    outstring := strings.Replace(string(output), "ResponseBody", "xml", -1)
		    fmt.Fprintf(w, outstring)
	    }else if v.MsgType == "event" {
		    Content := `"欢迎关注"`
		    v := ResponseBody{v.ToUserName, v.FromUserName,v.CreateTime,
		    v.MsgType,Content}
		    output, err := xml.MarshalIndent(v,"","")
		    if err != nil {
			    fmt.Printf("error:%v\n",err)
		    }
		    outstring := strings.Replace(string(output), "ResponseBody", "xml", -1)
		    fmt.Fprintf(w, string(outstring))
	    }else {
		    fmt.Fprintf(w,"hello")
	    }
	}
}
