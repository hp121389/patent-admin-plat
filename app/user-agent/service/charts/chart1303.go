package charts

import "encoding/json"

func init() {
	globalCharts[1303] = newChart1303()
}

type resp1303 struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Statistics [][]int  `json:"Statistics"`
		Classes    []string `json:"Classes"`
	} `json:"Option"`
}

type chart1303 struct {
	name string
}

func newChart1303() *chart1303 {
	return &chart1303{
		name: "省市申请量分析",
	}
}

func (c *chart1303) Serialize(params []byte) (string, error) {
	resp := &resp1303{}
	if err := json.Unmarshal(params, resp); err != nil {
		return "", err
	}

	data := make([]int, 0, len(resp.Option.Statistics))
	for _, s := range resp.Option.Statistics {
		data = append(data, s[0])
	}
	chart := genVerticalBarProfile(resp.Option.Classes, data, false)

	return chart, nil
}

func (c *chart1303) Name() string {
	return c.name
}
