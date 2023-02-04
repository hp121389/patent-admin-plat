package charts

import "encoding/json"

func init() {
	globalCharts[201] = newChart201()
}

type resp201 struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Statistics [][]int  `json:"Statistics"`
		Classes    []string `json:"Classes"`
	} `json:"Option"`
}

type chart201 struct {
	name string
}

func newChart201() *chart201 {
	return &chart201{
		name: "年度申请量分析",
	}
}

func (c *chart201) Serialize(params []byte) (string, error) {
	resp := &resp201{}
	if err := json.Unmarshal(params, resp); err != nil {
		return "", err
	}

	data := make([]int, 0, len(resp.Option.Statistics))
	for _, s := range resp.Option.Statistics {
		data = append(data, s[0])
	}
	bar := genBarProfile(resp.Option.Classes, data, false)

	return bar, nil
}

func (c *chart201) Name() string {
	return c.name
}
