package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyImageAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ImageName            string `position:"Query" name:"ImageName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyImageAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyImageAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyImageAttribute", "ecs", "")
	resp = &ModifyImageAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyImageAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
