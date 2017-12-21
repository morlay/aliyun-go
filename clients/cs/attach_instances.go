package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachInstancesRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *AttachInstancesRequest) Invoke(client *sdk.Client) (resp *AttachInstancesResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "AttachInstances", "/clusters/[ClusterId]/attach", "", "")
	req.Method = "POST"

	resp = &AttachInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type AttachInstancesResponse struct {
	responses.BaseResponse
}
