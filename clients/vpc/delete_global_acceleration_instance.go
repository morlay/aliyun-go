package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteGlobalAccelerationInstanceRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
}

func (r DeleteGlobalAccelerationInstanceRequest) Invoke(client *sdk.Client) (response *DeleteGlobalAccelerationInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteGlobalAccelerationInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteGlobalAccelerationInstance", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteGlobalAccelerationInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteGlobalAccelerationInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteGlobalAccelerationInstanceResponse struct {
	RequestId string
}
