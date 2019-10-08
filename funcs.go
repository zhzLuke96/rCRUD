package rCRUD

import (
	"strconv"

	"github.com/zhzluke96/Rango/rango"
)

func queryLs(c *crudify) func(vars *rango.ReqVars) interface{} {
	return func(vars *rango.ReqVars) interface{} {
		query := Param(excluding(vars.Query(), "page", "n"))
		cname, err := vars.Get("entity")
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		if msg, err := (*c).AuthHandler(vars.Request(), readRequestKey, cname.(string)); err != nil {
			return curdErrResp(401, msg, err)
		}
		count := vars.GetDefault("n", "-1").(string)
		countNum, err := strconv.Atoi(count)
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		page := vars.GetDefault("page", "0").(string)
		pageNum, err := strconv.Atoi(page)
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		res, err := (*c).QueryMap(cname.(string), &query, pageNum, countNum)
		if err != nil {
			return curdErrResp(404, "cant find entity.", err)
		}
		if res == nil {
			return []string{}
		}
		return res
	}
}

func queryOneIdx(c *crudify) func(vars *rango.ReqVars) interface{} {
	return func(vars *rango.ReqVars) interface{} {
		cname, err := vars.Get("entity")
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		if msg, err := (*c).AuthHandler(vars.Request(), readRequestKey, cname.(string)); err != nil {
			return curdErrResp(401, msg, err)
		}
		idx, err := vars.Get("idx")
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		idxNum, err := strconv.Atoi(idx.(string))
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		res, err := (*c).Query(cname.(string), idxNum)
		if err != nil {
			return curdErrResp(404, "cant find entity.", err)
		}
		return res
	}
}

func queryMap(c *crudify) func(vars *rango.ReqVars) interface{} {
	return func(vars *rango.ReqVars) interface{} {
		query := Param(vars.Query())
		cname, err := vars.Get("entity")
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		if msg, err := (*c).AuthHandler(vars.Request(), readRequestKey, cname.(string)); err != nil {
			return curdErrResp(401, msg, err)
		}
		if len(query) == 0 {
			return nil
		}
		res, err := (*c).QueryMap(cname.(string), &query, 0, 1)
		if err != nil {
			return curdErrResp(404, "cant find entity.", err)
		}
		if len(res) == 0 {
			return nil
		}
		return res[0]
	}
}
func insert(c *crudify) func(vars *rango.ReqVars) interface{} {
	return func(vars *rango.ReqVars) interface{} {
		cname, err := vars.Get("entity")
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		if msg, err := (*c).AuthHandler(vars.Request(), createRequestKey, cname.(string)); err != nil {
			return curdErrResp(401, msg, err)
		}
		idx, err := (*c).Insert(cname.(string), vars.JSON())
		if err != nil {
			return curdErrResp(404, "cant insert entity.", err)
		}
		return map[string]interface{}{
			"idx": idx,
		}
	}
}
func updata(c *crudify) func(vars *rango.ReqVars) interface{} {
	return func(vars *rango.ReqVars) interface{} {
		cname, err := vars.Get("entity")
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		if msg, err := (*c).AuthHandler(vars.Request(), updateRequestKey, cname.(string)); err != nil {
			return curdErrResp(401, msg, err)
		}
		idx, err := vars.Get("idx")
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		idxNum, err := strconv.Atoi(idx.(string))
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		err = (*c).Update(cname.(string), idxNum, vars.JSON())
		if err != nil {
			return curdErrResp(404, "cant updata entity.", err)
		}
		return "success"
	}
}
func delete(c *crudify) func(vars *rango.ReqVars) interface{} {
	return func(vars *rango.ReqVars) interface{} {
		cname, err := vars.Get("entity")
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		if msg, err := (*c).AuthHandler(vars.Request(), deleteRequestKey, cname.(string)); err != nil {
			return curdErrResp(401, msg, err)
		}
		idx, err := vars.Get("idx")
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		idxNum, err := strconv.Atoi(idx.(string))
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		err = (*c).Remove(cname.(string), idxNum)
		if err != nil {
			return curdErrResp(404, "cant updata entity.", err)
		}
		return "success"
	}
}
func deleteMap(c *crudify) func(vars *rango.ReqVars) interface{} {
	return func(vars *rango.ReqVars) interface{} {
		query := Param(excluding(vars.Query(), "n"))
		cname, err := vars.Get("entity")
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		if msg, err := (*c).AuthHandler(vars.Request(), deleteRequestKey, cname.(string)); err != nil {
			return curdErrResp(401, msg, err)
		}
		count := vars.GetDefault("n", "-1").(string)
		countNum, err := strconv.Atoi(count)
		if err != nil {
			return curdErrResp(500, "server is wrong.", err)
		}
		err = (*c).RemoveMap(cname.(string), &query, countNum)
		if err != nil {
			return curdErrResp(404, "cant remove entity.", err)
		}
		return "success"
	}
}
