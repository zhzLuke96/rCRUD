package rCRUD

import (
	"net/http"

	"github.com/zhzluke96/Rango/rango"
)

const (

	// CRUD
	readRequestKey   = "QUERY"
	createRequestKey = "INSERT"
	updateRequestKey = "UPDATE"
	deleteRequestKey = "DELETE"
)

var curdResponser = rango.NewResponser("curd")

type crudify interface {
	// pager
	// Mate(cname string) (mate map[string]interface{},err error)
	Mate(string) (map[string]interface{}, error)

	// Create
	// Insert(cname string, e map[string]interface{}) (idx int, err error)
	Insert(string, map[string]interface{}) (int, error)

	// Updata
	// Update(canme string, idx int, e map[string]interface{}) error
	Update(string, int, map[string]interface{}) error

	// Retrieve
	// Query(cname string, idx int) (e map[string]interface{}, err error)
	Query(string, int) (map[string]interface{}, error)
	// QueryMap(cname string, param map[string]interface{}, page, n int) (es []map[string]interface{}, err error)
	QueryMap(string, *Param, int, int) ([]map[string]interface{}, error)

	// Delete
	// Remove(cname string, idx int) error
	Remove(string, int) error
	// RemoveMap(cname string, param map[string]interface{}, n int) error
	RemoveMap(string, *Param, int) error

	// Auth
	AuthHandler(*http.Request, string, string) (string, error)
}

func curdErrResp(code int, msg string, err error) *rango.ErrResponse {
	resp := curdResponser.NewErrResponse()
	return resp.Set(code, msg, err)
}

type rCRUD struct {
	crudEntry crudify
	Sev       *rango.Server
}

func (c rCRUD) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Sev.Handler.ServeHTTP(w, r)
}

func New(c crudify, preFix string) *rCRUD {
	sev := rango.NewSev("__crud__")
	// query list
	sev.GET("/{entity:\\w+}s/", queryLs(&c))
	// query one index
	sev.GET("/{entity:\\w+}/{idx:\\d+}/", queryOneIdx(&c))
	// query one map
	sev.GET("/{entity:\\w+}/", queryMap(&c))
	// insert
	sev.POST("/{entity:\\w+}s/", insert(&c))
	// updata
	sev.Func("/{entity:\\w+}/{idx:\\d+}/", updata(&c)).Methods("PUT")
	// delete
	sev.Func("/{entity:\\w+}/{idx:\\d+}/", delete(&c)).Methods("DELETE")
	// delete map
	sev.Func("/{entity:\\w+}s/", deleteMap(&c)).Methods("DELETE")

	sev.Handler.Use(rango.StripPrefixMid(preFix))
	sev.Handler.Use(sev.Router.Mid)
	return &rCRUD{
		crudEntry: c,
		Sev:       sev,
	}
}
