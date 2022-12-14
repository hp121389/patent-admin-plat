package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/go-ego/gse"
	"github.com/prometheus/common/log"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
	"gorm.io/gorm"
	"math"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	cDto "go-admin/common/dto"
)

type Patent struct {
	service.Service
}
type Simlilarity struct {
	count int
	words []string
	score float64
}

// GetPage 获取patent列表
func (e *Patent) GetPage(c *dto.PatentGetPageReq, list *[]models.Patent, count *int64) error {
	var err error
	var data models.Patent
	var model models.Patent
	var list1 []models.Patent
	//fmt.Println("id在这里？", c.GetPatentId())
	db := e.Orm.First(&model, c.GetPatentId())
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

	var separator = "|"
	var sentence = model.TI + model.CL
	var seg gse.Segmenter
	seg.LoadDict()
	segments := seg.Segment([]byte(sentence))
	see := GetResult(segments)
	resWords := RemoveStop(see)
	result := unique(resWords)
	var sqlse2 = "CONCAT_WS(\" \", TI, CL) REGEXP \"" + strings.Join(result, separator) + "\""
	fmt.Println(result)

	err = e.Orm.Model(&data).
		Scopes(
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).Where(sqlse2).
		Find(list).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	list1 = *list
	var totalinfo []string
	for j := 0; j < len(list1); j++ {
		totalinfo = append(totalinfo, list1[j].TI+"，"+list1[j].CL)
	}
	ts := New(totalinfo)

	//var count1 []int
	var sims []Simlilarity
	for j := 0; j < len(list1); j++ {
		var temp = Simlilarity{}
		for i := 0; i < len(result); i++ {
			if strings.Contains(list1[j].TI, result[i]) || strings.Contains(list1[j].CL, result[i]) {
				temp.count++
				temp.words = append(temp.words, result[i])
			}
		}
		segments1 := seg.Segment([]byte(list1[j].TI + list1[j].CL))
		resWords1 := GetResult(segments1)
		result1 := RemoveStop(unique(resWords1))
		temp.score, _ = ts.Similarity(result1, result)
		sims = append(sims, temp)
	}
	keywords := ts.Keywords(0.2, 0.8)
	keywords = unique(keywords)
	fmt.Println("keywords222222 ", keywords)
	fmt.Println("检索词：\n")
	var searchlist string
	var searchword = make([][]string, 50)
	for i := 0; i < len(keywords); i++ {
		searchlist += keywords[i] + " " + toString(getSimilar(keywords[i])) + "\n"
		//fmt.Println(keywords[i], " ", getSimilar(keywords[i]))
		temp := getSimilar(keywords[i])
		searchword[i] = make([]string, 0)
		searchword[i] = append(searchword[i], keywords[i])
		for j := 0; j < len(temp); j++ {
			searchword[i] = append(searchword[i], temp[j])
		}
	}
	searchtype := getSearchType(searchword)
	fmt.Println("检索式：", searchtype)
	n := len(sims)
	var conclusion []string
	var count1 = 1
	for i := 0; i < n-1; i++ {
		maxNumIndex := i // 无序区第一个
		for j := i + 1; j < n; j++ {
			if sims[j].score > sims[maxNumIndex].score {
				maxNumIndex = j
			}
		}
		sims[i], sims[maxNumIndex] = sims[maxNumIndex], sims[i]
		list1[i], list1[maxNumIndex] = list1[maxNumIndex], list1[i]
		if sims[i].score > 0.3 {
			temp := strconv.Itoa(count1) + ".申请人: " + list1[i].INN + "\n申请单位:" + list1[i].PA + "\n专利名称:" + list1[i].TI + "\n申请号：" + list1[i].PNM + "\n申请日：" + list1[i].AD + "\n简介：" + list1[i].CL + "\n"
			conclusion = append(conclusion, temp)
			count1++
		}
	}
	str1 := html()
	str1 = strings.Replace(str1, "number", GetRandomString(10), -1)
	str1 = strings.Replace(str1, "pname", model.TI, -1)
	str1 = strings.Replace(str1, "pearson", "北京邮电大学 胡泊", -1)
	str1 = strings.Replace(str1, "startdate", getTime(), -1)
	str1 = strings.Replace(str1, "institution", "教育部科技查新工作站", -1)
	str1 = strings.Replace(str1, "finishdate", getTime(), -1)
	str1 = strings.Replace(str1, "cname", model.TI, -1)
	str1 = strings.Replace(str1, "telepoint", toHtml(model.CLAIMS), -1)
	str1 = strings.Replace(str1, "retWord", toHtml(searchlist), -1)
	str1 = strings.Replace(str1, "retType", toHtml(searchtype), -1)
	str1 = strings.Replace(str1, "num1", strconv.Itoa(len(list1)), -1)
	str1 = strings.Replace(str1, "num2", strconv.Itoa(count1-1), -1)
	str1 = strings.Replace(str1, "retResult", toHtml(toString2(conclusion)), -1)
	str1 = strings.Replace(str1, "retConclusion", toHtml(model.CL), -1)
	fileName := "./app/user-agent/mytest.html"
	dstFile, err := os.Create(fileName)
	defer dstFile.Close()
	dstFile.WriteString(str1 + "\n")
	fmt.Println("写入文档" + fileName + "成功!")

	//fmt.Println("str1在这里！", str1)
	list = &list1
	return nil
}

type kv struct {
	Key   string
	Value float64
}

// TextSimilarity is a struct containing internal
// data to be re-used by the package.
type TextSimilarity struct {
	corpus            []string
	documents         []string
	documentFrequency map[string]int
}

// Option type describes functional options that
// allow modification of the internals of TextSimilarity
// before initialization. They are optional, and not using them
// allows you to use the defaults.
type Option func(TextSimilarity) TextSimilarity

// Cosine returns the Cosine Similarity between two vectors.
func Cosine(a, b []float64) (float64, error) {
	count := 0
	lengthA := len(a)
	lengthB := len(b)
	if lengthA > lengthB {
		count = lengthA
	} else {
		count = lengthB
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= lengthA {
			s2 += math.Pow(b[k], 2)
			continue
		}
		if k >= lengthB {
			s1 += math.Pow(a[k], 2)
			continue
		}
		sumA += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("null vector")
	}
	return sumA / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}

func count(key string, a []string) int {
	count := 0
	for _, s := range a {
		if key == s {
			count = count + 1
		}
	}
	return count
}

func getTime() string {
	Year := time.Now().Year()
	Month := int(time.Now().Month())
	Day := time.Now().Day()
	time := strconv.Itoa(Year) + "年" + strconv.Itoa(Month) + "月" + strconv.Itoa(Day) + "日"
	return time
}
func GetRandomString(l int) string {
	str := "123456789ABCDEFGHIJKLMNPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	ok1, _ := regexp.MatchString(".[1|2|3|4|5|6|7|8|9]", string(result))
	ok2, _ := regexp.MatchString(".[Z|X|C|V|B|N|M|A|S|D|F|G|H|J|K|L|Q|W|E|R|T|Y|U|I|P]", string(result))
	if ok1 && ok2 {
		return string(result)
	} else {
		return GetRandomString(l)
	}

}

func tfidf(v string, tokens []string, n int, documentFrequency map[string]int) float64 {
	tf := float64(count(v, tokens)) / float64(documentFrequency[v])
	idf := math.Log(float64(n) / (float64(documentFrequency[v])))
	return tf * idf
}

func union(a, b []string) []string {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}

func minMaxKvSlice(s []kv) (min, max float64) {
	min = math.Inf(0)
	max = math.Inf(-1)
	for _, v := range s {
		max = math.Max(v.Value, max)
		min = math.Min(v.Value, min)
	}
	return min, max
}

func filter(vs []kv, f func(kv) bool) []kv {
	var vsf []kv
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// New accepts a slice of documents and
// creates the internal corpus and document frequency mapping.
func New(documents []string) *TextSimilarity {
	var (
		allTokens []string
	)

	ts := TextSimilarity{
		documents: documents,
	}

	ts.documentFrequency = map[string]int{}
	var seg gse.Segmenter
	seg.LoadDict()
	for _, doc := range documents {

		segments1 := seg.Segment([]byte(doc))
		resWords := RemoveStop(GetResult(segments1))
		allTokens = append(allTokens, resWords...)
	}

	// Generate a corpus.
	for _, t := range allTokens {
		if ts.documentFrequency[t] == 0 {
			ts.documentFrequency[t] = 1
			ts.corpus = append(ts.corpus, t)
		} else {
			ts.documentFrequency[t] = ts.documentFrequency[t] + 1
		}
	}

	return &ts
}

// Similarity returns the cosine similarity between two documents using
// Tf-Idf vectorization using the corpus.
func (ts *TextSimilarity) Similarity(a, b []string) (float64, error) {
	combinedTokens := union(a, b)
	// Populate the vectors using frequency in the corpus.
	n := len(combinedTokens)
	vectorA := make([]float64, n)
	vectorB := make([]float64, n)
	for k, v := range combinedTokens {
		vectorA[k] = tfidf(v, a, n, ts.documentFrequency)
		vectorB[k] = tfidf(v, b, n, ts.documentFrequency)
	}

	similarity, err := Cosine(vectorA, vectorB)
	if err != nil {
		return 0.0, err
	}
	return similarity, nil
}

// Keywords accepts thresholds, which can be used to filter keyswords that
// are either they are too common or too unique and returns a sorted list of
// keywords (index 0 being the lower tf-idf score). Play with the thresholds
// according to your corpus.
func (ts *TextSimilarity) Keywords(threshLower, threshUpper float64) []string {
	var (
		docKeywords = []kv{}
		result      = []string{}
	)
	var seg gse.Segmenter
	seg.LoadDict()
	for _, doc := range ts.documents {
		segments1 := seg.Segment([]byte(doc))
		tokens := RemoveStop(GetResult(segments1))
		n := len(tokens)
		mapper := map[string]float64{}

		for _, v := range tokens {
			val := tfidf(v, tokens, n, ts.documentFrequency)
			mapper[v] = val
		}

		// Convert to a kv pair for convenience.
		i := 0
		vector := make([]kv, len(mapper))
		for k, v := range mapper {
			vector[i] = kv{
				Key:   k,
				Value: v,
			}
			i++
		}

		// Filter tf-idf, using threshold.
		vector = filter(vector, func(v kv) bool {
			return v.Value >= threshLower && v.Value <= threshUpper
		})

		// Select the most common words relative to the corpus for this doc.

		docKeywords = append(docKeywords, vector...)
	}

	// Sort the vector based on tf-idf scores
	sort.Slice(docKeywords, func(i, j int) bool {
		return docKeywords[i].Value < docKeywords[j].Value
	})

	// Convert back to slice.
	for _, word := range docKeywords {
		result = append(result, word.Key)
	}
	return result
}

func unique(resWords []string) []string {
	result := make([]string, len(resWords))
	result[0] = resWords[0]
	result_idx := 1
	for i := 0; i < len(resWords); i++ {
		is_repeat := false
		for j := 0; j < len(result); j++ {
			if resWords[i] == result[j] {
				is_repeat = true
				break
			}
		}
		if !is_repeat {
			result[result_idx] = resWords[i]
			result_idx++
		}
	}
	return result[:result_idx]
}

func html() string {
	str := "<p>\n    报告编号：number\n</p>\n<p>\n    &nbsp;\n</p>\n<p>\n    &nbsp;\n</p>\n<p>\n    &nbsp;\n</p>\n<p style=\"text-align:center\">\n    <strong><span style=\"font-size:29px;font-family: 宋体\">科 技 查 新 报 告</span></strong>\n</p>\n<p style=\"text-align:center\">\n    <strong><span style=\"font-size:29px;font-family:宋体\">&nbsp;</span></strong>\n</p>\n<p style=\"text-align:center\">\n    <strong><span style=\"font-size:29px;font-family:宋体\">&nbsp;</span></strong>\n</p>\n<p style=\"margin-top:16px;margin-right:0;margin-bottom:16px;margin-left:140px;text-align:left;line-height:150%\">\n    <strong><span style=\"font-size:19px;line-height: 150%\">项目名称：&nbsp;&nbsp;&nbsp; </span></strong><span style=\"font-size:16px;line-height:150%\">pname</span>\n</p>\n<p style=\"margin-top:16px;margin-right:0;margin-bottom:16px;margin-left:140px;text-align:left;line-height:150%\">\n    <strong><span style=\"font-size:19px;line-height: 150%\">委 托 人 ：&nbsp;&nbsp;&nbsp; </span></strong><span style=\"font-size:16px;line-height:150%\">pearson</span>\n</p>\n<p style=\"margin-top:16px;margin-right:0;margin-bottom:16px;margin-left:140px;text-align:left;line-height:150%\">\n    <strong><span style=\"font-size:19px;line-height: 150%\">委托日期：&nbsp;&nbsp;&nbsp; </span></strong><span style=\"font-size:16px;line-height:150%\">startdate</span>\n</p>\n<p style=\"margin-top:16px;margin-right:0;margin-bottom:16px;margin-left:140px;text-align:left;line-height:150%\">\n    <strong><span style=\"font-size:19px;line-height: 150%\">查新机构：&nbsp;&nbsp;&nbsp; </span></strong><span style=\"font-size:16px;line-height:150%\">institution</span>\n</p>\n<p style=\"margin-top:16px;margin-right:0;margin-bottom:16px;margin-left:140px;text-align:left;line-height:150%\">\n    <strong><span style=\"font-size:19px;line-height: 150%\">完成日期：&nbsp;&nbsp;&nbsp; </span></strong><span style=\"font-size:16px;line-height:150%\">finishdate</span>\n</p>\n<p style=\"text-align:left\">\n    <strong><span style=\"font-size:21px\">&nbsp;</span></strong>\n</p>\n<p style=\"text-align:left\">\n    <strong><span style=\"font-size:16px\">&nbsp;</span></strong>\n</p>\n<p style=\"text-align:left\">\n    <strong><span style=\"font-size:16px\">&nbsp;</span></strong>\n</p>\n<p style=\"text-align:left\">\n    <strong><span style=\"font-size:16px\">&nbsp;</span></strong>\n</p>\n<p style=\"text-align:left\">\n    <strong><span style=\"font-size:21px\">&nbsp;</span></strong>\n</p>\n<p style=\"text-align:left\">\n    <strong><span style=\"font-size:21px\">&nbsp;</span></strong>\n</p>\n<p style=\"text-align:left\">\n    <strong><span style=\"font-size:21px\">&nbsp;</span></strong>\n</p>\n<p style=\"text-align:center\">\n    <strong><span style=\"font-size:16px\">教育部科技发展中心</span></strong>\n</p>\n<p style=\"text-align:center\">\n    <span style=\"font-size:16px\">二O一三年制</span>\n</p>\n<p>\n    <br/>\n</p>\n<table cellspacing=\"0\" cellpadding=\"0\">\n    <tbody>\n        <tr style=\";height:36px\" class=\"firstRow\">\n            <td rowspan=\"2\" style=\"border: 1px solid windowtext; padding: 0px 7px; word-break: break-all;\" width=\"77\" height=\"36\">\n                <p style=\"text-align:justify;text-justify:distribute-all-lines\">\n                    查新项目\n                </p>\n                <p style=\"text-align:justify;text-justify:distribute-all-lines\">\n                    名称\n                </p>\n            </td>\n            <td colspan=\"6\" style=\"border-color: windowtext windowtext windowtext currentcolor; border-style: solid solid solid none; border-width: 1px 1px 1px medium; border-image: none 100% / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;\" width=\"476\" height=\"36\">\n                <p>\n                    中文：cname\n                </p>\n            </td>\n        </tr>\n        <tr style=\";height:36px\">\n            <td colspan=\"6\" style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;\" width=\"476\" height=\"36\">\n                <p>\n                    英文：略\n                </p>\n            </td>\n        </tr>\n        <tr style=\";height:23px\">\n            <td rowspan=\"5\" style=\"border-color: currentcolor windowtext windowtext; border-style: none solid solid; border-width: medium 1px 1px; border-image: none 100% / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;\" width=\"77\" height=\"23\">\n                <p style=\"text-align:center\">\n                    查新机构\n                </p>\n            </td>\n            <td style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px;\" width=\"75\" height=\"23\">\n                <p>\n                    名称\n                </p>\n            </td>\n            <td colspan=\"5\" style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;\" width=\"401\" height=\"23\">\n                insName<br/>\n            </td>\n        </tr>\n        <tr style=\";height:23px\">\n            <td style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px;\" width=\"75\" height=\"23\">\n                <p>\n                    通信地址\n                </p>\n            </td>\n            <td colspan=\"3\" style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;\" width=\"232\" height=\"23\">\n                insAddress<br/>\n            </td>\n            <td style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px;\" width=\"67\" height=\"23\">\n                <p>\n                    邮政编码\n                </p>\n            </td>\n            <td style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;\" width=\"102\" height=\"23\">\n                insPost<br/>\n            </td>\n        </tr>\n        <tr style=\";height:17px\">\n            <td style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;\" width=\"75\" height=\"17\">\n                <p>\n                    负责人\n                </p>\n            </td>\n            <td style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;\" width=\"93\" height=\"17\">\n                pic<br/>\n            </td>\n            <td style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px;\" width=\"58\" height=\"17\">\n                <p>\n                    电话\n                </p>\n            </td>\n            <td colspan=\"3\" style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;\" width=\"250\" height=\"17\">\n                tele1<br/>\n            </td>\n        </tr>\n        <tr style=\";height:16px\">\n            <td style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;\" width=\"75\" height=\"16\">\n                <p>\n                    联系人\n                </p>\n            </td>\n            <td style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;\" width=\"93\" height=\"16\">\n                ptc<br/>\n            </td>\n            <td style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px;\" width=\"58\" height=\"16\">\n                <p>\n                    电话\n                </p>\n            </td>\n            <td colspan=\"3\" style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;\" width=\"250\" height=\"16\">\n                tele2<br/>\n            </td>\n        </tr>\n        <tr style=\";height:27px\">\n            <td style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px;\" width=\"75\" height=\"27\">\n                <p>\n                    电子邮箱\n                </p>\n            </td>\n            <td colspan=\"5\" style=\"border-color: currentcolor windowtext windowtext currentcolor; border-style: none solid solid none; border-width: medium 1px 1px medium; padding: 0px 7px; word-break: break-all;\" width=\"401\" height=\"27\">\n                insEamil<br/>\n            </td>\n        </tr>\n        <tr style=\";height:107px\">\n            <td colspan=\"7\" style=\"border-color: currentcolor windowtext windowtext; border-style: none solid solid; border-width: medium 1px 1px; border-image: none 100% / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;\" width=\"553\" valign=\"top\" height=\"107\">\n                <p>\n                    一、项目的科学技术要点\n                </p> <p style=\"text-indent:28px\">&nbsp;telepoint\n       </p>      </td>\n        </tr>\n        <tr style=\";height:107px\">\n            <td colspan=\"7\" style=\"border-color: currentcolor windowtext windowtext; border-style: none solid solid; border-width: medium 1px 1px; border-image: none 100% / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;\" width=\"553\" valign=\"top\" height=\"107\">\n                <p>\n                    二、专利检索范围及检索策略\n                </p>\n                <p style=\"text-indent:28px\">\n                    检索的中文数据库\n                </p>\n                <p style=\"text-indent:28px\">\n                    &nbsp;cDataBase\n                </p>\n                <p style=\"text-indent:28px\">\n                    &nbsp;\n                </p>\n                <p style=\"text-indent:28px\">\n                    检索词\n                </p>\n                              retWord\n                              <p style=\"text-indent:28px\">\n                    &nbsp;\n                </p>\n                <p style=\"text-indent:28px\">\n                    检索式\n                </p>\n                                   &nbsp;retType\n              &nbsp;\n                <p style=\"text-indent:28px\">\n                    &nbsp;\n                </p>\n            </td>\n        </tr>\n        <tr style=\";height:89px\">\n            <td colspan=\"7\" style=\"border-color: currentcolor windowtext windowtext; border-style: none solid solid; border-width: medium 1px 1px; border-image: none 100% / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;\" width=\"553\" valign=\"top\" height=\"89\">\n                <p>\n                    三、检索结果\n                </p>\n                <p style=\"text-indent:28px\">\n                    依据上专利检索范围和检索式，共检索出相专利 num1 项，其中密切相关专利 num2 项，题录为：\n                </p>&nbsp;retResult\n                <p style=\"text-indent:28px\">\n                    &nbsp;\n                </p>\n            </td>\n        </tr>\n        <tr style=\";height:89px\">\n            <td colspan=\"7\" style=\"border-color: currentcolor windowtext windowtext; border-style: none solid solid; border-width: medium 1px 1px; border-image: none 100% / 1 / 0 stretch; padding: 0px 7px; word-break: break-all;\" width=\"553\" valign=\"top\" height=\"89\">\n                <p>\n                    四、查新结论\n                </p>\n                                        <p style=\"text-indent:28px\">   经对检出的相关文献进行阅读、分析、对比，结论如下：</p> <p style=\"text-indent:28px\"> retConclusion </p> <p style=\"text-indent:28px\"> 本次查新在国内公开发表的中文文献中，尚未见有与本项目研究内容一致的文献报道，本项目研究内容在国内外具备新颖性。<br/>\n                </p>&nbsp;\n                <p>\n                    &nbsp;\n                </p>\n            </td>\n        </tr>\n    </tbody>\n</table>\n<p>\n    <br/>\n</p>"
	return str
}

func RemoveStop(unstop []string) []string {
	file, err := os.Open("./app/user-agent/file1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	stops := make([]string, 0)
	result := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stop := strings.Split(line, "\n")
		stops = append(stops, stop[0])
	}
	for i := 0; i < len(unstop); i++ {
		same := 0
		for j := 0; j < len(stops); j++ {
			if stops[j] == unstop[i] {
				same = 1
				break
			}
		}
		if same == 0 {
			result = append(result, unstop[i])
		}
	}
	return result
}

func getSimilar(word string) []string {
	var max = 4
	file, err := os.Open("./app/user-agent/cilin.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	similarwords := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		similarword := strings.Split(line, "\n")
		similarwords = append(similarwords, similarword[0])
	}
	for i := 0; i < len(similarwords); i++ {
		result := make([]string, 0)
		temp := strings.Split(similarwords[i], " ")
		var include = false
		var count = 0
		for j := 1; j < len(temp) && count < max; j++ {
			if temp[j] == word {
				include = true
			} else {
				result = append(result, temp[j])
				count++
			}
		}
		if include {
			return result
		}
	}
	return nil
}

func toString(list []string) string {
	var result string
	for i := 0; i < len(list); i++ {
		result = result + list[i] + " "
	}
	return result
}
func toString2(list []string) string {
	var result string
	for i := 0; i < len(list); i++ {
		result = result + list[i] + "\n"
	}
	return result
}

func GetResult(segs []gse.Segment, searchMode ...bool) []string {
	var mode bool
	var output []string
	if len(searchMode) > 0 {
		mode = searchMode[0]
	}

	if mode {
		for _, seg := range segs {
			output = append(output, seg.Token().Text())
		}
		return output
	}
	partOfSpeech := []string{"v", "n", "vn", "x", "an", "nz", "a", "l", "ns"}

	for _, seg := range segs {
		for i := 0; i < len(partOfSpeech); i++ {
			if seg.Token().Pos() == partOfSpeech[i] {
				output = append(output, seg.Token().Text())
				break
			}
		}
	}

	return output
}

func toHtml(word string) string {
	result := strings.Replace(word, "\n", "<br>", -1)
	return result

}

func getSearchType(word [][]string) string {
	resultt := ""
	num := 0
	for num = 0; word[num] != nil; num++ {
	}
	for k := 0; k < 3; k++ {
		resultt += strconv.Itoa(k+1) + ".  "
		for i := k * num / 3; i < (k+1)*num/3; i++ {
			temp := ""
			if word[i] != nil {
				temp += " ( "
				for j := 0; j < len(word[i]); j++ {
					if j < len(word[i])-1 {
						temp += word[i][j] + " OR "
					} else {
						temp += word[i][j]
					}
				}
				temp += " ) "
				if i < (k+1)*num/3-1 {
					resultt += temp + " AND "
				} else {
					resultt += temp
				}
			}
		}
		resultt += "\n"
	}
	return resultt

}

// Get 获取Patent对象
func (e *Patent) Get(d *dto.PatentById, model *models.Patent) error {
	//引用传递、函数名、形参、返回值
	var err error
	db := e.Orm.First(model, d.GetPatentId())
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

// Remove 根据专利id删除Patent
func (e *Patent) Remove(c *dto.PatentById) error {
	var err error
	var data models.Patent

	db := e.Orm.Delete(&data, c.GetPatentId())

	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// UpdateLists 根据PatentId修改Patent对象
func (e *Patent) UpdateLists(c *dto.PatentUpdateReq) error {
	var err error
	var model models.Patent
	db := e.Orm.First(&model, c.GetPatentId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update Patent error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

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

// Insert 根据PatentId 创建Patent对象
func (e *Patent) Insert(c *dto.PatentInsertReq) error {
	var err error
	var data models.Patent
	var i int64
	err = e.Orm.Model(&data).Where("patent_id = ?", c.PatentId).Count(&i).Error
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

// InsertIfAbsent 根据PatentId 创建Patent对象
func (e *Patent) InsertIfAbsent(c *dto.PatentInsertReq) (int, error) {
	var err error
	var data models.Patent
	var i int64
	err = e.Orm.Model(&data).Where("PNM = ?", c.PNM).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, err
	}
	if i > 0 {
		err = e.Orm.Model(&data).Where("PNM = ?", c.PNM).First(&data).Error
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return 0, err
		}
		return data.PatentId, nil
	}
	c.GenerateList(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, err
	}
	return data.PatentId, nil
}

// RemoveClaim 取消认领
func (e *Patent) RemoveClaim(c *dto.UserPatentObject) error {
	var err error
	var data models.UserPatent

	db := e.Orm.Where("Patent_Id = ? AND User_Id = ? AND Type = ?", c.PatentId, c.UserId, dto.ClaimType).
		Delete(&data)

	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// RemoveFocus 取消关注
func (e *Patent) RemoveFocus(c *dto.UserPatentObject) error {
	var err error
	var data models.UserPatent

	db := e.Orm.Where("Patent_Id = ? AND User_Id = ? AND Type = ?", c.PatentId, c.UserId, dto.FocusType).
		Delete(&data)

	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

//// InsertCollectionRelationship 创建关注关系
//func (e *UserPatent) InsertCollectionRelationship(c *dto.UserPatentObject) error {
//	var err error
//	var data models.UserPatent
//	var i int64
//	err = e.Orm.Model(&data).Where("Patent_Id = ? AND User_Id = ? AND Type = ?", c.PatentId, c.UserId, c.Type).
//		Count(&i).Error
//	if err != nil {
//		e.Log.Errorf("db error: %s", err)
//		return err
//	}
//	if i > 0 {
//		err := errors.New("关系已存在！")
//		e.Log.Errorf("db error: %s", err)
//		return err
//	}
//
//	c.GenerateUserPatent(&data)
//	c.Type = "关注"
//
//	err = e.Orm.Create(&data).Error
//	if err != nil {
//		e.Log.Errorf("db error: %s", err)
//		return err
//	}
//	return nil
//}

//// UpdateUserPatent
//func (e *UserPatent) UpdateUserPatent(c *dto.UpDateUserPatentObject) error {
//	var err error
//	var model models.UserPatent
//	var i int64
//
//	ids := e.Orm.Model(&model).Where("Patent_Id = ? AND User_Id = ? ", c.PatentId, c.UserId).First(&model).Count(&i)
//
//	fmt.Println("一共有", i, "个专利id为", c.PatentId, "且用户是", c.UserId, "的关系")
//
//	if i == 2 {
//		//先按照条件找到用户对应的专利，然后修改，且只找一个。
//		//如果一个用户即关注又认领了一个专利怎么办呢 ,model不是数组，只是一个model
//		return errors.New("您已同时认领和关注该专利！")
//	}
//
//	err = ids.Error
//
//	db := e.Orm.Model(&model).Where("Patent_Id = ? AND User_Id = ? ", c.PatentId, c.UserId).
//		First(&model)
//
//	if err = db.Error; err != nil {
//		e.Log.Errorf("Service Update User-Patent error: %s", err)
//		return err
//	}
//	if db.RowsAffected == 0 {
//		return errors.New("无权更新该数据")
//	}
//
//	c.GenerateUserPatent(&model)
//
//	update := e.Orm.Model(&model).Updates(&model)
//	if err = update.Error; err != nil {
//		e.Log.Errorf("db error: %s", err)
//		return err
//	}
//	if update.RowsAffected == 0 {
//		err = errors.New("update patent-info error maybe you dont need update or record not exist")
//		log.Warnf("db update error")
//		return err
//	}
//	return nil
//}

// GetClaimLists 通过UserId获得PatentId列表
func (e *Patent) GetClaimLists(c *dto.UserPatentGetPageReq, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).
		Where("Type = ? AND User_Id = ?", "认领", c.GetUserId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetCollectionLists 通过UserId获得PatentId列表
func (e *Patent) GetCollectionLists(c *dto.UserPatentGetPageReq, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).
		Where("Type = ? AND User_Id = ?", "关注", c.GetUserId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetPatentPagesByIds 获取patent列表
func (e *Patent) GetPatentPagesByIds(d *dto.PatentsByIdsForRelationshipUsers, list *[]models.Patent, count *int64) error {
	var err error
	var ids []int = d.GetPatentId()
	for i := 0; i < len(ids); i++ {
		if ids[i] != 0 {
			var data1 models.Patent
			err = e.Orm.Model(&data1).
				Where("Patent_Id = ? ", ids[i]).
				First(&data1).Limit(-1).Offset(-1).
				Count(count).Error
			*list = append(*list, data1)
			if err != nil {
				e.Log.Errorf("db error:%s", err)
				return err
			}
		}
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Inserts relationship between user and patent
func (e *Patent) Inserts(c *dto.UserPatentObject) error {
	var err error
	var data models.UserPatent
	var i int64
	err = e.Orm.Model(&data).Where("Patent_Id = ? AND User_Id = ? AND Type = ?", c.PatentId, c.UserId, c.Type).
		Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err = fmt.Errorf("%w, (p:%d, u:%d, t:%s) existed", ErrConflictBindPatent, c.PatentId, c.UserId, c.Type)
		e.Log.Errorf("db error: %s", err)
		return err
	}

	c.GenerateUserPatent(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
