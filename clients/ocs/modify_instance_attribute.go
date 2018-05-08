package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyInstanceAttributeRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	OcsInstanceName      string `position:"Query" name:"OcsInstanceName"`
	OldPassword          string `position:"Query" name:"OldPassword"`
	NewPassword          string `position:"Query" name:"NewPassword"`
}

func (req *ModifyInstanceAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceAttributeResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "ModifyInstanceAttribute", "", "")
	resp = &ModifyInstanceAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
