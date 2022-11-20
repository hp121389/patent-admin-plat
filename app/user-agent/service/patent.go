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
	"os"
	"sort"
	"strings"

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
	//fmt.Println("pid在这里", c.PId)
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
	see := gse.GetResult(segments)
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
		resWords1 := gse.GetResult(segments1)
		result1 := RemoveStop(unique(resWords1))
		temp.score, _ = ts.Similarity(result1, result)
		//keywords := ts.Keywords(0.2, 0.5)
		//fmt.Println("keywords %d : %s", j, keywords)
		//temp.score = CosineSimilar(result1, result)
		sims = append(sims, temp)
	}
	n := len(sims)
	var conclusion []string
	for i := 0; i < n-1; i++ {
		maxNumIndex := i // 无序区第一个
		for j := i + 1; j < n; j++ {
			if sims[j].score > sims[maxNumIndex].score {
				maxNumIndex = j
			}
		}
		sims[i], sims[maxNumIndex] = sims[maxNumIndex], sims[i]
		list1[i], list1[maxNumIndex] = list1[maxNumIndex], list1[i]
		fmt.Println("\n申请号：", list1[i].PNM, "\n专利名称：", list1[i].TI, "\n相似度：", sims[i].score)
		if sims[i].score > 0.48 {
			conclusion = append(conclusion, "对比文件 ", ": ", list1[i].CL, "\n\n")
		}
	}
	conclusion = append(conclusion, "基于以上对比文件，本申请的区别特征在于：", model.CL, "因此，本专利具备新颖性和创造性")
	fmt.Println(conclusion)

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
		resWords := RemoveStop(gse.GetResult(segments1))
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
		tokens := RemoveStop(gse.GetResult(segments1))
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
		min, _ := minMaxKvSlice(vector)
		vector = filter(vector, func(v kv) bool {
			return (v.Value == min)
		})

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

//func CosineSimilar(srcWords, dstWords []string) float64 {
//	// get all words
//	allWordsMap := make(map[string]int, 0)
//	for _, word := range srcWords {
//		allWordsMap[word] += 1
//	}
//	for _, word := range dstWords {
//		allWordsMap[word] += 1
//	}
//
//	// stable the sort
//	allWordsSlice := make([]string, 0)
//	for word, _ := range allWordsMap {
//		allWordsSlice = append(allWordsSlice, word)
//	}
//
//	// assemble vector
//	srcVector := make([]int, len(allWordsSlice))
//	dstVector := make([]int, len(allWordsSlice))
//	for _, word := range srcWords {
//		if index := indexOfSclie(allWordsSlice, word); index != -1 {
//			srcVector[index] += 1
//		}
//	}
//	for _, word := range dstWords {
//		if index := indexOfSclie(allWordsSlice, word); index != -1 {
//			dstVector[index] += 1
//		}
//	}
//	//fmt.Printf("srcVector:%v\n", srcVector)
//	//fmt.Printf("dstVector:%v\n", dstVector)
//
//	// calc cos
//	numerator := float64(0)
//	srcSq := 0
//	dstSq := 0
//	for i, srcCount := range srcVector {
//		dstCount := dstVector[i]
//		numerator += float64(srcCount * dstCount)
//		srcSq += srcCount * srcCount
//		dstSq += dstCount * dstCount
//	}
//	denominator := math.Sqrt(float64(srcSq * dstSq))
//
//	return numerator / denominator
//}
//
//func indexOfSclie(ss []string, s string) (index int) {
//	index = -1
//	for k, v := range ss {
//		if s == v {
//			index = k
//			break
//		}
//	}
//
//	return
//}

func RemoveStop(unstop []string) []string {
	file, err := os.Open("E:/下载/testgit/app/user-agent/file1.txt")
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
