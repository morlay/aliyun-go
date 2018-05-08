package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateWaterMarkTemplateRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	WaterMarkTemplateId  string `position:"Query" name:"WaterMarkTemplateId"`
	Config               string `position:"Query" name:"Config"`
}

func (req *UpdateWaterMarkTemplateRequest) Invoke(client *sdk.Client) (resp *UpdateWaterMarkTemplateResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateWaterMarkTemplate", "mts", "")
	resp = &UpdateWaterMarkTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateWaterMarkTemplateResponse struct {
	responses.BaseResponse
	RequestId         common.String
	WaterMarkTemplate UpdateWaterMarkTemplateWaterMarkTemplate
}

type UpdateWaterMarkTemplateWaterMarkTemplate struct {
	Id         common.String
	Name       common.String
	Width      common.String
	Height     common.String
	Dx         common.String
	Dy         common.String
	ReferPos   common.String
	Type       common.String
	State      common.String
	Timeline   UpdateWaterMarkTemplateTimeline
	RatioRefer UpdateWaterMarkTemplateRatioRefer
}

type UpdateWaterMarkTemplateTimeline struct {
	Start    common.String
	Duration common.String
}

type UpdateWaterMarkTemplateRatioRefer struct {
	Dx     common.String
	Dy     common.String
	Width  common.String
	Height common.String
}
