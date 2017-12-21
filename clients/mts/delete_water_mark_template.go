package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteWaterMarkTemplateRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	WaterMarkTemplateId  string `position:"Query" name:"WaterMarkTemplateId"`
}

func (r DeleteWaterMarkTemplateRequest) Invoke(client *sdk.Client) (response *DeleteWaterMarkTemplateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteWaterMarkTemplateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "DeleteWaterMarkTemplate", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteWaterMarkTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteWaterMarkTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteWaterMarkTemplateResponse struct {
	RequestId           string
	WaterMarkTemplateId string
}
