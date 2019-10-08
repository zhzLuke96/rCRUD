package rCRUD

import (
	"strconv"
	"strings"
)

type Param map[string]string

func (p Param) Pass(check map[string]interface{}) bool {
	for pk, pv := range p {
		keyFound := false
		for k, v := range check {
			if pk == k {
				if !queryEqule(pv, v) {
					return false
				} else {
					continue
				}
			}
			if len(pk)-len(k) == 1 && strings.HasPrefix(pk, k) {
				keyFound = true
				if !conditionPass(pk[len(pk)-1:], pv, v) {
					return false
				}
			}
		}
		if !keyFound {
			return false
		}
	}
	return true
}

func conditionPass(method string, pv string, v interface{}) bool {
	pvnum, pverr := strconv.ParseFloat(pv, 64)

	vstr, vstrok := v.(string)
	vnum, verr := toFloat(v)
	if !vstrok && verr != nil {
		return false
	}
	switch method {
	case "<":
		if pverr != nil || verr != nil {
			return false
		}
		return vnum <= pvnum
	case ">":
		if pverr != nil || verr != nil {
			return false
		}
		return vnum >= pvnum
	case "!":
		if pverr != nil || verr != nil {
			return vstr != pv
		}
		return vnum != pvnum
	case "^":
		return strings.HasPrefix(vstr, pv)
	case "$":
		return strings.HasSuffix(vstr, pv)
	case "*":
		return strings.Index(vstr, pv) != -1
	case "@", "%", ":", ";", "|", ".", ",":
		return false
	default:
		return false
	}
}
