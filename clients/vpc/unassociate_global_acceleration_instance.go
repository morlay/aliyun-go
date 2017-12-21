package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnassociateGlobalAccelerationInstanceRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	InstanceType                 string `position:"Query" name:"InstanceType"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
}

func (r UnassociateGlobalAccelerationInstanceRequest) Invoke(client *sdk.Client) (response *UnassociateGlobalAccelerationInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UnassociateGlobalAccelerationInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "UnassociateGlobalAccelerationInstance", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		UnassociateGlobalAccelerationInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UnassociateGlobalAccelerationInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type UnassociateGlobalAccelerationInstanceResponse struct {
	RequestId string
}
