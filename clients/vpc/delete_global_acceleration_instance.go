package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteGlobalAccelerationInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
}

func (req *DeleteGlobalAccelerationInstanceRequest) Invoke(client *sdk.Client) (resp *DeleteGlobalAccelerationInstanceResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteGlobalAccelerationInstance", "vpc", "")
	resp = &DeleteGlobalAccelerationInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteGlobalAccelerationInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
