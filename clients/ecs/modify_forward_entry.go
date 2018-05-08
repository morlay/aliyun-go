package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyForwardEntryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string `position:"Query" name:"IpProtocol"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InternalIp           string `position:"Query" name:"InternalIp"`
	ForwardEntryId       string `position:"Query" name:"ForwardEntryId"`
	InternalPort         string `position:"Query" name:"InternalPort"`
	ExternalIp           string `position:"Query" name:"ExternalIp"`
	ExternalPort         string `position:"Query" name:"ExternalPort"`
}

func (req *ModifyForwardEntryRequest) Invoke(client *sdk.Client) (resp *ModifyForwardEntryResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyForwardEntry", "ecs", "")
	resp = &ModifyForwardEntryResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyForwardEntryResponse struct {
	responses.BaseResponse
	RequestId common.String
}
