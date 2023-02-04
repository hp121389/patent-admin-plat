package charts

import "encoding/json"

func init() {
	globalCharts[1316] = newChart1316()
}

type resp1316 struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Statistics [][]int  `json:"Statistics"`
		Classes    []string `json:"Classes"`
	} `json:"Option"`
}

type chart1316 struct {
	name string
}

func newChart1316() *chart1316 {
	return &chart1316{
		name: "许可人排名分析",
	}
}

func (c *chart1316) Serialize(params []byte) (string, error) {
	resp := &resp1316{}
	if err := json.Unmarshal(params, resp); err != nil {
		return "", err
	}

	data := make([]int, 0, len(resp.Option.Statistics))
	for _, s := range resp.Option.Statistics {
		data = append(data, s[0])
	}
	chart := genVerticalBarProfile(resp.Option.Classes, data, true)

	return chart, nil
}

func (c *chart1316) Name() string {
	return c.name
}
