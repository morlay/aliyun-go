package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateNodeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	NodeType             string `position:"Query" name:"NodeType"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	NodeStorage          int    `position:"Query" name:"NodeStorage"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeClass            string `position:"Query" name:"NodeClass"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
}

func (req *CreateNodeRequest) Invoke(client *sdk.Client) (resp *CreateNodeResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "CreateNode", "dds", "")
	resp = &CreateNodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateNodeResponse struct {
	responses.BaseResponse
	RequestId common.String
	OrderId   common.String
}
