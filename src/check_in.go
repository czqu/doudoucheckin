package doudou

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	_ "strings"
)

func CheckIn(cookies []*http.Cookie) error {

	u := CHECK_IN_URL
	req, err := http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return err
	}

	var client = &http.Client{}
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("user-agent", UA)
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	for _, c := range cookies {
		req.AddCookie(c)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	res, _ := UnescapeUnicode(body)
	fmt.Println(res)
	return nil

}
func Login(email string, pwd string) ([]*http.Cookie, error) {
	u := LOGIN
	req, err := http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	var client = &http.Client{}
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("user-agent", UA)
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	query := req.URL.Query()
	query.Add("email", email)
	query.Add("passwd", pwd)
	req.URL.RawQuery = query.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	res, _ := UnescapeUnicode(body)
	if strings.Contains(res, "错误") {
		return nil, errors.New(res)
	}
	fmt.Println(res)

	return resp.Cookies(), nil

}
func UnescapeUnicode(raw []byte) (string, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}
