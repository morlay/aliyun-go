package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyInstanceAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	InstanceName         string `position:"Query" name:"InstanceName"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NewPassword          string `position:"Query" name:"NewPassword"`
}

func (req *ModifyInstanceAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceAttributeResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceAttribute", "redisa", "")
	resp = &ModifyInstanceAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
