package snippetsss

import (
	"fmt"
	"testing"
)

// abstract interface
type ABCBanker interface {
	DoWork()
}

// Single responsibility 单一职责
type SaveBanker struct{}

func (s *SaveBanker) DoWork() {
	fmt.Println("Save")
}

// Single responsibility 单一职责
type TransferBanker struct{}

func (s *TransferBanker) DoWork() {
	fmt.Println("Transfer")
}

// Single responsibility 单一职责
type PayBanker struct{}

func (s *PayBanker) DoWork() {
	fmt.Println("Pay")
}

func CallDoWork(banker ABCBanker) {
	banker.DoWork()
}

// TestPolymorphismType 多态
func TestPolymorphismType(t *testing.T) {
	CallDoWork(&SaveBanker{})
	CallDoWork(&TransferBanker{})
	CallDoWork(&PayBanker{})
}
