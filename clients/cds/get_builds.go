package cds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetBuildsRequest struct {
	requests.RoaRequest
	Start         int    `position:"Query" name:"Start"`
	NumberPerPage int    `position:"Query" name:"NumberPerPage"`
	JobName       string `position:"Path" name:"JobName"`
}

func (req *GetBuildsRequest) Invoke(client *sdk.Client) (resp *GetBuildsResponse, err error) {
	req.InitWithApiInfo("Cds", "2017-09-25", "GetBuilds", "/v1/job/[JobName]/builds", "", "")
	req.Method = "GET"

	resp = &GetBuildsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetBuildsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Builds    GetBuildsBuildList
}

type GetBuildsBuild struct {
	BuildNumber common.Integer
	Duration    common.Integer
	StartTime   common.Long
	Log         common.String
	BuildEnv    common.String
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
