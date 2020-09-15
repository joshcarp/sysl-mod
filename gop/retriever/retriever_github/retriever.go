package retriever_github

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/joshcarp/gop/gop"
)

type Retriever struct {
}

func New() Retriever {
	return Retriever{}
}

func (a Retriever) Retrieve(resource string) (res []byte, cached bool, err error) {
	var resp *http.Response
	repo, resource, version, err := gop.ProcessRequest(resource)
	if err != nil {
		return nil, false, gop.CreateError(gop.BadRequestError, "Can't process request")
	}
	req := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s",
		strings.ReplaceAll(repo, "github.com/", ""), version, resource)
	resp, err = http.Get(req)
	if err != nil {
		return res, false, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, false, err
	}
	return bytes, false, nil
}
