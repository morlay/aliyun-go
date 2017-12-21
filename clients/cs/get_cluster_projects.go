package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetClusterProjectsRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (r GetClusterProjectsRequest) Invoke(client *sdk.Client) (response *GetClusterProjectsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetClusterProjectsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "GetClusterProjects", "/clusters/[ClusterId]/projects", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetClusterProjectsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetClusterProjectsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetClusterProjectsResponse struct {
}
