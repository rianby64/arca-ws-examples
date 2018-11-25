package arca

func dummy(requestParams *interface{}, context *interface{}) (interface{}, error) {
	return nil, nil
}

// RegisterSource whatever
func RegisterSource(name string, methods DIRUD) {
	handlers[name] = map[string]requestHandler{
		"subscribe": dummy,
		"describe":  dummy,
		"insert":    dummy,
		"read":      dummy,
		"update":    dummy,
		"delete":    dummy,
	}

	handler := handlers[name]
	if methods.Insert != nil {
		handler["insert"] = methods.Insert
	}
	if methods.Read != nil {
		handler["read"] = methods.Read
	}
	if methods.Update != nil {
		handler["update"] = methods.Update
	}

	if methods.Delete != nil {
		handler["delete"] = methods.Delete
	}
}
