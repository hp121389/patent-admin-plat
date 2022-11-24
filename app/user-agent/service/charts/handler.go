package charts

import (
	"errors"
	"fmt"
)

var ErrNoSuchChart = errors.New("no such chart")
var ErrInvalidParams = errors.New("invalid params")

var globalCharts = make(map[int]Handler)

type Handler interface {
	Serialize(params []byte) (string, error)
	Name() string
}

func GetChart(id int) (Handler, error) {
	h, ok := globalCharts[id]
	if !ok {
		return nil, fmt.Errorf("get chart by id %d failed, %w", id, ErrNoSuchChart)
	}
	return h, nil
}

func strListTemplate(l []string) string {
	res := ""
	for i, s := range l {
		if i == len(l)-1 {
			res += fmt.Sprintf("\"%s\"", s)
			continue
		}
		res += fmt.Sprintf("\"%s\", ", s)
	}
	return fmt.Sprintf("[%s]", res)
}

func intListTemplate(l []int) string {
	res := ""
	for i, d := range l {
		if i == len(l)-1 {
			res += fmt.Sprintf("%d", d)
			continue
		}
		res += fmt.Sprintf("%d, ", d)
	}
	return fmt.Sprintf("[%s]", res)
}
