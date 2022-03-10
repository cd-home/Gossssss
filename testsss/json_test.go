package testsss

import (
	"encoding/json"
	"testing"
)

type TestJsonSuite struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Tips string      `json:"tips,omitempty"`
}

func TestJson(t *testing.T) {
	data := TestJsonSuite{
		Code: "1",
		Data: "data",
		Msg: "msg",
		Tips: "tips",
	}

	// bs, e := json.Marshal(data)
	bs, e := json.MarshalIndent(data, "", " ")
	if e != nil {
		t.Error(e)
	}
	t.Log(string(bs))

	res := TestJsonSuite{}
	json.Unmarshal(bs, &res)
	t.Log(res)
}
