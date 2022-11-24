package charts

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const demo_301 = `{
    "ReturnValue": 0,
    "Option": {
        "Statistics": [
            [
                1218
            ],
            [
                724
            ],
            [
                689
            ],
            [
                654
            ],
            [
                582
            ],
            [
                424
            ],
            [
                350
            ],
            [
                326
            ],
            [
                313
            ],
            [
                307
            ]
        ],
        "Scripts": "",
        "Classes": [
            "深圳市大疆创新科技有限公司",
            "西北工业大学",
            "北京航空航天大学",
            "南京航空航天大学",
            "国家电网有限公司",
            "国家电网公司",
            "广州极飞科技有限公司",
            "深圳市道通智能航空技术有限公司",
            "广东电网有限责任公司",
            "易瓦特科技股份公司"
        ],
        "Query": "ti=无人机",
        "Dbs": "wgzl,syxx,fmzl",
        "patentSearchConfig": {
            "GUID": "NONE-0497059e46f24547bee12c540dc67d42",
            "Query": "ti=无人机",
            "RecordCount": 0,
            "PageCount": 0,
            "Database": "wgzl,syxx,fmzl",
            "KaxOptions": 0,
            "IdeoSingle": true,
            "MixedSort": false,
            "wxStatisticsField": false,
            "Page": 1,
            "PageSize": 1,
            "RecordIndex": 0,
            "UseOrginalQuery": false,
            "UserID": 0,
            "Merge_Record_Num": 0,
            "familyType": 0,
            "ContainsFullText": false,
            "NoCache": false,
            "DBOnly": 0,
            "mypatentcount": 0,
            "SecondQuery": "",
            "IsPatentList": false,
            "OrginalQuery": "ti=无人机",
            "imageType": 0,
            "IsDeduct": true,
            "isRelatedSearchPatents": false,
            "IsFuzzyMatched": 0
        }
    },
    "AID": "301",
    "TotalCount": 81614,
    "TotalCount2": 0,
    "page": 0,
    "total": -1,
    "records": 0,
    "IsGetCheckPic": false,
    "isWriteMsg": false
}`

func Test301(t *testing.T) {
	h, err := GetChart(301)
	assert.NoError(t, err)
	bs, err := h.Serialize([]byte(demo_301))
	fmt.Println(string(bs))
}
