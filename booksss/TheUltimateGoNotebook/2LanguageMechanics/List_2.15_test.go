package example_test

import "testing"

const (
	A1 = iota // 0 : Start at 0
	B1 = iota // 1 : Increment by 1
	C1 = iota // 2 : Increment by 1
)

const (
	A2 = iota // 0 : Start at 0
	B2        // 1 : Increment by 1
	C2        // 2 : Increment by 1
)

const (
	A3 = iota + 1 // 1 : Start at 1
	B3            // 2 : Increment by 1
	C3            // 3 : Increment by 1
)

const (
	Ldate         = 1 << iota // 1 : Shift 1 to the left 0. 0000 0001
	Ltime                     // 2 : Shift 1 to the left 1. 0000 0010
	Lmicroseconds             // 4 : Shift 1 to the left 2. 0000 0100
	Llongfile                 // 8 : Shift 1 to the left 3. 0000 1000
	Lshortfile                // 16 : Shift 1 to the left 4. 0001 0000
	LUTC                      // 32 : Shift 1 to the left 5. 0010 0000
)

func TestIota(t *testing.T) {
	t.Log(A1)
	t.Log(B1)
	t.Log(C1)

	t.Log(A2)
	t.Log(B2)
	t.Log(C2)

	t.Log(A3)
	t.Log(B3)
	t.Log(C3)
}
