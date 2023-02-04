package service

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/google/uuid"
	"go-admin/app/admin-agent/model"
	"go-admin/app/user-agent/my_config"
	"go-admin/app/user-agent/service/dto"
	"go-admin/app/user-agent/service/report"
	"go-admin/app/user-agent/utils"
	cDto "go-admin/common/dto"
	"gorm.io/gorm"
	"strconv"
	"strings"
	//"gorm.io/gorm"
)

type Report struct {
	service.Service
}

const noveltyReportType = "novelty"

// GetCountByUserID 获取Report对象
func (e *Report) GetCountByUserID(uid int, count *int64) error {
	//引用传递、函数名、形参、返回值
	var err error
	var data model.ReportRelation
	err = e.Orm.Model(&data).Where("user_id = ?", uid).Count(count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetNovelty 获取查新报告
func (e *Report) GetNovelty(c *dto.NoveltyReportReq) (*dto.NoveltyReportResp, error) {
	query := report.GenQuery(c.KeyWords)

	checkMap := make(map[string]*dto.PatentDetail)
	for i := 0; i < len(query) && len(checkMap) < 30; i++ {
		searchReq := &dto.SimpleSearchReq{
			Pagination: cDto.Pagination{
				PageIndex: 1,
				PageSize:  100,
			},
			Query: query[i],
			DB:    "wgzl,syxx,fmzl",
		}
		checkListTemp, err := GetCurrentInnojoy().Search(searchReq, nil)
		if err != nil {
			return nil, err
		}
		for _, res := range checkListTemp {
			if _, ok := checkMap[res.Pnm]; !ok {
				checkMap[res.Pnm] = res
			}
		}
	}

	var checkList []*dto.PatentDetail
	for _, pd := range checkMap {
		checkList = append(checkList, pd)
	}

	var totalinfo []string
	for j := 0; j < len(checkList); j++ {
		totalinfo = append(totalinfo, checkList[j].Ti+"，"+checkList[j].Abst)
	}

	var sentence = c.Title + "。" + c.CL
	segments := report.Seg.Segment([]byte(sentence))
	result := report.GenKey(segments)
	GenVerb := report.GetResult(segments, 2)
	ts2 := report.NewTextSimilarity(strings.Split(sentence, "。"))

	var sims []report.Similarity
	for j := 0; j < len(checkList); j++ {
		var temp = report.Similarity{}
		for i := 0; i < len(result); i++ {
			if strings.Contains(checkList[j].Ti, result[i]) || strings.Contains(checkList[j].Abst, result[i]) {
				temp.Count++
				temp.Words = append(temp.Words, result[i])
			}
		}
		segments1 := report.Seg.Segment([]byte(checkList[j].Ti + checkList[j].Abst))
		resWords1 := report.GetResult(segments1, 0)
		result1 := report.RemoveStop(report.Unique(resWords1))
		temp.Score, _ = ts2.Similarity(result1, result)
		sims = append(sims, temp)
	}
	keywords := ts2.Keywords(-1, 5, 0)
	keywords = report.Unique(keywords)
	var searchList string
	var searchWord = make([][]string, len(keywords))
	for i := 0; i < len(keywords); i++ {
		contain := false
		for j := 0; j < len(GenVerb); j++ {
			if keywords[i] == GenVerb[j] {
				contain = true
				break
			}
		}
		if contain == true {
			continue
		}
		similar := report.GetSimilar(keywords[i])
		searchList += keywords[i] + " " + strings.Join(similar, " ") + "\n"
		searchWord[i] = make([]string, 0)
		searchWord[i] = append(searchWord[i], keywords[i])
		searchWord[i] = append(searchWord[i], similar...)
	}
	queryExpression := report.GenQueryExpression(searchWord)
	n := len(sims)
	var conclusion []string
	var retconc string
	var closeCount = 1 //密切相关的专利数量
	for i := 0; i < n-1 && closeCount < 15; i++ {
		maxNumIndex := i // 无序区第一个
		for j := i + 1; j < n; j++ {
			if sims[j].Score > sims[maxNumIndex].Score {
				maxNumIndex = j
			}
		}
		sims[i], sims[maxNumIndex] = sims[maxNumIndex], sims[i]
		checkList[i], checkList[maxNumIndex] = checkList[maxNumIndex], checkList[i]
		if sims[i].Score > 0.4 {
			if c.Title == checkList[i].Ti {
				continue
			}
			header := report.GenConclusionHeader(closeCount, checkList[i].Pinn, checkList[i].Pa, checkList[i].Ti,
				checkList[i].Pnm, checkList[i].Ad, report.Score2Str(sims[i].Score), checkList[i].Abst)
			conclusion = append(conclusion, header)
			retconc = retconc + "专利" + strconv.Itoa(closeCount) + "是" + checkList[i].Cl + "\n"
			closeCount++
		}
	}
	scale := float64(closeCount-1) / float64(len(checkList))
	if scale > 0.5 {
		retconc = retconc + "而本专利是" + c.CL + report.DisclaimerTemplate + report.NegativeResult
	} else {
		retconc = retconc + "本专利是" + c.CL + report.DisclaimerTemplate + report.PositiveResult
	}
	realteCount := 0
	if len(checkList) > 0 {
		realteCount = len(checkList) - 1
	} else {
		realteCount = len(checkList)
	}

	data := model.Report{
		ReportName: fmt.Sprintf("%s 查新报告", c.Title),
		Type:       noveltyReportType,
	}
	err := e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}

	return &dto.NoveltyReportResp{
		Number:          uuid.New().String(),
		DepartName:      my_config.CurrentPatentConfig.NoveltyReportConfig.DepartName,
		ContactAddr:     my_config.CurrentPatentConfig.NoveltyReportConfig.ContactAddr,
		ZipCode:         my_config.CurrentPatentConfig.NoveltyReportConfig.ZipCode,
		ManagerName:     my_config.CurrentPatentConfig.NoveltyReportConfig.ManagerName,
		ManagerTel:      my_config.CurrentPatentConfig.NoveltyReportConfig.ManagerTel,
		ContactName:     my_config.CurrentPatentConfig.NoveltyReportConfig.ContactName,
		ContactTel:      my_config.CurrentPatentConfig.NoveltyReportConfig.ContactTel,
		Email:           my_config.CurrentPatentConfig.NoveltyReportConfig.Email,
		Database:        my_config.CurrentPatentConfig.NoveltyReportConfig.DataBase,
		PatentName:      c.Title,
		UserName:        c.Applicant,
		Institution:     c.Org,
		FinishData:      utils.FormatCurrentTime(),
		TechPoint:       c.CL,
		QueryWord:       searchList,
		QueryExpression: queryExpression,
		RelativeNum:     strconv.Itoa(realteCount),
		VeryRelativeNum: strconv.Itoa(closeCount - 1),
		SearchResult:    strings.Join(conclusion, "\n"),
		Conclusion:      retconc,
	}, nil
}
