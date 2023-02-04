package dto

import (
	"strconv"
	"strings"
	"time"
)

type Dashboard struct {
	// 专利认领数量
	PatentClaimCount int `json:"patentClaimCount"`
	// 专利关注数量
	PatentFocusCount int `json:"patentFocusCount"`
	// 专利法律状态
	PatentStatus map[string]int `json:"patentStatus"`
	// 发表专利时间
	PublicationDates []PublicationDate `json:"publicationDates"`
	// 工艺包数量
	PackageCount int `json:"packageCount"`
	// 专利推荐
	PatentRecommendation []*PatentDetail `json:"patentRecommendation"`
	// 报告数量
	ReportCount int `json:"reportCount"`
	// 专利预估价
	PatentTotalPrice int `json:"patentTotalPrice"`
	// 合作者
	Collaborators []*Researcher `json:"collaborators"`
	// 竞争者
	Competitors []*Researcher `json:"competitors"`
}

type PublicationDate struct {
	Year  int `json:"year"`
	Count int `json:"count"`
}

func NewPublicationDates() *PublicationDates {
	res := make([]PublicationDate, 10)
	currentYear := time.Now().Year()
	for i := 0; i < 10; i++ {
		res[i] = PublicationDate{
			Year:  currentYear,
			Count: 0,
		}
		currentYear--
	}
	return &PublicationDates{pds: res}
}

type PublicationDates struct {
	pds []PublicationDate
}

func (pds *PublicationDates) AddYear(date string) {
	yearStr := strings.Split(date, ".")[0]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return
	}
	currentYear := time.Now().Year()
	index := currentYear - year
	if index > 9 {
		return
	}
	pds.pds[index].Count++
}

func (pds *PublicationDates) List() []PublicationDate {
	return pds.pds
}

type ReportInfo struct {
	Novelty      int `json:"novelty"`
	Infringement int `json:"infringement"`
	Valuation    int `json:"valuation"`
}

type Researcher struct {
	Name  string `json:"name"`
	Times int    `json:"times"`
}
