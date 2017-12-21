package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRecycleBinRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceIds          string `position:"Query" name:"ResourceIds"`
}

func (req *DeleteRecycleBinRequest) Invoke(client *sdk.Client) (resp *DeleteRecycleBinResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteRecycleBin", "ecs", "")
	resp = &DeleteRecycleBinResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteRecycleBinResponse struct {
	responses.BaseResponse
	RequestId string
}
