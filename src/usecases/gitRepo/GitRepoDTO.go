package gitRepo

type ResponseBodyDTO struct {
	LastProjects ServiceResponseDTO `json:"last_projects, omitempty"`
	Error        string             `json:"error,omitempty"`
}

type ServiceResponseDTO struct {
	ProjectsNames   string `json:"projects_names,omitempty"`
	TotalForksCount int    `json:"total_forks_count,omitempty"`
}
