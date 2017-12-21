package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateNodeRequest struct {
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

func (r CreateNodeRequest) Invoke(client *sdk.Client) (response *CreateNodeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateNodeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "CreateNode", "dds", "")

	resp := struct {
		*responses.BaseResponse
		CreateNodeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateNodeResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateNodeResponse struct {
	RequestId string
	OrderId   string
}
