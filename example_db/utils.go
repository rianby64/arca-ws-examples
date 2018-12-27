package example

import (
	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindArcaWithGrid whatever
func BindArcaWithGrid(
	connStr string,
	s *arca.JSONRPCExtensionWS,
	g *grid.Grid,
	methods *grid.QUID,
	source string,
) {

	var queryMethod arca.JSONRequestHandler = g.Query
	var updateMethod arca.JSONRequestHandler = g.Update
	var insertMethod arca.JSONRequestHandler = g.Insert
	var deleteMethod arca.JSONRequestHandler = g.Delete

	s.RegisterMethod(source, "read", &queryMethod)
	s.RegisterMethod(source, "update", &updateMethod)
	s.RegisterMethod(source, "insert", &insertMethod)
	s.RegisterMethod(source, "delete", &deleteMethod)

	g.Register(methods)
}
