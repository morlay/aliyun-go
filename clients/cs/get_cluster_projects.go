package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetClusterProjectsRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *GetClusterProjectsRequest) Invoke(client *sdk.Client) (resp *GetClusterProjectsResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "GetClusterProjects", "/clusters/[ClusterId]/projects", "", "")
	req.Method = "GET"

	resp = &GetClusterProjectsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetClusterProjectsResponse struct {
	responses.BaseResponse
}
