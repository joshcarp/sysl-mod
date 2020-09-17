package retriever_github

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/joshcarp/gop/gop"
)

type Retriever struct {
	token string
}

func New(token string) Retriever {
	return Retriever{token: token}
}

func (a Retriever) Retrieve(resource string) ([]byte, bool, error) {
	var resp *http.Response
	var apibase string
	var repo, path, version string
	var err error
	var res []byte

	repo, path, version, err = gop.ProcessRequest(resource)
	if err != nil {
		return nil, false, gop.CreateError(gop.BadRequestError, "Can't process request")
	}
	requestedurl, _ := url.Parse("https://" + resource)
	host := requestedurl.Host
	repo = strings.ReplaceAll(repo, host+"/", "")

	switch host {
	case "github.com":
		apibase = "api.github.com"
	default:
		apibase = fmt.Sprintf("%s/api/v3", host)
	}

	req, err := url.Parse(
		fmt.Sprintf(
			"https://%s/repos/%s/contents/%s?ref=%s",
			apibase, repo, path, version))
	heder := http.Header{}
	heder.Add("accept", "application/vnd.github.v3.raw+json")

	if a.token != "" {
		heder.Add("authorization", "token "+a.token)
	}

	r := &http.Request{
		Method: "GET",
		URL:    req,
		Header: heder,
	}

	resp, err = http.DefaultClient.Do(r)
	if err != nil {
		return res, false, err
	}
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}
	return res, false, nil
}
