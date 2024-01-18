package config

type MetaData struct {
	//Name of the config
	Name string
}

type RequestConfig struct {
	Name string
	//Url of the request
	Url string
	//Headers of the request
	Headers map[string]string
	//Method of the request
	Method string
}

type ExecutionConfig struct {
	//Interval in seconds
	Interval int
}

type Config struct {
	MetaData        MetaData
	Requests        []RequestConfig
	ExecutionConfig ExecutionConfig
}
