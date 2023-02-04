package charts

import "encoding/json"

func init() {
	globalCharts[101] = newChart101()
}

type resp101 struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Statistics [][]int  `json:"Statistics"`
		Classes    []string `json:"Classes"`
	} `json:"Option"`
}

type chart101 struct {
	name string
}

func newChart101() *chart101 {
	return &chart101{
		name: "专利类型分布",
	}
}

func (c *chart101) Serialize(params []byte) (string, error) {
	resp := &resp101{}
	if err := json.Unmarshal(params, resp); err != nil {
		return "", err
	}

	data := make([]int, 0, len(resp.Option.Statistics))
	for _, s := range resp.Option.Statistics {
		data = append(data, s[0])
	}
	pie := genPieProfile(resp.Option.Classes, data)

	return pie, nil
}

func (c *chart101) Name() string {
	return c.name
}
