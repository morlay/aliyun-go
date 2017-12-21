package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyNodeSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	NodeStorage          int    `position:"Query" name:"NodeStorage"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeClass            string `position:"Query" name:"NodeClass"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	NodeId               string `position:"Query" name:"NodeId"`
}

func (req *ModifyNodeSpecRequest) Invoke(client *sdk.Client) (resp *ModifyNodeSpecResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyNodeSpec", "dds", "")
	resp = &ModifyNodeSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyNodeSpecResponse struct {
	responses.BaseResponse
	RequestId string
	OrderId   string
}
