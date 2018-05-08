package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyInstanceConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Config               string `position:"Query" name:"Config"`
}

func (req *ModifyInstanceConfigRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceConfigResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceConfig", "redisa", "")
	resp = &ModifyInstanceConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
