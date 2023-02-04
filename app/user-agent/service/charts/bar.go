package charts

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
      $ROTATE
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

const ROTATE = `
      "axisTick": {
        "alignWithLabel": true
      },
      "axisLabel": {
        "interval": 0,
        "rotate": 45
      }`

func genBarProfile(cate []string, data []int, isRotate bool) string {
	p := newProfile(barProfile)
	if isRotate {
		p = p.replace("$ROTATE", ROTATE)
	} else {
		p = p.replace("$ROTATE", "").
			replace("$CATE,", "$CATE")
	}
	cateTemp := strListTemplate(cate)
	dataTemp := intListTemplate(data)

	return p.replace("$CATE", cateTemp).
		replace("$DATA", dataTemp).
		String()
}
