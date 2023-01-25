package registry

//Registration struct
type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
	Host        string
	Port        string
}

//ServiceName string
type ServiceName string

const (
	//LogService const
	LogService = ServiceName("LogService")
	//GeneratorService const
	GeneratorService = ServiceName("GeneratorService")
)
