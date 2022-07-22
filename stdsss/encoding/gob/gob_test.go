package gob_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

// gob: go binary

type MessageSender struct {
	ServiceName string
	CallNumber  int64
	Request     interface{}
}

type MessageReceiver struct {
	ServiceName string
	CallNumber  *int64
	Request     interface{}
}

func TestGobEncodeAndDecode(t *testing.T) {
	// Encode the message
	fmt.Println("Encode:...........................................")

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)

	// 必须先注册 , 因为Message中有 interface{} 类型, 并且赋值的时候可能是 map struct 类型
	gob.Register(map[string]string{})

	request := map[string]string{
		"method": "gob",
		"mode":   "encode",
	}
	message := MessageSender{
		ServiceName: "Hello",
		CallNumber:  1001,
		Request:     request,
	}

	enc.Encode(&message)
	fmt.Println("Encode Mesaage: ", buf)

	// Mock Network Transport
	fmt.Println("Decode:...........................................")
	
	// "buf" can be networkdata
	dec := gob.NewDecoder(buf)
	var message2 MessageReceiver
	dec.Decode(&message2)

	fmt.Println("Decode Mesaage: ", message2)
}
