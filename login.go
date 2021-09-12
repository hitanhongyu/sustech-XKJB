package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)
var(
	cookies []*http.Cookie
	header http.Header
)
func log(username,passwd string){
	var OpJar *cookiejar.Jar
	OpJar, _ = cookiejar.New(nil)
	cli:=http.Client{}
	cli.Jar = OpJar
	cli.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		fmt.Println(req.URL)
		return nil
	}
	//GET JSESSIONID route
	resp,_:=cli.Get("https://tis.sustech.edu.cn/session/invalid")
	//fmt.Println(resp)
	/*for _,v:=range resp.Cookies(){
		if v.Name=="JSESSIONID"||v.Name=="route"{
			cookies = append(cookies,&http.Cookie{Name: v.Name,Value: v.Value})
		}
	}
	*/
	req,_:= http.NewRequest("GET","https://tis.sustech.edu.cn/cas",nil)
	/*
	req.AddCookie(cookies[0])
	req.AddCookie(cookies[1])
	*/
	resp,_ = cli.Do(req)

	//Get the execution
	body,_:=io.ReadAll(resp.Body)
	reg,_:=regexp.Compile("<input type=\"hidden\" name=\"execution\" value=\".*\"/>")
	params:=reg.FindStringSubmatch(string(body))
	if len(params)==0{
		fmt.Println("GET execution error!")
		return
	}
	execution:=params[0][45:len(params[0])-97]

	//get DISSESSION
	req,_=http.NewRequest("GET","https://cas.sustech.edu.cn/cas/clientredirect?client_name=Wework&service=https%3A%2F%2Ftis.sustech.edu.cn%2Fcas",nil)
	resp,_=cli.Do(req)
	/*
	for _,v:=range resp.Cookies(){
		if v.Name=="DISSESSION"{
			cookies = append(cookies,&http.Cookie{Name: v.Name,Value: v.Value})
		}
	}
    */
	//GET TGC
	urls:=url.Values{}
	urls.Add("execution",execution)
	urls.Add("username",username)
	urls.Add("password",passwd)
	urls.Add("_eventId","submit")
	urls.Add("geolocation","")
	req,_= http.NewRequest("POST","https://cas.sustech.edu.cn/cas/login?service=https%3A%2F%2Ftis.sustech.edu.cn%2Fcas",strings.NewReader(urls.Encode()))
	//req.AddCookie(getCookie("DISSESSION"))
	req.Header.Add("Content-Length",strconv.Itoa(len([]byte(urls.Encode()))))
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
	//req.Header.Add("accept-encoding","deflate")
	//req.Header.Add("accept-language","zh-CN,zh;q=0.9")
	req.Header.Add("rolecode","02")
	//req.Header.Add("sec-ch-ua","\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	//req.Header.Add("sec-ch-ua-mobile","?0")
	//req.Header.Add("sec-fetch-dest","empty")
	//req.Header.Add("sec-fetch-mode","cors")
	//req.Header.Add("sec-fetch-site","same-origin")
	//req.Header.Add("user-agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 ")
	//req.Header.Add("x-requested-with","XMLHttpRequest")
	//req.Header.Add("content-length",strconv.Itoa(len([]byte(urls.Encode()))))
	//req.Header.Add("content-type", "application/json;charset=UTF-8")
	//req.Header.Add("origin","https://tis.sustech.edu.cn")
	//req.Header.Add("refer","https://tis.sustech.edu.cn/Xsxk/query/2")
	resp,_=cli.Do(req)
	/*for _,v:=range resp.Cookies() {
		if v.Name == "TGC" {
			cookies = append(cookies, &http.Cookie{Name: v.Name, Value: v.Value})
		}
	}
    */
	//GET https://tis.sustech.edu.cn/student_index
	req,_= http.NewRequest("GET","https://tis.sustech.edu.cn/student_index",strings.NewReader(urls.Encode()))
	//req.AddCookie(getCookie("route"))
	//req.AddCookie(getCookie("JSESSIONID"))
	resp,_ = cli.Do(req)
	//GET https://tis.sustech.edu.cn/Xsxk/query/2
	req,_= http.NewRequest("GET","https://tis.sustech.edu.cn/Xsxk/query/2",strings.NewReader(urls.Encode()))
	//req.AddCookie(getCookie("route"))
	//req.AddCookie(getCookie("JSESSIONID"))
	resp,_ = cli.Do(req)

	//Xk
	urls = url.Values{}

	urls.Add("p_pylx","2")
	urls.Add("mxpylx","2")
	urls.Add("p_sfgldjr","0")
	urls.Add("p_sfredis","0")
	urls.Add("p_sfsyxkgwc","0")
	urls.Add("p_xktjz","rwtjzyx")
	urls.Add("p_chaxunxh","")
	urls.Add("p_gjz","")
	urls.Add("p_skjs","")
	urls.Add("p_xn","2021-2022")
	urls.Add("p_xq","1")
	urls.Add("p_xnxq","2021-20221")
	urls.Add("p_dqxn","2021-2022")
	urls.Add("p_dqxq","1")
	urls.Add("p_dqxnxq","2021-20221")
	urls.Add("p_xkfsdm","jhnxk")
	urls.Add("p_xiaoqu","")
	urls.Add("p_kkyx","")
	urls.Add("p_kclb","")
	urls.Add("p_xkxs","")
	urls.Add("p_id","C2A9EFD22D070A6DE053CA0412ACBCCD")
	urls.Add("p_sfhlctkc","0")
	urls.Add("p_sfhllrlkc","0")
	urls.Add("p_kxsj_xqj","")
	urls.Add("p_kxsj_ksjc","")
	urls.Add("p_kxsj_jsjc","")
	urls.Add("p_kcdm_js","")
	urls.Add("p_kcdm_cxrw","")
	urls.Add("p_kc_gjz","")
	urls.Add("p_xzcxtjz_nj","")
	urls.Add("p_xzcxtjz_yx","")
	urls.Add("p_xzcxtjz_zy","")
	urls.Add("p_xzcxtjz_zyfx","")
	urls.Add("p_xzcxtjz_bj","")
	urls.Add("p_sfxsgwckb","1")
	urls.Add("p_skyy","")
	urls.Add("p_chaxunxkfsdm","")
	urls.Add("pageNum","1")
	urls.Add("pageSize","12")


	req,_=http.NewRequest("POST","https://tis.sustech.edu.cn/Xsxk/addGouwuche",strings.NewReader(urls.Encode()))
	//req.AddCookie(getCookie("JSESSIONID"))
	//req.AddCookie(getCookie("route"))
	//req.Header.Add("accept","*/*")
	//req.Header.Add("accept-encoding","deflate")
	//req.Header.Add("accept-language","zh-CN,zh;q=0.9")
	req.Header.Add("rolecode","02")
	//req.Header.Add("sec-ch-ua","\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	//req.Header.Add("sec-ch-ua-mobile","?0")
	//req.Header.Add("sec-fetch-dest","empty")
	//req.Header.Add("sec-fetch-mode","cors")
	//req.Header.Add("sec-fetch-site","same-origin")
	//req.Header.Add("user-agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")
	req.Header.Add("x-requested-with","XMLHttpRequest")
	req.Header.Add("content-length",strconv.Itoa(len([]byte(urls.Encode()))))
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	//req.Header.Add("origin","https://tis.sustech.edu.cn")
	//req.Header.Add("refer","https://tis.sustech.edu.cn/Xsxk/query/2")
	fmt.Println(req.Header)
	resp,_=cli.Do(req)
	body,_=io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func getCookie(key string)*http.Cookie {
	for i := 0; i < len(cookies); i++ {
		if cookies[i].Name == key {
			return cookies[i]
		}
	}
	return nil
}
func main(){
	username:="admin"
	passwd:="123456"
	log(username,passwd)
}
