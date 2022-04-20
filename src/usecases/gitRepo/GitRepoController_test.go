package gitRepo

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type gitRepoServiceMock struct {
	handleServiceFn func(n uint64) ([]string, error)
}

func (mock gitRepoServiceMock) lastProjects(n uint64) ([]string, error) {
	return mock.handleServiceFn(10)
}

func TestCreateUserControllerWithError(t *testing.T) {
	handleServiceMock := gitRepoServiceMock{}
	handleServiceMock.handleServiceFn = func(n uint64) ([]string, error) {
		return []string{}, errors.New("error")
	}

	HandleService = handleServiceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	HandleController(context)

	if response.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code 500, got %d", response.Code)
	}

	bodyRequest, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Errorf("Error reading response body: %s", err)
	}

	var responseBodyObject responseBodyDTO

	if err = json.Unmarshal(bodyRequest, &responseBodyObject); err != nil {
		t.Errorf("Error unmarshalling response body: %s", err)
	}

	expectedResponseBody := responseBodyDTO{
		Error: "error",
	}

	if responseBodyObject.Error != expectedResponseBody.Error {
		t.Errorf("Expected error message: %s, got %s", expectedResponseBody.Error, responseBodyObject.Error)
	}
}

func TestCreateUserControllerNoError(t *testing.T) {

	handleServiceMock := gitRepoServiceMock{}
	handleServiceMock.handleServiceFn = func(n uint64) ([]string, error) {
		return []string{
			"repo1",
			"repo2",
			"repo3",
		}, nil
	}

	HandleService = handleServiceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	HandleController(context)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}

	bodyRequest, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Errorf("Error reading response body: %s", err)
	}

	var responseBodyObject responseBodyDTO

	if err = json.Unmarshal(bodyRequest, &responseBodyObject); err != nil {
		t.Errorf("Error unmarshalling response body: %s", err)
	}

	expectedResponseBody := responseBodyDTO{
		Repos: []string{
			"repo1",
			"repo2",
			"repo3",
		},
	}

	if len(responseBodyObject.Repos) != len(expectedResponseBody.Repos) {
		t.Errorf("Expected repos length: %d, got %d", len(expectedResponseBody.Repos), len(responseBodyObject.Repos))
	}

	if responseBodyObject.Repos[0] != expectedResponseBody.Repos[0] {
		t.Errorf("Expected repo: %s, got %s", expectedResponseBody.Repos[0], responseBodyObject.Repos[0])
	}
}
