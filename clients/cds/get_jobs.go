package cds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetJobsRequest struct {
	Start         int `position:"Query" name:"Start"`
	NumberPerPage int `position:"Query" name:"NumberPerPage"`
}

func (r GetJobsRequest) Invoke(client *sdk.Client) (response *GetJobsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetJobsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Cds", "2017-09-25", "GetJobs", "/v1/jobs", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetJobsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetJobsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetJobsResponse struct {
	RequestId string
	Jobs      GetJobsJobList
}

type GetJobsJob struct {
	JobName             string
	SuccessRate         float32
	TotalBuilds         int
	LastFailedBuild     GetJobsLastFailedBuild
	LastSuccessfulBuild GetJobsLastSuccessfulBuild
}

type GetJobsLastFailedBuild struct {
	BuildNumber string
	Duration    int64
	StartTime   int64
	Log         string
	BuildEnv    string
}

type GetJobsLastSuccessfulBuild struct {
	BuildNumber int
	Duration    int
	StartTime   int64
	Log         string
	BuildEnv    string
}

type GetJobsJobList []GetJobsJob

func (list *GetJobsJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobsJob)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
