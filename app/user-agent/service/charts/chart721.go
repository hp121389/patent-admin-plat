package charts

import "encoding/json"

func init() {
	globalCharts[721] = newChart721()
}

type resp721 struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Statistics [][]int  `json:"Statistics"`
		Classes    []string `json:"Classes"`
	} `json:"Option"`
}

type chart721 struct {
	name string
}

func newChart721() *chart721 {
	return &chart721{
		name: "发明人排行榜",
	}
}

func (c *chart721) Serialize(params []byte) (string, error) {
	resp := &resp721{}
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

func (c *chart721) Name() string {
	return c.name
}
