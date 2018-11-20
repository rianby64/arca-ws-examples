package main

// JSONRPCerror is the structure of JSON-RPC response
type JSONRPCerror struct {
	Code    int
	Message string
	Data    interface{}
}

// JSONRPCrequest is the structure of JSON-RPC request
type JSONRPCrequest struct {
	Jsonrpc string
	ID      string
	Method  string
	Context interface{}
	Params  interface{}
}

// JSONRPCresponse is the structure of JSON-RPC response
type JSONRPCresponse struct {
	Jsonrpc string
	ID      string
	Method  string
	Context interface{}
	Result  interface{}
	Error   interface{}
}
