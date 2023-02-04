package report

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/go-ego/gse"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var synonymsDict map[string][]string
var ignoreWordsDict map[string]struct{}

const similarMaxLen = 4

var Seg gse.Segmenter

func init() {
	ignoreDict, err := os.Open("./app/user-agent/service/dict/ignore_words_dict.txt")
	if err != nil {
		panic(err)
	}
	defer ignoreDict.Close()
	ignoreWordsDict = make(map[string]struct{})
	br := bufio.NewReader(ignoreDict)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		ignoreWordsDict[string(line)] = struct{}{}
	}

	file, err := os.Open("./app/user-agent/service/dict/synonyms_dict.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	synonymsDict = make(map[string][]string)
	br = bufio.NewReader(file)
	for {
		bs, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else if len(bs) == 0 {
			break
		}
		line := string(bs)
		segs := strings.Split(line, "=")
		if len(segs) < 2 {
			continue
		}

		words := segs[1]
		allWords := make([]string, 0)
		for _, word := range strings.Split(words, " ") {
			if len(word) == 0 {
				continue
			}
			allWords = append(allWords, word)
		}

		cnt := len(allWords)
		for i, word := range allWords {
			if i == 0 {
				synonymsDict[word] = allWords[1:cnt]
				continue
			} else if i == cnt-1 {
				synonymsDict[word] = allWords[0 : cnt-1]
				continue
			}
			synonymsDict[word] = append(allWords[0:i], allWords[i+1:cnt]...)
		}
	}

	Seg.LoadDict()
}

type kv struct {
	Key   string
	Value float64
}

type Similarity struct {
	Count int
	Words []string
	Score float64
}

// TextSimilarity is a struct containing internal data to be re-used by the package.
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
	if documentFrequency[v] == 0 {
		return 0
	}
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

func filter(vs []kv, f func(kv) bool) []kv {
	var vsf []kv
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// NewTextSimilarity accepts a slice of documents and
// creates the internal corpus and document frequency mapping.
func NewTextSimilarity(documents []string) *TextSimilarity {
	var (
		allTokens []string
	)

	ts := TextSimilarity{
		documents: documents,
	}

	ts.documentFrequency = map[string]int{}
	for _, doc := range documents {
		segments1 := Seg.Segment([]byte(doc))
		resWords := RemoveStop(GetResult(segments1, 0))
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
// are either they are too common or too Unique and returns a sorted list of
// keywords (index 0 being the lower tf-idf Score). Play with the thresholds
// according to your corpus.
func (ts *TextSimilarity) Keywords(threshLower, threshUpper float64, pattern int) []string {
	var (
		docKeywords = []kv{}
		result      = []string{}
	)
	for _, doc := range ts.documents {
		segments1 := Seg.Segment([]byte(doc))
		tokens := RemoveStop(GetResult(segments1, 0))
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

		// Select the most common Words relative to the corpus for this doc.

		if pattern != 0 {
			sort.Slice(vector, func(i, j int) bool {
				return vector[i].Value < vector[j].Value
			})
			docKeywords = append(docKeywords, vector...)
			break
		}
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

func Unique(resWords []string) []string {
	result := make([]string, len(resWords))
	result[0] = resWords[0]
	resultIdx := 1
	for i := 0; i < len(resWords); i++ {
		isRepeat := false
		for j := 0; j < len(result); j++ {
			if resWords[i] == result[j] {
				isRepeat = true
				break
			}
		}
		if !isRepeat {
			result[resultIdx] = resWords[i]
			resultIdx++
		}
	}
	return result[:resultIdx]
}

func RemoveStop(unstop []string) []string {
	result := make([]string, 0)
	for i := 0; i < len(unstop); i++ {
		if _, ok := ignoreWordsDict[unstop[i]]; !ok {
			result = append(result, unstop[i])
		}
	}
	return result
}

func GetSimilar(word string) []string {
	if similar, ok := synonymsDict[word]; ok {
		if len(similar) > similarMaxLen {
			return similar[:similarMaxLen]
		} else {
			return similar
		}
	}
	return nil
}

func GetResult(segs []gse.Segment, pattern int, searchMode ...bool) []string {
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
	partOfSpeech := []string{"n", "v", "vn", "x", "an", "nz", "a", "l", "ns"}
	if pattern == 1 {
		partOfSpeech = []string{"n", "vn", "x", "an", "nz", "a", "l", "ns"}
	}
	if pattern == 2 {
		partOfSpeech = []string{"v"}
	}
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

func ToHtml(word string) string {
	result := strings.Replace(word, "\n", "<br> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;", -1)
	return result
}

func GenQueryExpression(word [][]string) string {
	result := ""
	num := len(word)
	for k := 0; k < 3; k++ {
		result += strconv.Itoa(k+1) + ".  "
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
					result += temp + " AND "
				} else {
					result += temp
				}
			}
		}
		result += "\n"
	}
	return result
}

func Score2Str(score float64) string {
	return strconv.FormatFloat(score*100, 'f', 2, 64) + "%"
}

func GenKey(segments []gse.Segment) []string {
	see := GetResult(segments, 0)
	resWords := RemoveStop(see)
	result := Unique(resWords)
	return result
}

func GenQueryAndCleanVerb(key, verb []string) []string {
	var unVerbWords []string
	for i := 0; i < len(key); i++ {
		contain := false
		for j := 0; j < len(verb); j++ {
			if key[i] == verb[j] {
				contain = true
				break
			}
		}
		if contain == false {
			unVerbWords = append(unVerbWords, key[i])
		}
	}
	length := len(unVerbWords)
	fmt.Println("query", unVerbWords)
	var query []string
	for i := length; i > 0; i-- {
		keyTemp := unVerbWords[0:i]
		query = append(query, strings.Join(keyTemp, " "))
	}
	return query
}

func GenQuery(keys []string) []string {
	var query []string
	for i := len(keys); i > 0; i-- {
		keyTemp := keys[0:i]
		query = append(query, strings.Join(keyTemp, " "))
	}
	return query
}
