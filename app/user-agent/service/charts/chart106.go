package charts

import "encoding/json"

func init() {
	globalCharts[106] = newChart106()
}

type resp106 struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Statistics [][]int  `json:"Statistics"`
		Classes    []string `json:"Classes"`
		Titles     []string `json:"Classes2"`
	} `json:"Option"`
}

type chart106 struct {
	name string
}

func newChart106() *chart106 {
	return &chart106{
		name: "技术生命周期分析",
	}
}

func (c *chart106) Serialize(params []byte) (string, error) {
	resp := &resp106{}
	if err := json.Unmarshal(params, resp); err != nil {
		return "", err
	}

	data := make([][2]int, 0, len(resp.Option.Statistics))
	for _, s := range resp.Option.Statistics {
		data = append(data, [2]int{s[0], s[1]})
	}
	names := resp.Option.Classes
	values := resp.Option.Classes
	xTitle := resp.Option.Titles[0]
	yTitle := resp.Option.Titles[1]

	points, err := makePoints(data, names, values)
	if err != nil {
		return "", err
	}

	bar := genLineWithMarkPoint(points, xTitle, yTitle)

	return bar, nil
}

func (c *chart106) Name() string {
	return c.name
}
