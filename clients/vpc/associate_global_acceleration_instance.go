package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AssociateGlobalAccelerationInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	BackendServerId              string `position:"Query" name:"BackendServerId"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
	BackendServerRegionId        string `position:"Query" name:"BackendServerRegionId"`
	BackendServerType            string `position:"Query" name:"BackendServerType"`
}

func (req *AssociateGlobalAccelerationInstanceRequest) Invoke(client *sdk.Client) (resp *AssociateGlobalAccelerationInstanceResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "AssociateGlobalAccelerationInstance", "vpc", "")
	resp = &AssociateGlobalAccelerationInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type AssociateGlobalAccelerationInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
