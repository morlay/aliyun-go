package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReleasePublicIpAddressRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PublicIpAddress      string `position:"Query" name:"PublicIpAddress"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReleasePublicIpAddressRequest) Invoke(client *sdk.Client) (resp *ReleasePublicIpAddressResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ReleasePublicIpAddress", "ecs", "")
	resp = &ReleasePublicIpAddressResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleasePublicIpAddressResponse struct {
	responses.BaseResponse
	RequestId common.String
}
