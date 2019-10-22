package pb

var (
	servicePrefix = "/fantasy/server/"
	webPrefix     = "/fantasy/web/"
	gwPrefix      = "/fantasy/gw/"
)

func ServicePrefix() string {
	return servicePrefix
}

func WebPrefix() string {
	return webPrefix
}

func GwPrefix() string {
	return gwPrefix
}

func GetName(id SERVICE) string {
	return "fantasy." + SERVICE_name[int32(id)]
}

func GetServerKey(id SERVICE) string {
	return servicePrefix + SERVICE_name[int32(id)]
}

func GetWebKey(id SERVICE) string {
	return webPrefix + SERVICE_name[int32(id)]
}

func GetGwKey(id SERVICE) string {
	return gwPrefix + SERVICE_name[int32(id)]
}
