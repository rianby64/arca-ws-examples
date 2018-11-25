package arca

// JSONRPCBase is the base for both request and response structures
type JSONRPCBase struct {
	Jsonrpc string
	ID      string
	Method  string
	Context interface{}
}

// JSONRPCerror is the structure of JSON-RPC response
type JSONRPCerror struct {
	Code    int
	Message string
	Data    interface{}
}

// JSONRPCrequest is the structure of JSON-RPC request
type JSONRPCrequest struct {
	JSONRPCBase
	Params interface{}
}

// JSONRPCresponse is the structure of JSON-RPC response
type JSONRPCresponse struct {
	JSONRPCBase
	Result interface{}
	Error  interface{}
}

// IRUD whatever
type IRUD struct {
	Insert func(requestParams *interface{}, context *interface{}) (interface{}, error)
	Read   func(requestParams *interface{}, context *interface{}) (interface{}, error)
	Update func(requestParams *interface{}, context *interface{}) (interface{}, error)
	Delete func(requestParams *interface{}, context *interface{}) (interface{}, error)
}

type sIRUDd struct {
	IRUD
	Subscribe func(requestParams *interface{}, context *interface{}) (interface{}, error)
	Describe  func(requestParams *interface{}, context *interface{}) (interface{}, error)
}

type requestHandler func(requestParams *interface{},
	context *interface{}) (interface{}, error)
type requestHandlers map[string]map[string]requestHandler
