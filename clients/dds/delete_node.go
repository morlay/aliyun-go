package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteNodeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeId               string `position:"Query" name:"NodeId"`
}

func (req *DeleteNodeRequest) Invoke(client *sdk.Client) (resp *DeleteNodeResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "DeleteNode", "dds", "")
	resp = &DeleteNodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteNodeResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskId    common.Integer
}
