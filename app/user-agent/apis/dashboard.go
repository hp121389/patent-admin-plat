package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
)

type Dashboard struct {
	api.Api
}

// GetDashboard
// @Summary Dashboard主页
// @Description  Dashboard主页
// @Tags Dashboard主页
// @Success 200 {object} dto.Dashboard
// @Router /apis/v1/user-agent/dashboard [get]
// @Security Bearer
func (e Dashboard) GetDashboard(c *gin.Context) {
	ups := service.UserPatent{}
	ps := service.Patent{}
	pks := service.Package{}
	rps := service.Report{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&ups.Service).
		MakeService(&ps.Service).
		MakeService(&pks.Service).
		MakeService(&rps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	userID := user.GetUserId(c)

	focusList := make([]models.UserPatent, 0)
	if err = ups.GetFocusLists(userID, &focusList); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	focusCount := len(focusList)

	// get competitors
	focusIds := make([]int, len(focusList))
	for _, fp := range focusList {
		focusIds = append(focusIds, fp.Id)
	}
	var _c int64
	focusPatents, err := ps.GetPatentsByIds(focusIds, &_c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	competitorNodes, _, err := ps.FindInventorsAndRelationsFromPatents(focusPatents)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	claimList := make([]models.UserPatent, 0)
	if err = ups.GetClaimLists(userID, &claimList); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	claimCount := len(claimList)

	// get collaborators
	claimIds := make([]int, len(claimList))
	for _, cp := range claimList {
		claimIds = append(claimIds, cp.PatentId)
	}
	var count int64
	claimPatents, err := ps.GetPatentsByIds(claimIds, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	collaboratorNodes, _, err := ps.FindInventorsAndRelationsFromPatents(claimPatents)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ids := make([]int, 0, len(claimList))
	for _, claim := range claimList {
		ids = append(ids, claim.PatentId)
	}
	var patentCount int64
	patents, err := ps.GetPatentsByIds(ids, &patentCount)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	patentStatus := make(map[string]int)
	publicationDates := dto.NewPublicationDates()
	var totalPrice int
	for _, p := range patents {
		pd := dto.PatentDetail{}
		if err = json.Unmarshal([]byte(p.PatentProperties), &pd); err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		if len(pd.Cls) == 0 {
			continue
		}
		if _, ok := patentStatus[pd.Cls]; ok {
			patentStatus[pd.Cls]++
		} else {
			patentStatus[pd.Cls] = 1
		}

		if len(pd.Ad) != 0 {
			publicationDates.AddYear(pd.Ad)
		}

		totalPrice += pd.Idx * dto.PatentPriceBase
	}

	pkgList := make([]models.Package, 0)
	listPkgReq := dto.PackageListReq{UserId: user.GetUserId(c)}
	err = pks.ListByUserId(&listPkgReq, &pkgList)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	pkgCount := len(pkgList)

	var reportCount int64
	err = rps.GetCountByUserID(user.GetUserId(c), &reportCount)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	res := &dto.Dashboard{
		PatentClaimCount:     int(claimCount),
		PatentFocusCount:     int(focusCount),
		PatentStatus:         patentStatus,
		PublicationDates:     publicationDates.List(),
		PackageCount:         pkgCount,
		PatentRecommendation: nil,
		ReportCount:          int(reportCount),
		PatentTotalPrice:     totalPrice,
		Collaborators:        covertNodesToResearchers(collaboratorNodes),
		Competitors:          covertNodesToResearchers(competitorNodes),
	}

	e.OK(res, "查询成功")
}

func covertNodesToResearchers(nodes []models.SimplifiedNode) []*dto.Researcher {
	res := make([]*dto.Researcher, 0, len(nodes))
	for _, n := range nodes {
		res = append(res, &dto.Researcher{
			Name:  n.Name,
			Times: n.TheNumberOfPatents,
		})
	}
	return res
}
