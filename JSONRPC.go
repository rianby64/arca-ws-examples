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
	Method  string
	Params  interface{}
	ID      string
}

// JSONRPCresponse is the structure of JSON-RPC response
type JSONRPCresponse struct {
	Jsonrpc string
	Result  interface{}
	Error   interface{}
	ID      string
}
