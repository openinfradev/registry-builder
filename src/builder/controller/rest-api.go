package controller

import (
	"builder/service"
	"builder/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// var sampleService *service.SampleService
// var httpsampleService *service.HTTPSampleService
var registryService *service.RegistryService
var dockerService *service.DockerService

func init() {
	// inject service
	injectServices()

	// health
	addRequestMapping(
		RequestMapper{
			Method:  "GET",
			Path:    "/health",
			Request: health,
		},
	)

	// registry catalog
	addRequestMapping(
		RequestMapper{
			Method:  "GET",
			Path:    "/registry/catalog",
			Request: getRegistryCatalog,
		},
	)

	// registry repositories tag list
	addRequestMapping(
		RequestMapper{
			Method:  "GET",
			Path:    "/registry/repositories/*name",
			Request: getRegistryRepositories,
		},
	)

	// docker build
	// needs POST (GET is test)
	addRequestMapping(
		RequestMapper{
			Method:  "GET",
			Path:    "/docker/build",
			Request: buildDockerFile,
		},
	)

	// docker tag
	// needs PATCH (GET is test)
	addRequestMapping(
		RequestMapper{
			Method:  "GET",
			Path:    "/docker/tag",
			Request: tagDockerImage,
		},
	)
}

func injectServices() {
	// sampleService = new(service.SampleService)
	// httpsampleService = new(service.HTTPSampleService)
	dockerService = new(service.DockerService)
	registryService = new(service.RegistryService)
}

/*
	Request Mapping Functions
*/

// health
// @Summary health check api
// @Description builder의 health를 체크할 목적의 api
// @name health
// @Produce  json
// @Router /health [get]
// @Success 200
func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// getRegistryCatalog
// @Summary docker registry catalog api
// @Description docker registry catalog api
// @name getRegistryCatalog
// @Produce  json
// @Router /registry/catalog [get]
// @Success 200
func getRegistryCatalog(c *gin.Context) {
	r := registryService.GetCatalog()

	c.JSON(http.StatusOK, r)
}

// getRegistryRepositories
// @Summary docker registry repositories api
// @Description docker registry repositories api
// @name getRegistryRepositories
// @Param name path string false "Repository Name"
// @Produce  json
// @Router /registry/repositories/{name} [get]
// @Success 200
func getRegistryRepositories(c *gin.Context) {
	repoName := c.Params.ByName("name")
	repoName = strings.Replace(repoName, "/", "", 1)

	if repoName == "" {
		r := registryService.GetRepositories()
		c.JSON(http.StatusOK, r)
	} else {
		r := registryService.GetRepository(repoName)
		if r.Tags == nil {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.JSON(http.StatusOK, r)
	}
}

// buildDockerFile
// @Summary docker build by dockerfile
// @Description docker build by dockerfile api
// @name buildDockerFile
// @Produce  json
// @Router /docker/build [get]
// @Success 200
func buildDockerFile(c *gin.Context) {
	// test arguments
	repoName := "exntu/sample1"
	dockerfilePath := "./sample"

	r := dockerService.Build(repoName, dockerfilePath)
	c.JSON(http.StatusOK, util.StringToMap(r))
}

// tagDockerImage
// @Summary docker image tag
// @Description docker image tag
// @name tagDockerImage
// @Produce  json
// @Router /docker/tag [get]
// @Success 200
func tagDockerImage(c *gin.Context) {
	// test arguments
	repoName := "exntu/sample1"
	oldTag := "latest"
	newTag := "v1"

	r := dockerService.Tag(repoName, oldTag, newTag)
	c.JSON(http.StatusOK, util.StringToMap(r))
}
