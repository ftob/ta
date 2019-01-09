package health


type Service interface {
	Health() Status
}


type service struct {
	nodes   []string
}


type Notes []string


type CPU struct {
	
}

type Memory struct {

}

type Uptime struct {
	
}

type Details struct {
	CPUUtilization CPU `json:"cpu:utilization"`
	MemoryUtilization Memory `json:"memory:utilization"`
	Uptime Uptime `json:"uptime"`
}

// https://tools.ietf.org/id/draft-inadarei-api-health-check-01.html#rfc.section.3
type Status struct {
	Status string `json:"status"`
	Version string `json:"version"`
	ReleaseID string `json:"releaseID"`
	Notes Notes `json:"notes"`
	Output string `json:"output"`
	Details string  `json:"details"`
	Links map[string]string `json:"links"`
	ServiceID string `json:"serviceID"`
	Description string `json:"description"`
}
