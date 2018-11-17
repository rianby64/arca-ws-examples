package main

// JSONRPCerror is the structure of JSON-RPC response
type JSONRPCerror struct {
	Code    int
	Message string
	Data    map[string]interface{}
}

// JSONRPCrequest is the structure of JSON-RPC request
type JSONRPCrequest struct {
	Jsonrpc string
	Method  string
	Params  map[string]interface{}
	ID      string
}

// JSONRPCresponse is the structure of JSON-RPC response
type JSONRPCresponse struct {
	Jsonrpc string
	Result  map[string]interface{}
	Error   JSONRPCerror
	ID      string
}
