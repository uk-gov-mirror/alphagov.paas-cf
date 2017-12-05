// This file was generated by counterfeiter
package concoursefakes

import (
	"io"
	"net/http"
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/go-concourse/concourse"
)

type FakeClient struct {
	URLStub        func() string
	uRLMutex       sync.RWMutex
	uRLArgsForCall []struct{}
	uRLReturns     struct {
		result1 string
	}
	HTTPClientStub        func() *http.Client
	hTTPClientMutex       sync.RWMutex
	hTTPClientArgsForCall []struct{}
	hTTPClientReturns     struct {
		result1 *http.Client
	}
	BuildsStub        func(concourse.Page) ([]atc.Build, concourse.Pagination, error)
	buildsMutex       sync.RWMutex
	buildsArgsForCall []struct {
		arg1 concourse.Page
	}
	buildsReturns struct {
		result1 []atc.Build
		result2 concourse.Pagination
		result3 error
	}
	BuildStub        func(buildID string) (atc.Build, bool, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		buildID string
	}
	buildReturns struct {
		result1 atc.Build
		result2 bool
		result3 error
	}
	BuildEventsStub        func(buildID string) (concourse.Events, error)
	buildEventsMutex       sync.RWMutex
	buildEventsArgsForCall []struct {
		buildID string
	}
	buildEventsReturns struct {
		result1 concourse.Events
		result2 error
	}
	BuildResourcesStub        func(buildID int) (atc.BuildInputsOutputs, bool, error)
	buildResourcesMutex       sync.RWMutex
	buildResourcesArgsForCall []struct {
		buildID int
	}
	buildResourcesReturns struct {
		result1 atc.BuildInputsOutputs
		result2 bool
		result3 error
	}
	AbortBuildStub        func(buildID string) error
	abortBuildMutex       sync.RWMutex
	abortBuildArgsForCall []struct {
		buildID string
	}
	abortBuildReturns struct {
		result1 error
	}
	CreateBuildStub        func(plan atc.Plan) (atc.Build, error)
	createBuildMutex       sync.RWMutex
	createBuildArgsForCall []struct {
		plan atc.Plan
	}
	createBuildReturns struct {
		result1 atc.Build
		result2 error
	}
	BuildPlanStub        func(buildID int) (atc.PublicBuildPlan, bool, error)
	buildPlanMutex       sync.RWMutex
	buildPlanArgsForCall []struct {
		buildID int
	}
	buildPlanReturns struct {
		result1 atc.PublicBuildPlan
		result2 bool
		result3 error
	}
	CreatePipeStub        func() (atc.Pipe, error)
	createPipeMutex       sync.RWMutex
	createPipeArgsForCall []struct{}
	createPipeReturns     struct {
		result1 atc.Pipe
		result2 error
	}
	ListContainersStub        func(queryList map[string]string) ([]atc.Container, error)
	listContainersMutex       sync.RWMutex
	listContainersArgsForCall []struct {
		queryList map[string]string
	}
	listContainersReturns struct {
		result1 []atc.Container
		result2 error
	}
	ListVolumesStub        func() ([]atc.Volume, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct{}
	listVolumesReturns     struct {
		result1 []atc.Volume
		result2 error
	}
	ListWorkersStub        func() ([]atc.Worker, error)
	listWorkersMutex       sync.RWMutex
	listWorkersArgsForCall []struct{}
	listWorkersReturns     struct {
		result1 []atc.Worker
		result2 error
	}
	GetInfoStub        func() (atc.Info, error)
	getInfoMutex       sync.RWMutex
	getInfoArgsForCall []struct{}
	getInfoReturns     struct {
		result1 atc.Info
		result2 error
	}
	GetCLIReaderStub        func(arch, platform string) (io.ReadCloser, http.Header, error)
	getCLIReaderMutex       sync.RWMutex
	getCLIReaderArgsForCall []struct {
		arch     string
		platform string
	}
	getCLIReaderReturns struct {
		result1 io.ReadCloser
		result2 http.Header
		result3 error
	}
	ListPipelinesStub        func() ([]atc.Pipeline, error)
	listPipelinesMutex       sync.RWMutex
	listPipelinesArgsForCall []struct{}
	listPipelinesReturns     struct {
		result1 []atc.Pipeline
		result2 error
	}
	TeamStub        func(teamName string) concourse.Team
	teamMutex       sync.RWMutex
	teamArgsForCall []struct {
		teamName string
	}
	teamReturns struct {
		result1 concourse.Team
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) URL() string {
	fake.uRLMutex.Lock()
	fake.uRLArgsForCall = append(fake.uRLArgsForCall, struct{}{})
	fake.recordInvocation("URL", []interface{}{})
	fake.uRLMutex.Unlock()
	if fake.URLStub != nil {
		return fake.URLStub()
	} else {
		return fake.uRLReturns.result1
	}
}

func (fake *FakeClient) URLCallCount() int {
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	return len(fake.uRLArgsForCall)
}

func (fake *FakeClient) URLReturns(result1 string) {
	fake.URLStub = nil
	fake.uRLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeClient) HTTPClient() *http.Client {
	fake.hTTPClientMutex.Lock()
	fake.hTTPClientArgsForCall = append(fake.hTTPClientArgsForCall, struct{}{})
	fake.recordInvocation("HTTPClient", []interface{}{})
	fake.hTTPClientMutex.Unlock()
	if fake.HTTPClientStub != nil {
		return fake.HTTPClientStub()
	} else {
		return fake.hTTPClientReturns.result1
	}
}

func (fake *FakeClient) HTTPClientCallCount() int {
	fake.hTTPClientMutex.RLock()
	defer fake.hTTPClientMutex.RUnlock()
	return len(fake.hTTPClientArgsForCall)
}

func (fake *FakeClient) HTTPClientReturns(result1 *http.Client) {
	fake.HTTPClientStub = nil
	fake.hTTPClientReturns = struct {
		result1 *http.Client
	}{result1}
}

func (fake *FakeClient) Builds(arg1 concourse.Page) ([]atc.Build, concourse.Pagination, error) {
	fake.buildsMutex.Lock()
	fake.buildsArgsForCall = append(fake.buildsArgsForCall, struct {
		arg1 concourse.Page
	}{arg1})
	fake.recordInvocation("Builds", []interface{}{arg1})
	fake.buildsMutex.Unlock()
	if fake.BuildsStub != nil {
		return fake.BuildsStub(arg1)
	} else {
		return fake.buildsReturns.result1, fake.buildsReturns.result2, fake.buildsReturns.result3
	}
}

func (fake *FakeClient) BuildsCallCount() int {
	fake.buildsMutex.RLock()
	defer fake.buildsMutex.RUnlock()
	return len(fake.buildsArgsForCall)
}

func (fake *FakeClient) BuildsArgsForCall(i int) concourse.Page {
	fake.buildsMutex.RLock()
	defer fake.buildsMutex.RUnlock()
	return fake.buildsArgsForCall[i].arg1
}

func (fake *FakeClient) BuildsReturns(result1 []atc.Build, result2 concourse.Pagination, result3 error) {
	fake.BuildsStub = nil
	fake.buildsReturns = struct {
		result1 []atc.Build
		result2 concourse.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) Build(buildID string) (atc.Build, bool, error) {
	fake.buildMutex.Lock()
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		buildID string
	}{buildID})
	fake.recordInvocation("Build", []interface{}{buildID})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(buildID)
	} else {
		return fake.buildReturns.result1, fake.buildReturns.result2, fake.buildReturns.result3
	}
}

func (fake *FakeClient) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeClient) BuildArgsForCall(i int) string {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.buildArgsForCall[i].buildID
}

func (fake *FakeClient) BuildReturns(result1 atc.Build, result2 bool, result3 error) {
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 atc.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) BuildEvents(buildID string) (concourse.Events, error) {
	fake.buildEventsMutex.Lock()
	fake.buildEventsArgsForCall = append(fake.buildEventsArgsForCall, struct {
		buildID string
	}{buildID})
	fake.recordInvocation("BuildEvents", []interface{}{buildID})
	fake.buildEventsMutex.Unlock()
	if fake.BuildEventsStub != nil {
		return fake.BuildEventsStub(buildID)
	} else {
		return fake.buildEventsReturns.result1, fake.buildEventsReturns.result2
	}
}

func (fake *FakeClient) BuildEventsCallCount() int {
	fake.buildEventsMutex.RLock()
	defer fake.buildEventsMutex.RUnlock()
	return len(fake.buildEventsArgsForCall)
}

func (fake *FakeClient) BuildEventsArgsForCall(i int) string {
	fake.buildEventsMutex.RLock()
	defer fake.buildEventsMutex.RUnlock()
	return fake.buildEventsArgsForCall[i].buildID
}

func (fake *FakeClient) BuildEventsReturns(result1 concourse.Events, result2 error) {
	fake.BuildEventsStub = nil
	fake.buildEventsReturns = struct {
		result1 concourse.Events
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) BuildResources(buildID int) (atc.BuildInputsOutputs, bool, error) {
	fake.buildResourcesMutex.Lock()
	fake.buildResourcesArgsForCall = append(fake.buildResourcesArgsForCall, struct {
		buildID int
	}{buildID})
	fake.recordInvocation("BuildResources", []interface{}{buildID})
	fake.buildResourcesMutex.Unlock()
	if fake.BuildResourcesStub != nil {
		return fake.BuildResourcesStub(buildID)
	} else {
		return fake.buildResourcesReturns.result1, fake.buildResourcesReturns.result2, fake.buildResourcesReturns.result3
	}
}

func (fake *FakeClient) BuildResourcesCallCount() int {
	fake.buildResourcesMutex.RLock()
	defer fake.buildResourcesMutex.RUnlock()
	return len(fake.buildResourcesArgsForCall)
}

func (fake *FakeClient) BuildResourcesArgsForCall(i int) int {
	fake.buildResourcesMutex.RLock()
	defer fake.buildResourcesMutex.RUnlock()
	return fake.buildResourcesArgsForCall[i].buildID
}

func (fake *FakeClient) BuildResourcesReturns(result1 atc.BuildInputsOutputs, result2 bool, result3 error) {
	fake.BuildResourcesStub = nil
	fake.buildResourcesReturns = struct {
		result1 atc.BuildInputsOutputs
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) AbortBuild(buildID string) error {
	fake.abortBuildMutex.Lock()
	fake.abortBuildArgsForCall = append(fake.abortBuildArgsForCall, struct {
		buildID string
	}{buildID})
	fake.recordInvocation("AbortBuild", []interface{}{buildID})
	fake.abortBuildMutex.Unlock()
	if fake.AbortBuildStub != nil {
		return fake.AbortBuildStub(buildID)
	} else {
		return fake.abortBuildReturns.result1
	}
}

func (fake *FakeClient) AbortBuildCallCount() int {
	fake.abortBuildMutex.RLock()
	defer fake.abortBuildMutex.RUnlock()
	return len(fake.abortBuildArgsForCall)
}

func (fake *FakeClient) AbortBuildArgsForCall(i int) string {
	fake.abortBuildMutex.RLock()
	defer fake.abortBuildMutex.RUnlock()
	return fake.abortBuildArgsForCall[i].buildID
}

func (fake *FakeClient) AbortBuildReturns(result1 error) {
	fake.AbortBuildStub = nil
	fake.abortBuildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateBuild(plan atc.Plan) (atc.Build, error) {
	fake.createBuildMutex.Lock()
	fake.createBuildArgsForCall = append(fake.createBuildArgsForCall, struct {
		plan atc.Plan
	}{plan})
	fake.recordInvocation("CreateBuild", []interface{}{plan})
	fake.createBuildMutex.Unlock()
	if fake.CreateBuildStub != nil {
		return fake.CreateBuildStub(plan)
	} else {
		return fake.createBuildReturns.result1, fake.createBuildReturns.result2
	}
}

func (fake *FakeClient) CreateBuildCallCount() int {
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	return len(fake.createBuildArgsForCall)
}

func (fake *FakeClient) CreateBuildArgsForCall(i int) atc.Plan {
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	return fake.createBuildArgsForCall[i].plan
}

func (fake *FakeClient) CreateBuildReturns(result1 atc.Build, result2 error) {
	fake.CreateBuildStub = nil
	fake.createBuildReturns = struct {
		result1 atc.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) BuildPlan(buildID int) (atc.PublicBuildPlan, bool, error) {
	fake.buildPlanMutex.Lock()
	fake.buildPlanArgsForCall = append(fake.buildPlanArgsForCall, struct {
		buildID int
	}{buildID})
	fake.recordInvocation("BuildPlan", []interface{}{buildID})
	fake.buildPlanMutex.Unlock()
	if fake.BuildPlanStub != nil {
		return fake.BuildPlanStub(buildID)
	} else {
		return fake.buildPlanReturns.result1, fake.buildPlanReturns.result2, fake.buildPlanReturns.result3
	}
}

func (fake *FakeClient) BuildPlanCallCount() int {
	fake.buildPlanMutex.RLock()
	defer fake.buildPlanMutex.RUnlock()
	return len(fake.buildPlanArgsForCall)
}

func (fake *FakeClient) BuildPlanArgsForCall(i int) int {
	fake.buildPlanMutex.RLock()
	defer fake.buildPlanMutex.RUnlock()
	return fake.buildPlanArgsForCall[i].buildID
}

func (fake *FakeClient) BuildPlanReturns(result1 atc.PublicBuildPlan, result2 bool, result3 error) {
	fake.BuildPlanStub = nil
	fake.buildPlanReturns = struct {
		result1 atc.PublicBuildPlan
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) CreatePipe() (atc.Pipe, error) {
	fake.createPipeMutex.Lock()
	fake.createPipeArgsForCall = append(fake.createPipeArgsForCall, struct{}{})
	fake.recordInvocation("CreatePipe", []interface{}{})
	fake.createPipeMutex.Unlock()
	if fake.CreatePipeStub != nil {
		return fake.CreatePipeStub()
	} else {
		return fake.createPipeReturns.result1, fake.createPipeReturns.result2
	}
}

func (fake *FakeClient) CreatePipeCallCount() int {
	fake.createPipeMutex.RLock()
	defer fake.createPipeMutex.RUnlock()
	return len(fake.createPipeArgsForCall)
}

func (fake *FakeClient) CreatePipeReturns(result1 atc.Pipe, result2 error) {
	fake.CreatePipeStub = nil
	fake.createPipeReturns = struct {
		result1 atc.Pipe
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListContainers(queryList map[string]string) ([]atc.Container, error) {
	fake.listContainersMutex.Lock()
	fake.listContainersArgsForCall = append(fake.listContainersArgsForCall, struct {
		queryList map[string]string
	}{queryList})
	fake.recordInvocation("ListContainers", []interface{}{queryList})
	fake.listContainersMutex.Unlock()
	if fake.ListContainersStub != nil {
		return fake.ListContainersStub(queryList)
	} else {
		return fake.listContainersReturns.result1, fake.listContainersReturns.result2
	}
}

func (fake *FakeClient) ListContainersCallCount() int {
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	return len(fake.listContainersArgsForCall)
}

func (fake *FakeClient) ListContainersArgsForCall(i int) map[string]string {
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	return fake.listContainersArgsForCall[i].queryList
}

func (fake *FakeClient) ListContainersReturns(result1 []atc.Container, result2 error) {
	fake.ListContainersStub = nil
	fake.listContainersReturns = struct {
		result1 []atc.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListVolumes() ([]atc.Volume, error) {
	fake.listVolumesMutex.Lock()
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct{}{})
	fake.recordInvocation("ListVolumes", []interface{}{})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub()
	} else {
		return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
	}
}

func (fake *FakeClient) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeClient) ListVolumesReturns(result1 []atc.Volume, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 []atc.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListWorkers() ([]atc.Worker, error) {
	fake.listWorkersMutex.Lock()
	fake.listWorkersArgsForCall = append(fake.listWorkersArgsForCall, struct{}{})
	fake.recordInvocation("ListWorkers", []interface{}{})
	fake.listWorkersMutex.Unlock()
	if fake.ListWorkersStub != nil {
		return fake.ListWorkersStub()
	} else {
		return fake.listWorkersReturns.result1, fake.listWorkersReturns.result2
	}
}

func (fake *FakeClient) ListWorkersCallCount() int {
	fake.listWorkersMutex.RLock()
	defer fake.listWorkersMutex.RUnlock()
	return len(fake.listWorkersArgsForCall)
}

func (fake *FakeClient) ListWorkersReturns(result1 []atc.Worker, result2 error) {
	fake.ListWorkersStub = nil
	fake.listWorkersReturns = struct {
		result1 []atc.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetInfo() (atc.Info, error) {
	fake.getInfoMutex.Lock()
	fake.getInfoArgsForCall = append(fake.getInfoArgsForCall, struct{}{})
	fake.recordInvocation("GetInfo", []interface{}{})
	fake.getInfoMutex.Unlock()
	if fake.GetInfoStub != nil {
		return fake.GetInfoStub()
	} else {
		return fake.getInfoReturns.result1, fake.getInfoReturns.result2
	}
}

func (fake *FakeClient) GetInfoCallCount() int {
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	return len(fake.getInfoArgsForCall)
}

func (fake *FakeClient) GetInfoReturns(result1 atc.Info, result2 error) {
	fake.GetInfoStub = nil
	fake.getInfoReturns = struct {
		result1 atc.Info
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetCLIReader(arch string, platform string) (io.ReadCloser, http.Header, error) {
	fake.getCLIReaderMutex.Lock()
	fake.getCLIReaderArgsForCall = append(fake.getCLIReaderArgsForCall, struct {
		arch     string
		platform string
	}{arch, platform})
	fake.recordInvocation("GetCLIReader", []interface{}{arch, platform})
	fake.getCLIReaderMutex.Unlock()
	if fake.GetCLIReaderStub != nil {
		return fake.GetCLIReaderStub(arch, platform)
	} else {
		return fake.getCLIReaderReturns.result1, fake.getCLIReaderReturns.result2, fake.getCLIReaderReturns.result3
	}
}

func (fake *FakeClient) GetCLIReaderCallCount() int {
	fake.getCLIReaderMutex.RLock()
	defer fake.getCLIReaderMutex.RUnlock()
	return len(fake.getCLIReaderArgsForCall)
}

func (fake *FakeClient) GetCLIReaderArgsForCall(i int) (string, string) {
	fake.getCLIReaderMutex.RLock()
	defer fake.getCLIReaderMutex.RUnlock()
	return fake.getCLIReaderArgsForCall[i].arch, fake.getCLIReaderArgsForCall[i].platform
}

func (fake *FakeClient) GetCLIReaderReturns(result1 io.ReadCloser, result2 http.Header, result3 error) {
	fake.GetCLIReaderStub = nil
	fake.getCLIReaderReturns = struct {
		result1 io.ReadCloser
		result2 http.Header
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListPipelines() ([]atc.Pipeline, error) {
	fake.listPipelinesMutex.Lock()
	fake.listPipelinesArgsForCall = append(fake.listPipelinesArgsForCall, struct{}{})
	fake.recordInvocation("ListPipelines", []interface{}{})
	fake.listPipelinesMutex.Unlock()
	if fake.ListPipelinesStub != nil {
		return fake.ListPipelinesStub()
	} else {
		return fake.listPipelinesReturns.result1, fake.listPipelinesReturns.result2
	}
}

func (fake *FakeClient) ListPipelinesCallCount() int {
	fake.listPipelinesMutex.RLock()
	defer fake.listPipelinesMutex.RUnlock()
	return len(fake.listPipelinesArgsForCall)
}

func (fake *FakeClient) ListPipelinesReturns(result1 []atc.Pipeline, result2 error) {
	fake.ListPipelinesStub = nil
	fake.listPipelinesReturns = struct {
		result1 []atc.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Team(teamName string) concourse.Team {
	fake.teamMutex.Lock()
	fake.teamArgsForCall = append(fake.teamArgsForCall, struct {
		teamName string
	}{teamName})
	fake.recordInvocation("Team", []interface{}{teamName})
	fake.teamMutex.Unlock()
	if fake.TeamStub != nil {
		return fake.TeamStub(teamName)
	} else {
		return fake.teamReturns.result1
	}
}

func (fake *FakeClient) TeamCallCount() int {
	fake.teamMutex.RLock()
	defer fake.teamMutex.RUnlock()
	return len(fake.teamArgsForCall)
}

func (fake *FakeClient) TeamArgsForCall(i int) string {
	fake.teamMutex.RLock()
	defer fake.teamMutex.RUnlock()
	return fake.teamArgsForCall[i].teamName
}

func (fake *FakeClient) TeamReturns(result1 concourse.Team) {
	fake.TeamStub = nil
	fake.teamReturns = struct {
		result1 concourse.Team
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	fake.hTTPClientMutex.RLock()
	defer fake.hTTPClientMutex.RUnlock()
	fake.buildsMutex.RLock()
	defer fake.buildsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.buildEventsMutex.RLock()
	defer fake.buildEventsMutex.RUnlock()
	fake.buildResourcesMutex.RLock()
	defer fake.buildResourcesMutex.RUnlock()
	fake.abortBuildMutex.RLock()
	defer fake.abortBuildMutex.RUnlock()
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	fake.buildPlanMutex.RLock()
	defer fake.buildPlanMutex.RUnlock()
	fake.createPipeMutex.RLock()
	defer fake.createPipeMutex.RUnlock()
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	fake.listWorkersMutex.RLock()
	defer fake.listWorkersMutex.RUnlock()
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	fake.getCLIReaderMutex.RLock()
	defer fake.getCLIReaderMutex.RUnlock()
	fake.listPipelinesMutex.RLock()
	defer fake.listPipelinesMutex.RUnlock()
	fake.teamMutex.RLock()
	defer fake.teamMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ concourse.Client = new(FakeClient)