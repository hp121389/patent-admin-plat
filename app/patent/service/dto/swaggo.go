package dto

type SwagSearchListResp struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      struct {
		Count     int `json:"count"`
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
		List      []struct {
			NO        int    `json:"NO"`
			AN        string `json:"AN"`
			ExamAN    string `json:"ExamAN"`
			AD        string `json:"AD"`
			PNM       string `json:"PNM"`
			PNM2      string `json:"PNM2"`
			PD        string `json:"PD"`
			TI        string `json:"TI"`
			PA        string `json:"PA"`
			DB        string `json:"DB"`
			ISNEWDATA bool   `json:"ISNEWDATA"`
			PINN      string `json:"PINN"`
			RNO       int    `json:"RNO,omitempty"`
		} `json:"list"`
	} `json:"data"`
}
