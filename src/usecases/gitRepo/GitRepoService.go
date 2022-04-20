package gitRepo

import (
	"context"
	"github.com/hasura/go-graphql-client"
	"github.com/lucchesisp/go-test-egroupas/src/config"
	"golang.org/x/oauth2"
)

type gitRepoServiceInterface interface {
	lastProjects(n int) (ServiceResponseDTO, error)
}

type gitRepoServiceImplementation struct{}

var HandleService gitRepoServiceInterface = gitRepoServiceImplementation{}

func (service gitRepoServiceImplementation) lastProjects(n int) (ServiceResponseDTO, error) {

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GetEnvVariable("GITLAB_TOKEN")},
	)

	httpClient := oauth2.NewClient(context.Background(), src)

	client := graphql.NewClient("https://gitlab.com/api/graphql", httpClient)

	var last_projects struct {
		Projects struct {
			Nodes []struct {
				Name        string
				Description string
				ForksCount  int
			}
		} `graphql:"projects(last: $n)"`
	}

	variables := map[string]interface{}{
		"n": graphql.Int(n),
	}

	err := client.Query(context.Background(), &last_projects, variables)

	if err != nil {
		return ServiceResponseDTO{}, err
	}

	var projects_names string

	for i, project := range last_projects.Projects.Nodes {
		var delimiter string

		if i < len(last_projects.Projects.Nodes)-1 {
			delimiter = ", "
		}

		projects_names += project.Name + delimiter
	}

	var totalForksCount int

	for _, project := range last_projects.Projects.Nodes {
		totalForksCount += project.ForksCount
	}

	var response ServiceResponseDTO
	response.ProjectsNames = projects_names
	response.TotalForksCount = totalForksCount

	return response, nil
}
