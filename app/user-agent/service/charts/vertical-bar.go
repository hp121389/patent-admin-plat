package charts

const verticalBarProfile = `{
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
      "type": "value"
    }
  ],
  "yAxis": [
    {
      "type": "category",
      "data": $CATE,
      $ROTATE
    }
  ],
  "series": [
    {
      "type": "bar",
      "stack": "total",
      "label": {
        "show": true
      },
      "emphasis": {
        "focus": "series"
      },
      "data": $DATA
    }
  ]
}`

func genVerticalBarProfile(cate []string, data []int, isRotate bool) string {
	p := newProfile(verticalBarProfile)
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
