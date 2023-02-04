package charts

import "encoding/json"

func init() {
	globalCharts[1315] = newChart1315()
}

type resp1315 struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Statistics [][]int  `json:"Statistics"`
		Classes    []string `json:"Classes"`
	} `json:"Option"`
}

type chart1315 struct {
	name string
}

func newChart1315() *chart1315 {
	return &chart1315{
		name: "年度专利许可趋势分析",
	}
}

func (c *chart1315) Serialize(params []byte) (string, error) {
	resp := &resp1315{}
	if err := json.Unmarshal(params, resp); err != nil {
		return "", err
	}

	data := make([]int, 0, len(resp.Option.Statistics))
	for _, s := range resp.Option.Statistics {
		data = append(data, s[0])
	}
	chart := genLineSmoothProfile(resp.Option.Classes, data, true)

	return chart, nil
}

func (c *chart1315) Name() string {
	return c.name
}
