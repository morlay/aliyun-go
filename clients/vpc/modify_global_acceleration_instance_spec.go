package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyGlobalAccelerationInstanceSpecRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                    string `position:"Query" name:"Bandwidth"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
}

func (r ModifyGlobalAccelerationInstanceSpecRequest) Invoke(client *sdk.Client) (response *ModifyGlobalAccelerationInstanceSpecResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyGlobalAccelerationInstanceSpecRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyGlobalAccelerationInstanceSpec", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyGlobalAccelerationInstanceSpecResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyGlobalAccelerationInstanceSpecResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyGlobalAccelerationInstanceSpecResponse struct {
	RequestId string
}
