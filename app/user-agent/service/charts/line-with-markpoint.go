package charts

import (
	"encoding/json"
	"fmt"
)

const lineWithMarkPointProfile = `{
  "xAxis": {
    "name": "$X_NAME"
  },
  "yAxis": {
    "name": "$Y_NAME"
  },
  "series": [
    {
      "markPoint": $MART_POINT,
      "data": $DATA,
      "type": "line"
    }
  ]
}`

func genLineWithMarkPoint(points Points, xTitle string, yTitle string) string {
	p := newProfile(lineWithMarkPointProfile)

	return p.replace("$X_NAME", xTitle).
		replace("$Y_NAME", yTitle).
		replace("$MART_POINT", points.String()).
		replace("$DATA", points.PairsStr()).
		String()
}

type Point struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	XAxis int    `json:"xAxis"`
	YAxis int    `json:"yAxis"`
}

type Points []Point

func makePoints(data [][2]int, names []string, values []string) (Points, error) {
	if len(data) != len(names) || len(names) != len(values) {
		return nil, fmt.Errorf("invalid parameters")
	}

	res := make(Points, 0, len(names))

	for i := 0; i < len(data); i++ {
		res = append(res, Point{
			Name:  names[i],
			Value: values[i],
			XAxis: data[i][0],
			YAxis: data[i][1],
		})
	}

	return res, nil
}

func (ps Points) Pairs() [][2]int {
	pairs := make([][2]int, 0, len(ps))
	for _, p := range ps {
		pairs = append(pairs, [2]int{p.XAxis, p.YAxis})
	}
	return pairs
}

func (ps Points) PairsStr() string {
	pairs := ps.Pairs()
	bs, _ := json.Marshal(pairs)
	return string(bs)
}

func (ps Points) String() string {
	bs, _ := json.Marshal(ps)
	return string(bs)
}
