package charts

import (
	"encoding/json"
)

const pieProfile = `{
  "tooltip": {
    "trigger": "item"
  },
  "legend": {
    "top": "5%",
    "left": "center"
  },
  "series": [
    {
      "name": "Access From",
      "type": "pie",
      "radius": ["40%", "70%"],
      "avoidLabelOverlap": false,
      "itemStyle": {
        "borderRadius": 10,
        "borderColor": "#fff",
        "borderWidth": 2
      },
      "label": {
        "show": false,
        "position": "center"
      },
      "emphasis": {
        "label": {
          "show": true,
          "fontSize": 40,
          "fontWeight": "bold"
        }
      },
      "labelLine": {
        "show": false
      },
      "data": $DATA
    }
  ]
}`

func pieDataTemplate(classes []string, values []int) string {
	pieList := make(PieDataList, 0)
	for i := 0; i < len(classes); i++ {
		pieList = append(pieList, PieData{
			Name:  classes[i],
			Value: values[i],
		})
	}
	return pieList.String()
}

type PieData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type PieDataList []PieData

func (d *PieDataList) String() string {
	bs, _ := json.Marshal(d)
	return string(bs)
}

func genPieProfile(classes []string, data []int) string {
	template := pieDataTemplate(classes, data)
	p := newProfile(pieProfile)
	return p.replace("$DATA", template).String()
}
