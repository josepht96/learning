package registry

//Registration struct
type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

//ServiceName string
type ServiceName string

const (
	LogService = ServiceName("LogService")
)
