package rCRUD

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func anyLess(a, b interface{}) bool {
	switch a.(type) {
	case nil:
		return true
	case string:
		return strings.Compare(a.(string), b.(string)) == -1
	case int:
		return a.(int) < b.(int)
	case float64:
		return a.(float64) < b.(float64)
	case bool:
		return a.(bool)
	default:
		return true
	}
}

func excluding(in map[string]string, ex ...string) map[string]string {
	if len(ex) == 0 {
		return in
	}
	ret := make(map[string]string)
	for k, v := range in {
		isPassKey := false
		for _, v := range ex {
			if k == v {
				isPassKey = true
				break
			}
		}
		if isPassKey {
			continue
		}
		ret[k] = v
	}
	return ret
}

func queryEqule(a, b interface{}) bool {
	switch at := a.(type) {
	case nil:
		return b == nil
	case bool:
		if bt, ok := b.(bool); ok {
			return at == bt
		}
		return false
	case string, int, float64, float32, uint, complex64, complex128:
		return fmt.Sprint(a) == fmt.Sprint(b)
	default:
		return false
	}
}

func toFloat(v interface{}) (float64, error) {
	switch vt := v.(type) {
	case string:
		return strconv.ParseFloat(vt, 64)
	case int:
		return float64(vt), nil
	case float64:
		return vt, nil
	case float32:
		return float64(vt), nil
	default:
		return 0, fmt.Errorf("cant parsing")
	}
}

type DictArr struct {
	arr     []map[string]interface{}
	lessKey string
}

func NewDictArr(a []map[string]interface{}) *DictArr {
	return &DictArr{
		arr:     a,
		lessKey: "idx",
	}
}

func (a *DictArr) Len() int {
	return len(a.arr)
}

func (a *DictArr) Less(i, j int) bool {
	if a.lessKey == "" {
		return true
	}
	iV := a.arr[i][a.lessKey]
	jV := a.arr[j][a.lessKey]
	return anyLess(iV, jV)
}

func (a *DictArr) Swap(i, j int) {
	a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
}

func (a *DictArr) Sort(key string) *DictArr {
	a.lessKey = key
	sort.Sort(a)
	return a
}
