package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteWaterMarkTemplateRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	WaterMarkTemplateId  string `position:"Query" name:"WaterMarkTemplateId"`
}

func (req *DeleteWaterMarkTemplateRequest) Invoke(client *sdk.Client) (resp *DeleteWaterMarkTemplateResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "DeleteWaterMarkTemplate", "", "")
	resp = &DeleteWaterMarkTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteWaterMarkTemplateResponse struct {
	responses.BaseResponse
	RequestId           string
	WaterMarkTemplateId string
}
