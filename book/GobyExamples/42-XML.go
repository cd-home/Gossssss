package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	coffee := &Plant{
		Id:      100,
		Name:    "Coffee",
	}
	coffee.Origin = []string{"Eth", "Bra"}

	out, _ := xml.MarshalIndent(coffee, "", "  ")
	//fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))
}

type Plant struct {
	// XMLName 字段名称规定了该结构的 XML 元素名称
	XMLName xml.Name `xml:"plant"`
	// id,attrr 表示 Id 字段是一个 XML 属性，而不是嵌套元素
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%s, origin=%v", p.Id, p.Name, p.Origin)
}
