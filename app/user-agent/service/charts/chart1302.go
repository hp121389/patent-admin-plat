package charts

import "encoding/json"

func init() {
	globalCharts[1302] = newChart1302()
}

type resp1302 struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Statistics [][]int  `json:"Statistics"`
		Classes    []string `json:"Classes"`
	} `json:"Option"`
}

type chart1302 struct {
	name string
}

func newChart1302() *chart1302 {
	return &chart1302{
		name: "年度专利转让趋势分析",
	}
}

func (c *chart1302) Serialize(params []byte) (string, error) {
	resp := &resp1302{}
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

func (c *chart1302) Name() string {
	return c.name
}
