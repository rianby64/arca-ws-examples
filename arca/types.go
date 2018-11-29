package arca

// JSONRPCBase is the base for both request and response structures
type JSONRPCBase struct {
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

// DIRUD whatever
type DIRUD struct {
	Describe requestHandler
	Insert   requestHandler
	Read     requestHandler
	Update   requestHandler
	Delete   requestHandler
}

type requestHandler func(requestParams *interface{},
	context *interface{}, response chan interface{}) error
type requestHandlers map[string]map[string]*requestHandler
