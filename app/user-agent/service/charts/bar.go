package charts

import "strings"

const barProfile = `{
  "tooltip": {
    "trigger": "axis",
    "axisPointer": {
      "type": "shadow"
    }
  },
  "grid": {
    "left": "3%",
    "right": "4%",
    "bottom": "3%",
    "containLabel": true
  },
  "xAxis": [
    {
      "type": "category",
      "data": $CATE,
      "axisTick": {
        "alignWithLabel": true
      },
      "axisLabel": {
        "interval": 0,
        "rotate": 45
      }
    }
  ],
  "yAxis": [
    {
      "type": "value"
    }
  ],
  "series": [
    {
      "name": "Direct",
      "type": "bar",
      "barWidth": "60%",
      "data": $DATA
    }
  ]
}`

func genBarProfile(cate []string, data []int) string {
	cateTemp := strListTemplate(cate)
	dataTemp := intListTemplate(data)
	return strings.Replace(strings.Replace(barProfile, "$CATE", cateTemp, 1), "$DATA", dataTemp, 1)
}
