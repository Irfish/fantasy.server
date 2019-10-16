package pb

var (
	servicePrefix = "/fantacy/server/"
	webPrefix     = "/fantacy/web/"
	gwPrefix      = "/fantacy/gw/"
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
	return "fantacy." + SERVICE_name[int32(id)]
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
