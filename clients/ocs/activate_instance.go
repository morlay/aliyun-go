package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ActivateInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
}

func (r ActivateInstanceRequest) Invoke(client *sdk.Client) (response *ActivateInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ActivateInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "ActivateInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		ActivateInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ActivateInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type ActivateInstanceResponse struct {
	RequestId string
}
