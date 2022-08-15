package gob_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type LinkFace interface {
	GenerateLink() string
}

type Web struct {
	Proto  string
	Domain string
	Port   string
}

func (w Web) GenerateLink() string {
	return fmt.Sprintf("%s://%s:%s", w.Proto, w.Domain, w.Port)
}

// How to encode an interface value
func TestEmptyInterfaceGob(t *testing.T) {
	var network bytes.Buffer

	// We must register the concrete type for the encoder and decoder (which would
	// normally be on a separate machine from the encoder).
	// On each end, this tells the engine which concrete type is being sent that implements the interface.
	// 必须为编码与解码注册具体的类型, 另一方面也是告知引擎正在发送那个具体类型来实现了接口
	gob.Register(&Web{})

	// Encode
	encoder := gob.NewEncoder(&network)

	var webSendFace LinkFace = &Web{Proto: "http", Domain: "godyao.com.cn", Port: "8090"}

	// Must Pointer to interface
	if err := encoder.Encode(&webSendFace); err != nil {
		t.Fatal(err)
	}

	// Decoder also Register gob.Register(&Web{})
	// Decode
	decoder := gob.NewDecoder(&network)
	var webReciverFace LinkFace
	if err := decoder.Decode(&webReciverFace); err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", webReciverFace.GenerateLink())
}
