package gob_test

import (
	"bytes"
	"encoding/gob"
	"testing"
)

// gob: go binary

type MessageSender struct {
	Name    string
	Seq     int64
	Request interface{}
}

type MessageReceiver struct {
	Name    string
	Seq     *int64
	Request interface{}
}

func TestGobEncodeAndDecodeIncludeMap(t *testing.T) {

	var network bytes.Buffer

	encoder := gob.NewEncoder(&network)

	// 如果值中含有interface类型, 必须先注册传输的具体类型 (如果该类型是 Map 或者 Struct)
	gob.Register(map[string]string{})

	request := map[string]string{
		"method": "gob",
		"mode":   "encode",
	}
	message := MessageSender{
		Name:    "Hello",
		Seq:     1001,
		Request: request,
	}

	if err := encoder.Encode(&message); err != nil {
		t.Fatal(err)
	}

	// Receiver
	dec := gob.NewDecoder(&network)
	var messageReceiver MessageReceiver
	dec.Decode(&messageReceiver)

	t.Log(messageReceiver)
}

func TestGobEncodeAndDecodeIncludeStruct(t *testing.T) {
	var network bytes.Buffer

	encoder := gob.NewEncoder(&network)

	// 如果值中含有interface类型, 必须先注册传输的具体类型 (如果该类型是 Map 或者 Struct)
	type Request struct {
		Method string
		Mode   string
	}

	gob.Register(Request{})

	request := &Request{
		Method: "gob",
		Mode:   "encode",
	}

	message := MessageSender{
		Name:    "Hello",
		Seq:     1001,
		Request: request,
	}

	if err := encoder.Encode(&message); err != nil {
		t.Fatal(err)
	}

	// Receiver
	dec := gob.NewDecoder(&network)
	var messageReceiver MessageReceiver
	dec.Decode(&messageReceiver)

	t.Log(messageReceiver)
}
