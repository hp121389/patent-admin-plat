package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/prometheus/common/log"
	"github.com/yanyiwu/gojieba"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
	cDto "go-admin/common/dto"
	"gorm.io/gorm"
	"sort"
	"strconv"
	"strings"
)

type Patent struct {
	service.Service
}

// MaxSimplifiedNodes is package relation graph max inventors included
const MaxSimplifiedNodes = 200

// Some useless Words of Patent
const uselessWords = "本发明"

// GetPage 获取Patent列表
func (e *Patent) GetPage(c *dto.PatentReq, list *[]models.Patent, count *int64) error {
	var err error
	var data models.Patent
	err = e.Orm.Model(&data).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetPatentsByIds 通过Id数组获取Patent对象列表
func (e *Patent) GetPatentsByIds(ids []int, count *int64) ([]models.Patent, error) {
	if len(ids) == 0 {
		return []models.Patent{}, nil
	}

	var err error
	patents := make([]models.Patent, 0)

	var patentData models.Patent
	err = e.Orm.Model(&patentData).
		Find(&patents, ids).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return nil, err
	}

	for i, pdData := range patents {
		pd := dto.PatentDetail{}
		err = json.Unmarshal([]byte(pdData.PatentProperties), &pd)
		if err != nil {
			return nil, err
		}
		patents[i].Price = pd.Idx * dto.PatentPriceBase
	}

	return patents, nil
}

func (e *Patent) GetPatentPagesByIds(ids []int, req dto.PatentPagesReq, count *int64) ([]models.Patent, error) {
	if len(ids) == 0 {
		return []models.Patent{}, nil
	}

	var err error
	patents := make([]models.Patent, 0)

	var patentData models.Patent
	err = e.Orm.Model(&patentData).
		Scopes(
			cDto.Paginate(req.GetPageSize(), req.GetPageIndex()),
		).
		Find(&patents, ids).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return nil, err
	}

	for i, pdData := range patents {
		pd := dto.PatentDetail{}
		err = json.Unmarshal([]byte(pdData.PatentProperties), &pd)
		if err != nil {
			return nil, err
		}
		patents[i].Price = pd.Idx * dto.PatentPriceBase
	}

	return patents, nil
}

func (e *Patent) FindPatentPages(ids []int, req dto.FindPatentPagesReq, count *int64) ([]models.Patent, error) {
	if len(ids) == 0 {
		return []models.Patent{}, nil
	}

	var err error
	patents := make([]models.Patent, 0)

	var patentData models.Patent
	err = e.Orm.Model(&patentData).
		Scopes(
			cDto.Paginate(req.GetPageSize(), req.GetPageIndex()),
		).
		Where("patent_properties LIKE ?", fmt.Sprintf("%%%s%%", req.Query)).
		Find(&patents, ids).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return nil, err
	}

	for i, pdData := range patents {
		pd := dto.PatentDetail{}
		err = json.Unmarshal([]byte(pdData.PatentProperties), &pd)
		if err != nil {
			return nil, err
		}
		patents[i].Price = pd.Idx * dto.PatentPriceBase
	}

	return patents, nil
}

// Get 获取Patent对象
func (e *Patent) Get(d *dto.PatentById, model *models.Patent) error {
	var err error
	db := e.Orm.First(model, d.PatentId)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看专利不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GeByPNM 获取Patent对象
func (e *Patent) GeByPNM(d *dto.PatentBriefInfo, model *models.Patent) error {
	//引用传递、函数名、形参、返回值
	var err error
	db := e.Orm.First(model, d.PNM)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看专利不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 添加Patent对象
func (e *Patent) Insert(c *dto.PatentReq) error {
	var err error
	var data models.Patent
	var i int64
	err = e.Orm.Model(&data).Where("PNM = ?", c.PNM).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("专利ID已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.GenerateList(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// InsertIfAbsent 返回值与上面不同
func (e *Patent) InsertIfAbsent(c *dto.PatentReq) (*models.Patent, error) {
	var err error
	var data models.Patent
	var i int64
	err = e.Orm.Model(&data).Where("PNM = ? OR patent_id = ?", c.PNM, c.PatentId).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	if i > 0 {
		err = e.Orm.Model(&data).Where("PNM = ?", c.PNM).First(&data).Error
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return nil, err
		}
		return &data, nil
	}
	c.GenerateList(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	return &data, nil
}

// UpdateLists 根据PatentId修改Patent对象
func (e *Patent) UpdateLists(c *dto.PatentReq) error {
	var err error
	var model models.Patent
	db := e.Orm.First(&model, c.PatentId)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update Patent error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("专利不存在")
	}
	c.GenerateList(&model)
	update := e.Orm.Model(&model).Where("patent_id = ?", &model.PatentId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update patent-info error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// Remove 根据id删除Patent
func (e *Patent) Remove(c *dto.PatentById) error {
	var err error
	var data models.Patent

	db := e.Orm.Delete(&data, c.PatentId)
	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("删除数据不存在")
		return err
	}
	return nil
}

// GetGraphByPatents 通过Patent数组获得专利发明人的关系图
func (e *Patent) GetGraphByPatents(simplifiedNodes []models.SimplifiedNode, Relations []int) (models.Graph, error) {
	var err error
	NodeList := make([]models.Node, 0)       //return node
	LinkList := make([]models.Link, 0)       //return link
	PreNodeList := make([]models.PreNode, 0) //Similar to node struct and need change some attributes type to become NodeList
	PreLinkList := make([]models.PreLink, 0) //Similar to link struct and need change some attributes type to become LinkList
	InventorGraph := models.Graph{}
	if len(simplifiedNodes) == 0 {
		err = fmt.Errorf("simplifiedNodes is null")
		return InventorGraph, err
	}
	StrongRelationInventors := MinResult(len(simplifiedNodes), 10) //chose the top10%(maximum is 10) inventors as StrongRelationInventors(must show)
	for i, inventor := range simplifiedNodes[0:StrongRelationInventors] {
		PreNodeList = append(PreNodeList, models.PreNode{NodeId: inventor.Id, NodeCategory: i})
	}
	//build top10% to top10% relationship
	MaxNumberOfLinks := 15 //the maximum find link
	MaxExpansion := 4      //every Node can extend the number of nodes
	NowLink := FindRelationFrequencyAndSort(Relations, simplifiedNodes[0:StrongRelationInventors], simplifiedNodes[0:StrongRelationInventors], MaxNumberOfLinks, MaxExpansion, MaxSimplifiedNodes)
	PreLinkList = append(PreLinkList, NowLink...)
	//build top10% to others relationship
	MaxNumberOfLinks = 200
	MaxExpansion = 5
	NowLink = FindRelationFrequencyAndSort(Relations, simplifiedNodes[0:StrongRelationInventors], simplifiedNodes[StrongRelationInventors:], MaxNumberOfLinks, MaxExpansion, MaxSimplifiedNodes)
	PreLinkList = append(PreLinkList, NowLink...)
	//add the extended nodes
	for _, i := range NowLink {
		if !simplifiedNodes[i.Target].InTheGraph {
			PreNodeList = append(PreNodeList, models.PreNode{NodeId: i.Target, NodeCategory: PreNodeList[i.Source].NodeCategory})
			simplifiedNodes[i.Target].InTheGraph = true
		}
	}
	//build others to others relationship
	SourceNode := make([]models.SimplifiedNode, 0)
	TargetNode := make([]models.SimplifiedNode, 0)
	for _, node := range PreNodeList[StrongRelationInventors:] {
		SourceNode = append(SourceNode, models.SimplifiedNode{Id: node.NodeId})
		TargetNode = append(SourceNode, models.SimplifiedNode{Id: node.NodeId})
	}
	MaxNumberOfLinks = 100
	MaxExpansion = 3
	NowLink = FindRelationFrequencyAndSort(Relations, SourceNode, TargetNode, MaxNumberOfLinks, MaxExpansion, MaxSimplifiedNodes)
	PreLinkList = append(PreLinkList, NowLink...)
	//deal the struct PreLink and PreNode
	MaxSizeofNode := 50
	for _, node := range PreNodeList {
		NodeList = append(NodeList, models.Node{
			NodeId:            strconv.FormatInt(int64(node.NodeId), 10),
			NodeName:          simplifiedNodes[node.NodeId].Name,
			NodeValue:         simplifiedNodes[node.NodeId].TheNumberOfPatents,
			NodeSymbolizeSize: float32((simplifiedNodes[node.NodeId].TheNumberOfPatents) * MaxSizeofNode / simplifiedNodes[0].TheNumberOfPatents),
			NodeCategory:      node.NodeCategory,
		})
	}
	for _, link := range PreLinkList {
		LinkList = append(LinkList, models.Link{
			Source: strconv.FormatInt(int64(link.Source), 10),
			Target: strconv.FormatInt(int64(link.Target), 10),
			Value:  link.Value,
		})
	}
	InventorGraph.Links = LinkList
	InventorGraph.Nodes = NodeList
	return InventorGraph, nil
}

// FindRelationFrequencyAndSort 建立边关系
func FindRelationFrequencyAndSort(relations []int, sources []models.SimplifiedNode, targets []models.SimplifiedNode, MaxNumberOfLinks int, MaxExtend int, MaxInventors int) []models.PreLink {
	LinkList := make([]models.PreLink, 0)
	LinkSum := 0
	LinkExtend := make(map[int]int)
	LinkReturnList := make([]models.PreLink, 0)
	LinkSearch := make(map[int]bool) //avoid duplicate link
	//init LinkExtend
	for _, i2 := range sources {
		LinkExtend[i2.Id] = 0
	}
	//find all Relation between sources and targets
	for _, source := range sources {
		for _, target := range targets {
			if source.Id == target.Id || relations[target.Id*MaxInventors+source.Id] == 0 && relations[source.Id*MaxInventors+target.Id] == 0 {
				continue
			}
			if _, ok := LinkSearch[source.Id*MaxInventors+target.Id]; ok {
				if _, ok1 := LinkSearch[target.Id*MaxInventors+source.Id]; ok1 {
					continue
				}
			}
			LinkSearch[source.Id*MaxInventors+target.Id] = true
			LinkSearch[target.Id*MaxInventors+source.Id] = true
			if source.Id < target.Id {
				LinkList = append(LinkList, models.PreLink{Source: source.Id, Target: target.Id, Value: relations[source.Id*MaxInventors+target.Id]})
			} else {
				LinkList = append(LinkList, models.PreLink{Source: source.Id, Target: target.Id, Value: relations[target.Id*MaxInventors+source.Id]})
			}
		}
	}
	//sort LinkList
	sort.Slice(LinkList, func(i, j int) bool {
		if LinkList[i].Value > LinkList[j].Value {
			return true
		}
		return false
	})
	//pick return links condition(MaxExtend,MaxNumberOfLinks)
	MaxNumberOfLinks = MinResult(MaxNumberOfLinks, len(LinkList))
	for _, link := range LinkList {
		if LinkSum >= MaxNumberOfLinks {
			break
		}
		if LinkExtend[link.Source] >= MaxExtend {
			continue
		}
		LinkSum++
		LinkExtend[link.Source]++
		LinkReturnList = append(LinkReturnList, link)
	}
	return LinkReturnList
}

// FindInventorsAndRelationsFromPatents --------------------------------------------------------------------------
// 通过patents数组查找patents数组中的发明人以及专利和发明人的关系
func (e *Patent) FindInventorsAndRelationsFromPatents(listPatents []models.Patent) ([]models.SimplifiedNode, []int, error) {

	ListInventors := make([]models.Inventor, 0)
	//find every patents' SimplifiedNode and count
	for z := 0; z < len(listPatents); z++ {
		patentDetail := dto.PatentDetail{}
		if err := json.Unmarshal([]byte(listPatents[z].PatentProperties), &patentDetail); err != nil {
			return nil, nil, err
		}

		raw := strings.Split(patentDetail.Inn, ";")
		// filter the English name, such as 沈航;阮辰晖;白光伟;SHEN HANG;RUAN CHENHUI;BAI GUANGWEI
		inventors := raw[:len(raw)/2]

		for i := 0; i < len(inventors); i++ {
			InventorExist := false
			for j := 0; j < len(ListInventors); j++ {
				if ListInventors[j].Name == inventors[i] {
					ListInventors[j].TheNumberOfPatents++
					InventorExist = true
					ListInventors[j].PatentsId = append(ListInventors[j].PatentsId, listPatents[z].PatentId)
					break
				}
			}
			if !InventorExist {
				NewPatents := make([]int, 0)
				ListInventors = append(ListInventors, models.Inventor{
					Name:               inventors[i],
					TheNumberOfPatents: 1,
					PatentsId:          append(NewPatents, listPatents[z].PatentId),
				})
			}
		}
	}
	//sort Inventors
	sort.Slice(ListInventors, func(i, j int) bool {
		if ListInventors[i].TheNumberOfPatents > ListInventors[j].TheNumberOfPatents {
			return true
		}
		return false
	})

	//write the id to Inventors
	for i := 0; i < len(ListInventors); i++ {
		ListInventors[i].Id = i
	}
	//create Relations
	NowInventorNumbers := MinResult(MaxSimplifiedNodes, len(ListInventors))
	ListRelations := make([]int, MaxSimplifiedNodes*MaxSimplifiedNodes)
	for i := 0; i < NowInventorNumbers; i++ {
		for j := i; j < NowInventorNumbers; j++ {
			var count int
			source := make(map[int]bool)
			for _, OneOfPatentId := range ListInventors[i].PatentsId {
				source[OneOfPatentId] = true

			}
			for _, OneOfPatentId := range ListInventors[j].PatentsId {
				if _, ok := source[OneOfPatentId]; ok {
					count++
				}

			}
			ListRelations[i*(MaxSimplifiedNodes)+j] = count
		}
	}

	//change preInventors to Inventors(delete preInventor->patents)
	ListSimplifiedNodes := make([]models.SimplifiedNode, 0)
	for _, i := range ListInventors {
		NowInventor := models.SimplifiedNode{Id: i.Id, Name: i.Name, TheNumberOfPatents: i.TheNumberOfPatents}
		ListSimplifiedNodes = append(ListSimplifiedNodes, NowInventor)
	}
	return ListSimplifiedNodes, ListRelations, nil
}

// FindKeywordsAndRelationsFromPatents --------------------------------------------------------------------------
// 通过patents数组查找patents数组中的关键字以及专利和关键字的关系
func (e *Patent) FindKeywordsAndRelationsFromPatents(listPatents []models.Patent) ([]models.SimplifiedNode, []int, error) {

	ListKeyWords := make([]models.KeyWord, 0)
	//find every patents' SimplifiedNode and count
	TiOfPatentsList := make([]string, 0)
	//keywordsList :=make([][]string,len(listPatents))
	for _, p := range listPatents {
		patentDetail := dto.PatentDetail{}
		if err := json.Unmarshal([]byte(p.PatentProperties), &patentDetail); err != nil {
			return nil, nil, err
		}
		TiOfPatentsList = append(TiOfPatentsList, patentDetail.Ti+patentDetail.Abst)
	}
	keywordsList := FindKeyWords(TiOfPatentsList)
	for z := 0; z < len(listPatents); z++ {
		for i := 0; i < len(keywordsList[z]); i++ {
			InventorExist := false
			for j := 0; j < len(ListKeyWords); j++ {
				if ListKeyWords[j].Name == keywordsList[z][i] {
					InventorExist = true
					if ListKeyWords[j].PatentsId[len(ListKeyWords[j].PatentsId)-1] != listPatents[z].PatentId {
						ListKeyWords[j].TheNumberOfPatents++
						ListKeyWords[j].PatentsId = append(ListKeyWords[j].PatentsId, listPatents[z].PatentId)
						break
					}
				}
			}
			if !InventorExist {
				NewPatents := make([]int, 0)
				ListKeyWords = append(ListKeyWords, models.KeyWord{
					Name:               keywordsList[z][i],
					TheNumberOfPatents: 1,
					PatentsId:          append(NewPatents, listPatents[z].PatentId),
				})
			}
		}
	}
	//sort Inventors
	sort.Slice(ListKeyWords, func(i, j int) bool {
		if ListKeyWords[i].TheNumberOfPatents > ListKeyWords[j].TheNumberOfPatents {
			return true
		}
		return false
	})

	//write the id to Inventors
	for i := 0; i < len(ListKeyWords); i++ {
		ListKeyWords[i].Id = i
	}
	//create Relations
	NowKeyWordsNumbers := MinResult(MaxSimplifiedNodes, len(ListKeyWords))
	ListRelations := make([]int, MaxSimplifiedNodes*MaxSimplifiedNodes)
	for i := 0; i < NowKeyWordsNumbers; i++ {
		for j := i; j < NowKeyWordsNumbers; j++ {
			var count int
			source := make(map[int]bool)
			for _, OneOfPatentId := range ListKeyWords[i].PatentsId {
				source[OneOfPatentId] = true

			}
			for _, OneOfPatentId := range ListKeyWords[j].PatentsId {
				if _, ok := source[OneOfPatentId]; ok {
					count++
				}

			}
			ListRelations[i*(MaxSimplifiedNodes)+j] = count
		}
	}

	//change preInventors to Inventors(delete preInventor->patents)
	ListSimplifiedNodes := make([]models.SimplifiedNode, 0)
	for _, i := range ListKeyWords {
		NowInventor := models.SimplifiedNode{Id: i.Id, Name: i.Name, TheNumberOfPatents: i.TheNumberOfPatents}
		ListSimplifiedNodes = append(ListSimplifiedNodes, NowInventor)
	}
	return ListSimplifiedNodes, ListRelations, nil
}

// FindKeyWords find the KeyWords from Sentence
func FindKeyWords(Sentences []string) [][]string {
	jieBa := gojieba.NewJieba()
	Results := make([][]string, 0)
	for _, s := range Sentences {
		TagReturn := jieBa.Tag(s)
		WordResult := make([]string, 0)
		for _, word := range TagReturn {
			nowWord := strings.Split(word, "/")
			if len(nowWord[0]) > 6 && nowWord[0] != uselessWords && nowWord[1][0] == 'n' {
				WordResult = append(WordResult, nowWord[0])
			}
		}
		Results = append(Results, WordResult)
	}
	return Results
}

// MinResult --------------------------------------------------------------------------------------------------------------------
func MinResult(a1 int, a2 int) int {
	if a1 >= a2 {
		return a2
	} else {
		return a1
	}
}
