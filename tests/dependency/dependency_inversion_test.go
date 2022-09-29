package dependency

import (
	"fmt"
	"testing"
)

type ICar interface {
	Run()
}

type Car struct {
	Type string `json:"type"`
}

func (c *Car) Run() {
	fmt.Printf("%s: running\n", c.Type)
}

type IDriver interface {
	DriverCar(car Car)
}

type Driver struct {
	Name string `json:"name"`
}

func (d *Driver) DriverCar(car ICar) {
	fmt.Printf("%s: driver\n", d.Name)
	car.Run()
}

func TestDependencyInversion(t *testing.T) {
	jack := Driver{Name: "jack"}
	bmw := &Car{Type: "bmm"}
	jack.DriverCar(bmw)

	ft := &Car{Type: "ft"}
	jack.DriverCar(ft)

	mike := Driver{Name: "mike"}
	benz := &Car{Type: "benz"}
	mike.DriverCar(benz)
}
