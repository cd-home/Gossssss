package from

import "fmt"

// 被导入时就会执行
func init() {
	fmt.Println("from init function")
}
