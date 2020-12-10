package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

const (
	Touser     = "haochuzhan"                                  //企业号中的用户帐号，在zabbix用户Media中配置，如果配置不正常，将按部门发送。
	Toparty    = "1"                                           //企业号中的部门id。
	Agentid    = 1000011                                       //企业号中的应用id。
	Corpid     = "wwee4c443448294ecb"                          //企业号的标识
	Corpsecret = "0uo2AX0A-firYFi8kQHCujSeJ4PLlc_13_qL7_-lA2k" ///企业号中的应用的Secret
)

/*
type JSON struct {
	AccessToken string `json:"access_token"`
}


type Message struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe int `json:"safe"`
}

func httpGetJson(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func httpPostJson(url string, data map[string]interface{}) (map[string]interface{}, error) {
	xxx, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(xxx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data2 map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data2, nil
}
*/

func HttpGetJson(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func HttpPostJson(url string, data map[string]interface{}) (map[string]interface{}, error) {
	xxx, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(xxx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data2 map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data2)
	if err != nil {
		return nil, err
	}

	return data2, nil
}

//
func SendMsg() {
	//var todo models.Todo
	//c.BindJSON(&todo)
	//corporate_name := c.PostForm("corporate_name")
	//name := c.PostForm("name")
	//phone := c.PostForm("phone")
	//email := c.PostForm("email")
	//advice := c.PostForm("advice")

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", Corpid, Corpsecret)
	data, err := HttpGetJson(url)
	/*
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})*/
	if err != nil {
		fmt.Println(err)

		return
	}
	errcode := data["errcode"].(float64) //注意这里必须强制类型转换
	if errcode != 0 {
		fmt.Println("errcode:", errcode)
		return
	}
	access_token := data["access_token"]

	req := map[string]interface{}{
		"touser":  Touser,
		"msgtype": "text",
		"agentid": Agentid,
		"text": map[string]interface{}{
			"content": "官网有企业提交合作申请，注意查看！\n可查看<a href=\"http://work.weixin.qq.com\">赛程后台</a>。",
		},
		"safe": 0,
	}
	url = fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", access_token)
	data, err = HttpPostJson(url, req)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("发送ok")

}

// 正则过滤sql注入的方法
// 参数 : 要匹配的语句
func FilteredSQLInject(to_match_str string) bool {
	//过滤 ‘
	//ORACLE 注解 --  /**/
	//关键字过滤 update ,delete
	// 正则的字符串, 不能用 " " 因为" "里面的内容会转义
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		return false
	}
	return re.MatchString(to_match_str)
}
