package service

import (
	"encoding/json"
	"fmt"
	s_cache "github.com/Gleiphir2769/s-cache"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/my_config"
	"go-admin/app/user-agent/service/charts"
	"go-admin/app/user-agent/service/dto"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	authUrl      = "http://www.innojoy.com/accountAuth.aspx"
	searchUrl    = "http://www.innojoy.com/service/patentSearch.aspx"
	statisticUrl = "http://www.innojoy.com/service/patentStat.aspx"

	searchListFields = "TI,AN,AD,PNM,PD,PA,PINN,CL,CD,AR,CLS"
	searchSortBy     = "-公开（公告）日,公开（公告）号"

	defaultCacheExpire     = time.Hour * 24
	defaultCleanupInterval = time.Minute
	defaultCacheCapacity   = 100000
)

var CurrentInnojoy *InnojoyClient
var innojoyCreateOnce sync.Once

func GetCurrentInnojoy() *InnojoyClient {
	innojoyCreateOnce.Do(func() {
		CurrentInnojoy = newInnojoyClient()
	})
	return CurrentInnojoy
}

type callback func() error

type InnojoyClient struct {
	email    string
	password string

	hc    *httpClient
	token string

	pc *pageCache
}

func newInnojoyClient() *InnojoyClient {
	return &InnojoyClient{
		email:    my_config.CurrentPatentConfig.InnojoyUser,
		password: my_config.CurrentPatentConfig.InnojoyPassword,
		hc:       newHttpClient(),
		pc:       newPageCache(),
	}
}

func (ic *InnojoyClient) autoLogin() error {
	req := &dto.LoginReq{UserConfig: dto.UserConfig{
		EMail:    ic.email,
		Password: ic.password,
	}}

	resp, err := ic.hc.Post(authUrl, req, nil)
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	loginRes := dto.LoginResp{}
	if err = json.Unmarshal(buf, &loginRes); err != nil {
		return err
	} else if loginRes.ReturnValue != 0 {
		return fmt.Errorf("login to innojoy failed: %s", loginRes.ErrorInfo)
	}

	ic.token = strings.Split(strings.Split(loginRes.Option, ":\"")[1], "\"")[0]

	return nil
}

func (ic *InnojoyClient) Search(req *dto.SimpleSearchReq, relatedPatents []models.UserPatent) (result []*dto.PatentDetail, err error) {
	sr := ic.parseSearchQuery(req.Query, req.DB, req.PageIndex, req.PageSize)
	res, err := ic.search(sr, ic.autoLogin)
	if err != nil {
		return nil, err
	}

	// mark focused or claimed
	if relatedPatents != nil {
		markRelation(res, relatedPatents)
	}

	return res, nil
}

func (ic *InnojoyClient) parseSearchQuery(query string, db string, pageIndex int, pageSize int) *dto.SearchReq {
	var guid string
	if pageIndex > 0 {
		guid = ic.pc.Get(query)
	}
	return &dto.SearchReq{
		Token: ic.token,
		PatentSearchConfig: &dto.PatentSearchConfig{
			GUID:      guid,
			Action:    "Search",
			Query:     query,
			Database:  db,
			Page:      strconv.Itoa(pageIndex),
			PageSize:  strconv.Itoa(pageSize),
			Sortby:    searchSortBy,
			FieldList: searchListFields,
		},
	}
}

func (ic *InnojoyClient) search(sr *dto.SearchReq, cb callback) (result []*dto.PatentDetail, err error) {
	var retried bool
	for {
		resp, err := ic.hc.Post(searchUrl, sr, nil)
		if err != nil {
			return nil, err
		}

		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		searchRes := dto.InnojoySearchResult{}
		if err = json.Unmarshal(buf, &searchRes); err != nil {
			return nil, err
		}

		if searchRes.ReturnValue != 0 {
			if retried {
				return nil, fmt.Errorf("patent search failed: %s", searchRes.ErrorInfo)
			}
			if err = cb(); err != nil {
				return nil, fmt.Errorf("seatch call callback error: %w", err)
			}
			// reset token
			sr.Token = ic.token
			retried = true
		} else {
			// refresh page GUID cache
			ic.pc.Put(sr.PatentSearchConfig.Query, searchRes.Option.GUID)
			// remove useless data
			refinePatentDetails(searchRes.Option.PatentList)
			return searchRes.Option.PatentList, nil
		}
	}
}

// refine patent title
func refinePatentDetails(pds []*dto.PatentDetail) {
	for _, pd := range pds {
		pd.Ti = strings.Split(pd.Ti, "[ZH]")[0]
		pd.Pa = strings.Split(pd.Pa, ";")[0]
		if len(strings.Split(pd.Ar, " ")) > 1 {
			temp := strings.Split(pd.Ar, " ")
			pd.Ar = temp[1]
		}
	}
}

type pageCache struct {
	cache *s_cache.Cache
}

func newPageCache() *pageCache {
	c := s_cache.NewCache(defaultCacheExpire, defaultCleanupInterval, s_cache.NewLRU(defaultCacheCapacity))
	return &pageCache{cache: c}
}

func (c *pageCache) Put(key string, guid string) {
	c.cache.Delete(key, nil)
	h := c.cache.Get(key, func() (size int, value s_cache.Value, d time.Duration) {
		return 1, guid, s_cache.DefaultExpiration
	})
	h.Release()
}

func (c *pageCache) Get(key string) string {
	h := c.cache.Get(key, nil)
	if h == nil {
		return ""
	}
	defer h.Release()
	return h.Value().(string)
}

func markRelation(res []*dto.PatentDetail, related []models.UserPatent) {
	rm := make(map[string]models.UserPatent, len(related))
	for _, rel := range related {
		rm[rel.PNM] = rel
	}
	for _, r := range res {
		if rel, ok := rm[r.Pnm]; ok {
			if rel.Type == dto.ClaimType {
				r.IsClaimed = true
			} else if rel.Type == dto.FocusType {
				r.IsFocused = true
			}
			r.PatentId = rel.PatentId
		}
	}
}

func (ic *InnojoyClient) GetChart(aid int, req *dto.SimpleSearchReq) (*dto.ChartProfile, error) {
	sr := ic.genStatisticQuery(aid, req)
	res, err := ic.statistic(sr, ic.autoLogin)
	if err != nil {
		return nil, err
	}
	c, err := charts.GetChart(aid)
	if err != nil {
		return nil, err
	}
	option, err := c.Serialize(res)
	if err != nil {
		return nil, err
	}
	return &dto.ChartProfile{
		Name:   c.Name(),
		Option: option,
	}, nil
}

func (ic *InnojoyClient) genStatisticQuery(aid int, req *dto.SimpleSearchReq) *dto.StatisticReq {
	return &dto.StatisticReq{
		Token: ic.token,
		PatentSearchConfig: &dto.PatentSearchConfig{
			Query:    req.Query,
			Database: req.DB,
			PageSize: "1",
		},
		AnalyseConfig: &dto.AnalyseConfig{AID: strconv.Itoa(aid)},
		Language:      "zh",
	}
}

func (ic *InnojoyClient) statistic(sr *dto.StatisticReq, cb callback) (result []byte, err error) {
	var retried bool
	for {
		resp, err := ic.hc.Post(statisticUrl, sr, nil)
		if err != nil {
			return nil, err
		}

		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		searchRes := dto.InnojoySearchResult{}
		if err = json.Unmarshal(buf, &searchRes); err != nil {
			return nil, err
		}

		if searchRes.ReturnValue != 0 {
			if retried {
				return nil, fmt.Errorf("patent search failed: %s", searchRes.ErrorInfo)
			}
			if err = cb(); err != nil {
				return nil, fmt.Errorf("search call callback error: %w", err)
			}
			// reset token
			sr.Token = ic.token
			retried = true
		} else {
			return buf, nil
		}
	}
}
