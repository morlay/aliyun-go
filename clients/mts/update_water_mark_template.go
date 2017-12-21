package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateWaterMarkTemplateRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	WaterMarkTemplateId  string `position:"Query" name:"WaterMarkTemplateId"`
	Config               string `position:"Query" name:"Config"`
}

func (r UpdateWaterMarkTemplateRequest) Invoke(client *sdk.Client) (response *UpdateWaterMarkTemplateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateWaterMarkTemplateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateWaterMarkTemplate", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateWaterMarkTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateWaterMarkTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateWaterMarkTemplateResponse struct {
	RequestId         string
	WaterMarkTemplate UpdateWaterMarkTemplateWaterMarkTemplate
}

type UpdateWaterMarkTemplateWaterMarkTemplate struct {
	Id         string
	Name       string
	Width      string
	Height     string
	Dx         string
	Dy         string
	ReferPos   string
	Type       string
	State      string
	Timeline   UpdateWaterMarkTemplateTimeline
	RatioRefer UpdateWaterMarkTemplateRatioRefer
}

type UpdateWaterMarkTemplateTimeline struct {
	Start    string
	Duration string
}

type UpdateWaterMarkTemplateRatioRefer struct {
	Dx     string
	Dy     string
	Width  string
	Height string
}
