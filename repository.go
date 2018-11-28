package core

func message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}
