package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnassociateEipAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r UnassociateEipAddressRequest) Invoke(client *sdk.Client) (response *UnassociateEipAddressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UnassociateEipAddressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "UnassociateEipAddress", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		UnassociateEipAddressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UnassociateEipAddressResponse

	err = client.DoAction(&req, &resp)
	return
}

type UnassociateEipAddressResponse struct {
	RequestId string
}
