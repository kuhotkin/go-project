package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (jiraConnector *jiraConnector) Get(ctx context.Context) (string, error) {
	//res, err := http.Get("https://jfrog.atlassian.net/rest/api/2/project/RTFACT")
	res, err := http.Get("https://jfrog.atlassian.net/rest/api/2/issue/RTFACT-29469")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(resBody))
	return "sgwreg", nil
}
