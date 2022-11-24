package charts

import "encoding/json"

func init() {
	globalCharts[301] = newChart301()
}

type resp301 struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Statistics [][]int  `json:"Statistics"`
		Classes    []string `json:"Classes"`
	} `json:"Option"`
}

type chart301 struct {
	name string
}

func newChart301() *chart301 {
	return &chart301{
		name: "申请人排行榜",
	}
}

func (c *chart301) Serialize(params []byte) (string, error) {
	resp := &resp301{}
	if err := json.Unmarshal(params, resp); err != nil {
		return "", err
	}

	data := make([]int, 0, len(resp.Option.Statistics))
	for _, s := range resp.Option.Statistics {
		data = append(data, s[0])
	}
	bar := genBarProfile(resp.Option.Classes, data)

	return bar, nil
}

func (c *chart301) Name() string {
	return c.name
}
