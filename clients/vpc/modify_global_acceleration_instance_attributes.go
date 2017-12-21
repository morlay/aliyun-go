package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyGlobalAccelerationInstanceAttributesRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	Name                         string `position:"Query" name:"Name"`
	Description                  string `position:"Query" name:"Description"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
}

func (r ModifyGlobalAccelerationInstanceAttributesRequest) Invoke(client *sdk.Client) (response *ModifyGlobalAccelerationInstanceAttributesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyGlobalAccelerationInstanceAttributesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyGlobalAccelerationInstanceAttributes", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyGlobalAccelerationInstanceAttributesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyGlobalAccelerationInstanceAttributesResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyGlobalAccelerationInstanceAttributesResponse struct {
	RequestId string
}
