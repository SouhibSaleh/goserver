package tasks

import "fmt"

func solve() <-chan struct{} {
	ch := make(chan struct{})
	defer close(ch)
	ch <- struct{}{}
	return ch
}

func TestArea() {

	l := solve()
	for s := range l {
		fmt.Println(s, "t")
	}

}
