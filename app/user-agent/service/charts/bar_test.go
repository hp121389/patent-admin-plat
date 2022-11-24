package charts

import (
	"fmt"
	"testing"
)

func TestBar(t *testing.T) {
	cate := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	data := []int{10, 52, 200, 334, 390, 330, 220}
	s := genBarProfile(cate, data)
	fmt.Println(s)
}
