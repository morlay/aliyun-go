package cds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetBuildsRequest struct {
	Start         int    `position:"Query" name:"Start"`
	NumberPerPage int    `position:"Query" name:"NumberPerPage"`
	JobName       string `position:"Path" name:"JobName"`
}

func (r GetBuildsRequest) Invoke(client *sdk.Client) (response *GetBuildsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetBuildsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Cds", "2017-09-25", "GetBuilds", "/v1/job/[JobName]/builds", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetBuildsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetBuildsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetBuildsResponse struct {
	RequestId string
	Builds    GetBuildsBuildList
}

type GetBuildsBuild struct {
	BuildNumber int
	Duration    int
	StartTime   int64
	Log         string
	BuildEnv    string
}

type GetBuildsBuildList []GetBuildsBuild

func (list *GetBuildsBuildList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetBuildsBuild)
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
