package main

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
