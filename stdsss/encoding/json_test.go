package encoding

import (
	"encoding/json"
	"testing"
)

type Data struct {
	Code int    `json:"code"`
	Key  string `json:"key"`
}

func TestSJson(t *testing.T) {
	bs := []byte(`{"code": 200}`)
	var data Data
	json.Unmarshal(bs, &data)
	t.Log(data)
}
