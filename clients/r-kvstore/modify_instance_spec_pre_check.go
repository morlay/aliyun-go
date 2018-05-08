package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyInstanceSpecPreCheckRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TargetInstanceClass  string `position:"Query" name:"TargetInstanceClass"`
}

func (req *ModifyInstanceSpecPreCheckRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceSpecPreCheckResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceSpecPreCheck", "redisa", "")
	resp = &ModifyInstanceSpecPreCheckResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceSpecPreCheckResponse struct {
	responses.BaseResponse
	RequestId       common.String
	IsAllowModify   bool
	DisableCommands common.String
}
