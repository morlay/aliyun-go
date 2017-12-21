package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceMaintainTimeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	MaintainStartTime    string `position:"Query" name:"MaintainStartTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MaintainEndTime      string `position:"Query" name:"MaintainEndTime"`
}

func (req *ModifyInstanceMaintainTimeRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceMaintainTimeResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceMaintainTime", "redisa", "")
	resp = &ModifyInstanceMaintainTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceMaintainTimeResponse struct {
	responses.BaseResponse
	RequestId string
}
