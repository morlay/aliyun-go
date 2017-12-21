package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AssociateEipAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r AssociateEipAddressRequest) Invoke(client *sdk.Client) (response *AssociateEipAddressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AssociateEipAddressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "AssociateEipAddress", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		AssociateEipAddressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AssociateEipAddressResponse

	err = client.DoAction(&req, &resp)
	return
}

type AssociateEipAddressResponse struct {
	RequestId string
}
