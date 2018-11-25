package arca

// RegisterSource whatever
func RegisterSource(name string, methods IRUD) {
	handlers[name] = map[string]requestHandler{
		"insert": methods.Insert,
		"read":   methods.Read,
		"update": methods.Update,
		"delete": methods.Delete,
	}
}
