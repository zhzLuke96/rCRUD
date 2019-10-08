# rCRUD
Rapid Prototyping on rango.

# Usage
```golang
type xxx struct{
    // ...
}
func (x xxx)AuthHandler(*http.Request, string, string) (string, error)
func (x xxx)Insert(string, map[string]interface{}) (int, error)
func (x xxx)Update(string, int, map[string]interface{}) error
func (x xxx)Mate(string) (map[string]interface{}, error)
func (x xxx)Query(string, int) (map[string]interface{}, error)
func (x xxx)QueryMap(string, *Param, int, int) ([]map[string]interface{}, error)
func (x xxx)Remove(string, int) error
func (x xxx)RemoveMap(string, *Param, int) error

// xxx.(curdify)

c := rCURD.New(new(xxx))

_,sev := sev.Group("/api")
sev.Router = c.Sev.Router
sev.Handler = c.Sev.Handler
```

# crudify
```golang
type crudify interface {
	AuthHandler(*http.Request, string, string) (string, error)
    
    Insert(string, map[string]interface{}) (int, error)
    
    Update(string, int, map[string]interface{}) error
    
    Mate(string) (map[string]interface{}, error)
	Query(string, int) (map[string]interface{}, error)
    QueryMap(string, *Param, int, int) ([]map[string]interface{}, error)
    
	Remove(string, int) error
	RemoveMap(string, *Param, int) error
}
```

# Param Query
| eg          | desc            |
| ----------- | --------------- |
| code<=100   | 小于            |
| code>=100   | 大于            |
| text!=admin | text不等于admin |
| code!=100   | code不等于100   |
| text^=admin | text开头为admin |
| text$=admin | text结尾为admin |
| text*=admin | text中包含admin |

# Idempotence
> restful base

# LICENSE
GPL-3.0
