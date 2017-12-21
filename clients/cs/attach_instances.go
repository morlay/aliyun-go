package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachInstancesRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (r AttachInstancesRequest) Invoke(client *sdk.Client) (response *AttachInstancesResponse, err error) {
	req := struct {
		*requests.RoaRequest
		AttachInstancesRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "AttachInstances", "/clusters/[ClusterId]/attach", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		AttachInstancesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AttachInstancesResponse

	err = client.DoAction(&req, &resp)
	return
}

type AttachInstancesResponse struct {
}
