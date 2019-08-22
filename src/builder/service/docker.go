package service

import (
	"bufio"
	"builder/constant"
	tacoconst "builder/constant/taco"
	"builder/model"
	"builder/repository"
	"builder/util/logger"
	tacoutil "builder/util/taco"
	"encoding/base64"
	"fmt"
	"os/exec"
)

// DockerService is docker command relative service
type DockerService struct{}

var fileManager *FileManager
var registryRepository *repository.RegistryRepository

func init() {
	fileManager = new(FileManager)
	registryRepository = new(repository.RegistryRepository)
}

// BuildByDockerfile is docker building by dockerfile
func (d *DockerService) BuildByDockerfile(params *model.DockerBuildByFileParam) *model.BasicResult {
	// needs using goroutine
	// and saving log line by line

	// phase - preparing
	p := tacoutil.MakePhaseLog(params.BuildID, tacoconst.PhasePreparing.StartSeq, tacoconst.PhasePreparing.Status)
	registryRepository.InsertBuildLog(p)

	// decoding contents
	decoded, err := base64.StdEncoding.DecodeString(params.Contents)
	if err != nil {
		return &model.BasicResult{
			Code:    constant.ResultFail,
			Message: "contents isn't base64 encoded",
		}
	}

	// phase - unpacking
	p = tacoutil.MakePhaseLog(params.BuildID, tacoconst.PhaseUnpacking.StartSeq, tacoconst.PhaseUnpacking.Status)
	registryRepository.InsertBuildLog(p)

	path, err := fileManager.WriteDockerfile(string(decoded))
	if err != nil {
		return &model.BasicResult{
			Code:    constant.ResultFail,
			Message: "",
		}
	}

	return d.Build(params.BuildID, params.Name, path, params.UseCache, params.Push)
}

// BuildByGitRepository is docker building by git repository
func (d *DockerService) BuildByGitRepository(params *model.DockerBuildByGitParam) *model.BasicResult {
	// needs using goroutine
	// and saving log line by line

	// phase - preparing
	p := tacoutil.MakePhaseLog(params.BuildID, tacoconst.PhasePreparing.StartSeq, tacoconst.PhasePreparing.Status)
	registryRepository.InsertBuildLog(p)

	// decoding userPW
	decoded, err := base64.StdEncoding.DecodeString(params.UserPW)
	if err != nil {
		return &model.BasicResult{
			Code:    constant.ResultFail,
			Message: "userPW isn't base64 encoded",
		}
	}

	// phase - unpacking
	p = tacoutil.MakePhaseLog(params.BuildID, tacoconst.PhaseUnpacking.StartSeq, tacoconst.PhaseUnpacking.Status)
	registryRepository.InsertBuildLog(p)

	// not using go-routine (not yet)
	// ch := make(chan string, 1)	// dirPath
	path, err := fileManager.PullGitRepository(params.GitRepository, params.UserID, string(decoded))
	if err != nil {
		return &model.BasicResult{
			Code:    constant.ResultFail,
			Message: "",
		}
	}

	return d.Build(params.BuildID, params.Name, path, params.UseCache, params.Push)
}

// Build is docker building by file path
func (d *DockerService) Build(buildID string, repoName string, dockerfilePath string, useCache bool, push bool) *model.BasicResult {

	// async
	ch := make(chan string)
	if push {
		go d.BuildAndPush(ch, buildID, repoName, dockerfilePath, useCache)
	} else {
		go buildJob(ch, buildID, repoName, dockerfilePath, useCache)
	}

	// only ok
	return &model.BasicResult{
		Code:    constant.ResultSuccess,
		Message: "",
	}
}

// BuildAndPush is docker build and push
func (d *DockerService) BuildAndPush(ch chan<- string, buildID string, repoName string, dockerfilePath string, useCache bool) {

	proc := make(chan string)
	// build
	go buildJob(proc, buildID, repoName, dockerfilePath, useCache)
	r := <-proc
	if r == constant.ResultFail {
		procBuildError(buildID)
		ch <- constant.ResultFail
	}

	// tag
	go tagJob(proc, repoName, "latest", "latest")
	r = <-proc
	if r == constant.ResultFail {
		procBuildError(buildID)
		ch <- constant.ResultFail
	}

	// push
	// phase - push
	registryRepository.UpdateBuildPhase(buildID, tacoconst.PhasePushing.Status)
	p := tacoutil.MakePhaseLog(buildID, tacoconst.PhasePushing.StartSeq, tacoconst.PhasePushing.Status)
	registryRepository.InsertBuildLog(p)

	go pushJob(proc, repoName, "latest")
	r = <-proc
	if r == constant.ResultFail {
		procBuildError(buildID)
		ch <- constant.ResultFail
	}

	// phase - complete
	procBuildComplete(buildID)
	ch <- constant.ResultSuccess
}

// Tag is image tagging
func (d *DockerService) Tag(params *model.DockerTagParam) *model.BasicResult {

	// needs using goroutine
	// and saving log line by line

	// sync
	ch := make(chan string)
	go tagJob(ch, params.Name, params.OldTag, params.NewTag)
	r := <-ch

	return &model.BasicResult{
		Code:    r,
		Message: "",
	}
}

// Push is docker image pushing
func (d *DockerService) Push(params *model.DockerPushParam) *model.BasicResult {
	// needs using goroutine
	// and saving log line by line

	// async
	ch := make(chan string)
	go pushJob(ch, params.Name, params.Tag)

	// only ok
	return &model.BasicResult{
		Code:    constant.ResultSuccess,
		Message: "",
	}
}

func pushJob(ch chan<- string, repoName string, tag string) {
	logger.DEBUG("service/docker.go", "pushJob", fmt.Sprintf("pushJob start [%s:%s]", repoName, tag))

	repoName = basicinfo.RegistryEndpoint + "/" + repoName + ":" + tag
	push := exec.Command("docker", "push", repoName)

	r := ""
	stdout, _ := push.StdoutPipe()
	stderr, _ := push.StderrPipe()
	push.Start()
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		r += m + "\n"
		logger.DEBUG("service/docker.go", "pushJob", m)
	}
	errscan := bufio.NewScanner(stderr)
	errscan.Split(bufio.ScanLines)
	for errscan.Scan() {
		m := errscan.Text()
		logger.ERROR("service/docker.go", "buildJob", m)

		ch <- constant.ResultFail
	}
	push.Wait()

	ch <- constant.ResultSuccess

	logger.DEBUG("service/docker.go", "pushJob", fmt.Sprintf("pushJob end [%s]", repoName))
}

func tagJob(ch chan<- string, repoName string, oldTag string, newTag string) {
	logger.DEBUG("service/docker.go", "tagJob", fmt.Sprintf("tagJob [%s] [%s] to [%s]", repoName, oldTag, newTag))

	oldRepo := repoName + ":" + oldTag
	newRepo := basicinfo.RegistryEndpoint + "/" + repoName + ":" + newTag

	tag := exec.Command("docker", "tag", oldRepo, newRepo)

	err := tag.Run()
	if err != nil {
		logger.ERROR("service/docker.go", "tagJob", "tagJob is failed")
		ch <- constant.ResultFail
	} else {
		logger.DEBUG("service/docker.go", "tagJob", "tagJob is success")
		ch <- constant.ResultSuccess
	}
}

func buildJob(ch chan<- string, buildID string, repoName string, dockerfilePath string, useCache bool) {
	logger.DEBUG("service/docker.go", "buildJob", "buildJob start "+repoName)

	seq := tacoconst.PhaseBuilding.StartSeq

	// phase - build
	// updating build phase
	registryRepository.UpdateBuildPhase(buildID, tacoconst.PhaseBuilding.Status)
	p := tacoutil.MakePhaseLog(buildID, seq, tacoconst.PhaseBuilding.Status)
	registryRepository.InsertBuildLog(p)

	repoName = repoName + ":latest"
	var build *exec.Cmd
	if useCache {
		// phase - checking cache
		p = tacoutil.MakePhaseLog(buildID, seq, tacoconst.PhaseCheckingCache.Status)
		registryRepository.InsertBuildLog(p)

		build = exec.Command("docker", "build", "--network=host", "-t", repoName, dockerfilePath)
	} else {
		build = exec.Command("docker", "build", "--no-cache", "--network=host", "-t", repoName, dockerfilePath)
	}

	stdout, _ := build.StdoutPipe()
	stderr, _ := build.StderrPipe()
	build.Start()
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		seq++
		m := scanner.Text()
		row := tacoutil.ParseLog(buildID, seq, m)
		registryRepository.InsertBuildLog(row)

		logger.DEBUG("service/docker.go", "buildJob", m)
	}
	errscan := bufio.NewScanner(stderr)
	errscan.Split(bufio.ScanLines)
	for errscan.Scan() {
		seq++
		m := errscan.Text()
		errrow := tacoutil.ParseLog(buildID, seq, m)
		errrow.Type = "error"
		registryRepository.InsertBuildLog(errrow)
		logger.ERROR("service/docker.go", "buildJob", m)
		// path removeall - because of breaking
		fileManager.DeleteDirectory(dockerfilePath)

		ch <- constant.ResultFail
	}

	build.Wait()

	// path removeall
	fileManager.DeleteDirectory(dockerfilePath)

	ch <- constant.ResultSuccess

	logger.DEBUG("service/docker.go", "buildJob", "buildJob end "+repoName)
}

func garbageCollectJob(ch chan<- string) {
	logger.DEBUG("service/docker.go", "garbageCollectJob", "garbage collect start")

	gc := exec.Command("docker", "exec", basicinfo.RegistryName, "bin/registry", "garbage-collect", "/etc/docker/registry/config.yml")

	r := ""
	stdout, _ := gc.StdoutPipe()
	gc.Start()
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		r += m + "\n"
		logger.DEBUG("service/docker.go", "garbageCollectJob", m)
	}
	gc.Wait()

	logger.DEBUG("service/docker.go", "garbageCollectJob", "garbage collect end")

	ch <- r
}

func procBuildComplete(buildID string) {
	registryRepository.UpdateBuildPhase(buildID, tacoconst.PhaseComplete.Status)
	p := tacoutil.MakePhaseLog(buildID, tacoconst.PhaseComplete.StartSeq, tacoconst.PhaseComplete.Status)
	registryRepository.InsertBuildLog(p)
}

func procBuildError(buildID string) {
	registryRepository.UpdateBuildPhase(buildID, tacoconst.PhaseError.Status)
	p := tacoutil.MakePhaseLog(buildID, tacoconst.PhaseError.StartSeq, tacoconst.PhaseError.Status)
	registryRepository.InsertBuildLog(p)
}
