package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteTemplateRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TemplateId           string `position:"Query" name:"TemplateId"`
}

func (req *DeleteTemplateRequest) Invoke(client *sdk.Client) (resp *DeleteTemplateResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "DeleteTemplate", "mts", "")
	resp = &DeleteTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteTemplateResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TemplateId common.String
}
