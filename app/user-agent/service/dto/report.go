package dto

import "encoding/json"

type NoveltyReportReq struct {
	KeyWords  []string `json:"keyWords"`
	Title     string   `json:"title"`
	CL        string   `json:"CL"`
	Applicant string   `json:"applicant"`
	Org       string   `json:"org"`
}

type NoveltyReportResp struct {
	Number          string `json:"NUMBER"`
	DepartName      string `json:"DEPART_NAME"`
	ContactAddr     string `json:"CONTACT_ADDR"`
	ZipCode         string `json:"ZIP_CODE"`
	ManagerName     string `json:"MANAGER_NAME"`
	ManagerTel      string `json:"MANAGER_TEL"`
	ContactName     string `json:"CONTACT_NAME"`
	ContactTel      string `json:"CONTACT_TEL"`
	Email           string `json:"EMAIL"`
	Database        string `json:"DATABASE"`
	PatentName      string `json:"PATENT_NAME"`
	UserName        string `json:"USER_NAME"`
	Institution     string `json:"INSTITUTION"`
	FinishData      string `json:"FINISH_DATE"`
	TechPoint       string `json:"TECH_POINT"`
	QueryWord       string `json:"QUERY_WORD"`
	QueryExpression string `json:"QUERY_EXPRESSION"`
	RelativeNum     string `json:"RELATIVE_NUM"`
	VeryRelativeNum string `json:"VERY_RELATIVE_NUM"`
	SearchResult    string `json:"SEARCH_RESULT"`
	Conclusion      string `json:"CONCLUSION"`
}

func (r *NoveltyReportResp) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (r *NoveltyReportResp) Map() map[string]interface{} {
	bs, _ := json.Marshal(r)
	res := make(map[string]interface{})
	_ = json.Unmarshal(bs, &res)
	return res
}
