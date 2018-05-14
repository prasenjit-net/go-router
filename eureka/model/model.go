package model

// EurekaResponse represetn a eureka response structure
type EurekaResponse struct {
	Application ApplicationInfo
}

// ApplicationInfo represents a application data in eureka response
type ApplicationInfo struct {
	Name     string
	Instance []AppInstance
}

// AppInstance represents a instance in eureka response
type AppInstance struct {
	InstanceID string
	HostName   string
	IPAddr     string
	Status     string
	Port       Port
	SecurePort Port
}

// Port holds port information
type Port struct {
	Value   int    `json:"$"`
	Enabled string `json:"@enabled"`
}
