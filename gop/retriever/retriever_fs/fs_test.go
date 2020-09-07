package retriever_fs

import (
	"testing"

	"github.com/joshcarp/pb-mod/app"
	"github.com/joshcarp/pb-mod/gen/pkg/servers/pbmod"
	"github.com/stretchr/testify/require"
)

func TestFsRetrieve(t *testing.T) {
	r := New(app.AppConfig{FsType: "memory"})
	file, err := r.fs.Create("github.com/anz-bank/sysl/tests/bananatree.sysl@e78f4afc524ad8d1a1a4740779731d706b7b079b")
	require.NoError(t, err)
	file.Write([]byte(bananatree))
	file, err = r.fs.Create("github.com/anz-bank/sysl/tests/bananatree.sysl.pb.json@e78f4afc524ad8d1a1a4740779731d706b7b079b")
	require.NoError(t, err)
	file.Write([]byte(bananatreepbjson))
	req := app.NewObject("github.com/anz-bank/sysl/tests/bananatree.sysl", "e78f4afc524ad8d1a1a4740779731d706b7b079b")
	require.NoError(t, r.Retrieve(req))
	banana := &pbmod.Object{Imported: true, Repo: "github.com/anz-bank/sysl", Version: "e78f4afc524ad8d1a1a4740779731d706b7b079b", Resource: "tests/bananatree.sysl", Content: "Bananatree [package=\"bananatree\"]:\n  !type Banana:\n    id \u003c: int\n    title \u003c: string\n\n  /banana:\n    /{id\u003c:int}:\n      GET:\n        return Banana\n\n  /morebanana:\n    /{id\u003c:int}:\n      GET:\n        return Banana\n"}
	req.Processed = nil
	require.Equal(t, banana, req)
}

const bananatree = `Bananatree [package="bananatree"]:
  !type Banana:
    id <: int
    title <: string

  /banana:
    /{id<:int}:
      GET:
        return Banana

  /morebanana:
    /{id<:int}:
      GET:
        return Banana
`

const bananatreepbjson = `{"apps":{"Bananatree":{"name":{"part":["Bananatree"]}, "attrs":{"package":{"s":"bananatree"}}, "endpoints":{"GET /banana/{id}":{"name":"GET /banana/{id}", "attrs":{"patterns":{"a":{"elt":[{"s":"rest"}]}}}, "stmt":[{"ret":{"payload":"Banana"}}], "restParams":{"method":"GET", "path":"/banana/{id}", "urlParam":[{"name":"id", "type":{"primitive":"INT", "sourceContext":{"file":"temp.sysl", "start":{"line":7, "col":5}, "end":{"line":7, "col":13}}}}]}, "sourceContext":{"file":"temp.sysl", "start":{"line":8, "col":6}, "end":{"line":11, "col":2}}}, "GET /morebanana/{id}":{"name":"GET /morebanana/{id}", "attrs":{"patterns":{"a":{"elt":[{"s":"rest"}]}}}, "stmt":[{"ret":{"payload":"Banana"}}], "restParams":{"method":"GET", "path":"/morebanana/{id}", "urlParam":[{"name":"id", "type":{"primitive":"INT", "sourceContext":{"file":"temp.sysl", "start":{"line":12, "col":5}, "end":{"line":12, "col":13}}}}]}, "sourceContext":{"file":"temp.sysl", "start":{"line":13, "col":6}, "end":{"line":15}}}}, "types":{"Banana":{"tuple":{"attrDefs":{"id":{"primitive":"INT", "sourceContext":{"file":"temp.sysl", "start":{"line":3, "col":10}, "end":{"line":3, "col":10}}}, "title":{"primitive":"STRING", "sourceContext":{"file":"temp.sysl", "start":{"line":4, "col":13}, "end":{"line":4, "col":13}}}}}, "sourceContext":{"file":"temp.sysl", "start":{"line":2, "col":2}, "end":{"line":6, "col":2}}}}, "sourceContext":{"file":"temp.sysl", "start":{"line":1, "col":1}, "end":{"line":1, "col":32}}}}}`