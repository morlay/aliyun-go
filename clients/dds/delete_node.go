package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNodeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeId               string `position:"Query" name:"NodeId"`
}

func (r DeleteNodeRequest) Invoke(client *sdk.Client) (response *DeleteNodeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteNodeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "DeleteNode", "dds", "")

	resp := struct {
		*responses.BaseResponse
		DeleteNodeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteNodeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteNodeResponse struct {
	RequestId string
	TaskId    int
}
