package bss

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetResourceBusinessStatusRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	BusinessStatus       string `position:"Query" name:"BusinessStatus"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

func (req *SetResourceBusinessStatusRequest) Invoke(client *sdk.Client) (resp *SetResourceBusinessStatusResponse, err error) {
	req.InitWithApiInfo("Bss", "2014-07-14", "SetResourceBusinessStatus", "", "")
	resp = &SetResourceBusinessStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetResourceBusinessStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
}
