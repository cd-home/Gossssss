package main

import (
	"errors"
	"fmt"
)

// GunFace 公共接口
type GunFace interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

type Gun struct {
	Name  string
	Power int
}

func (g *Gun) SetName(name string) {
	g.Name = name
}

func (g *Gun) SetPower(power int) {
	g.Power = power
}

func (g *Gun) GetName() string {
	return g.Name
}

func (g *Gun) GetPower() int {
	return g.Power
}

// Ak47 产品
type Ak47 struct {
	Gun
}

// NewAk47 具体工厂
func NewAk47() GunFace {
	return &Ak47{Gun{
		Name:  "AK-47",
		Power: 10,
	}}
}

// Musket 产品
type Musket struct {
	Gun
}

// NewMusket 具体工厂
func NewMusket() GunFace {
	return &Musket{Gun{
		Name:  "Musket",
		Power: 6,
	}}
}

// GetGun 工厂方法
func GetGun(gunType string) (GunFace, error) {
	switch gunType {
	case "AK47":
		return NewAk47(), nil
	case "Musket":
		return NewMusket(), nil
	default:
		return nil, errors.New("wrong Gun Type")
	}
}

func main() {
	gun, _ := GetGun("AK47")
	var gunFace GunFace = gun
	fmt.Println(gunFace.GetName())
	fmt.Println(gunFace.GetPower())
}
