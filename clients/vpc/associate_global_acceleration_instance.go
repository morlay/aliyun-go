package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AssociateGlobalAccelerationInstanceRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	BackendServerId              string `position:"Query" name:"BackendServerId"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
	BackendServerRegionId        string `position:"Query" name:"BackendServerRegionId"`
	BackendServerType            string `position:"Query" name:"BackendServerType"`
}

func (r AssociateGlobalAccelerationInstanceRequest) Invoke(client *sdk.Client) (response *AssociateGlobalAccelerationInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AssociateGlobalAccelerationInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "AssociateGlobalAccelerationInstance", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		AssociateGlobalAccelerationInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AssociateGlobalAccelerationInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type AssociateGlobalAccelerationInstanceResponse struct {
	RequestId string
}
