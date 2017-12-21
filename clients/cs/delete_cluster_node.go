package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteClusterNodeRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
	Ip        string `position:"Path" name:"Ip"`
	Force     string `position:"Query" name:"Force"`
}

func (r DeleteClusterNodeRequest) Invoke(client *sdk.Client) (response *DeleteClusterNodeResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DeleteClusterNodeRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DeleteClusterNode", "/clusters/[ClusterId]/ip/[Ip]", "", "")
	req.Method = "DELETE"

	resp := struct {
		*responses.BaseResponse
		DeleteClusterNodeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteClusterNodeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteClusterNodeResponse struct {
}
