package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceCapacityRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	Capacity             int64  `position:"Query" name:"Capacity"`
}

func (r ModifyInstanceCapacityRequest) Invoke(client *sdk.Client) (response *ModifyInstanceCapacityResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceCapacityRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "ModifyInstanceCapacity", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceCapacityResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyInstanceCapacityResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceCapacityResponse struct {
	RequestId string
}
