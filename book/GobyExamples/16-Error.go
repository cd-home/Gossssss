package main

import (
	"errors"
	"fmt"
)

func main() {
	v, err := division(10, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v)
	}

	v1, err := division(12, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v1)
	}

	v2, err := division2(100, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v2)
	}
}

// error 类型一般作为最后一个返回参数
func division(molecule, denominator float64) (float64, error) {
	if denominator == 0 {
		// new返回的是 error 指针类型
		return -999999, errors.New("0 can not be denominator")
	}
	return molecule / denominator, nil
}

// ZeroError 自定义错误类型
type ZeroError struct {
	arg float64
	msg string
}

// 必须实现Error方法
func (z *ZeroError) Error() string {
	return fmt.Sprintf("%f %s", z.arg, z.msg)
}

func division2(molecule, denominator float64) (float64, error) {
	if denominator == 0 {
		return -999999, &ZeroError{
			arg: denominator,
			msg: "can not be denominator",
		}
	}
	return molecule / denominator, nil
}
