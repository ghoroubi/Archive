package main

var Messages map[string]string

func InitMessages() {
	Messages = make(map[string]string)

	Messages["ConfigNotFound"] = "Config not found!"
	//Messages["ElasticError"] = "Error in elasticserach!"

	Messages["ActiveUserSms"] = "کد فعال سازی شما: %s "
}
