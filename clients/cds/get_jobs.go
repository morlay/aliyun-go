package cds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetJobsRequest struct {
	requests.RoaRequest
	Start         int `position:"Query" name:"Start"`
	NumberPerPage int `position:"Query" name:"NumberPerPage"`
}

func (req *GetJobsRequest) Invoke(client *sdk.Client) (resp *GetJobsResponse, err error) {
	req.InitWithApiInfo("Cds", "2017-09-25", "GetJobs", "/v1/jobs", "", "")
	req.Method = "GET"

	resp = &GetJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetJobsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Jobs      GetJobsJobList
}

type GetJobsJob struct {
	JobName             common.String
	SuccessRate         common.Float
	TotalBuilds         common.Integer
	LastFailedBuild     GetJobsLastFailedBuild
	LastSuccessfulBuild GetJobsLastSuccessfulBuild
}

type GetJobsLastFailedBuild struct {
	BuildNumber common.String
	Duration    common.Long
	StartTime   common.Long
	Log         common.String
	BuildEnv    common.String
}

type GetJobsLastSuccessfulBuild struct {
	BuildNumber common.Integer
	Duration    common.Integer
	StartTime   common.Long
	Log         common.String
	BuildEnv    common.String
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
