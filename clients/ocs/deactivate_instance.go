package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeactivateInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
}

func (r DeactivateInstanceRequest) Invoke(client *sdk.Client) (response *DeactivateInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeactivateInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "DeactivateInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		DeactivateInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeactivateInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeactivateInstanceResponse struct {
	RequestId string
}
