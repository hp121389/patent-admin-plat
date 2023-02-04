package charts

const lineSmoothProfile = `{
  "xAxis": {
    "type": "category",
    "data": $CATE,
	$ROTATE
  },
  "yAxis": {
    "type": "value"
  },
  "series": [
    {
      "data": $DATA,
      "type": "line",
      "smooth": true
    }
  ]
}`

func genLineSmoothProfile(cate []string, data []int, isRotate bool) string {
	p := newProfile(lineSmoothProfile)
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
