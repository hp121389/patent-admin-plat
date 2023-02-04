package charts

import "encoding/json"

func init() {
	globalCharts[501] = newChart501()
}

type resp501 struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Statistics [][]int  `json:"Statistics"`
		Classes    []string `json:"Classes"`
	} `json:"Option"`
}

type chart501 struct {
	name string
}

func newChart501() *chart501 {
	return &chart501{
		name: "发明人排行榜",
	}
}

func (c *chart501) Serialize(params []byte) (string, error) {
	resp := &resp501{}
	if err := json.Unmarshal(params, resp); err != nil {
		return "", err
	}

	data := make([]int, 0, len(resp.Option.Statistics))
	for _, s := range resp.Option.Statistics {
		data = append(data, s[0])
	}
	chart := genBarProfile(resp.Option.Classes, data, true)

	return chart, nil
}

func (c *chart501) Name() string {
	return c.name
}
