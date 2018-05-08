package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddWaterMarkTemplateRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Config               string `position:"Query" name:"Config"`
}

func (req *AddWaterMarkTemplateRequest) Invoke(client *sdk.Client) (resp *AddWaterMarkTemplateResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddWaterMarkTemplate", "mts", "")
	resp = &AddWaterMarkTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddWaterMarkTemplateResponse struct {
	responses.BaseResponse
	RequestId         common.String
	WaterMarkTemplate AddWaterMarkTemplateWaterMarkTemplate
}

type AddWaterMarkTemplateWaterMarkTemplate struct {
	Id         common.String
	Name       common.String
	Width      common.String
	Height     common.String
	Dx         common.String
	Dy         common.String
	ReferPos   common.String
	Type       common.String
	State      common.String
	Timeline   AddWaterMarkTemplateTimeline
	RatioRefer AddWaterMarkTemplateRatioRefer
}

type AddWaterMarkTemplateTimeline struct {
	Start    common.String
	Duration common.String
}

type AddWaterMarkTemplateRatioRefer struct {
	Dx     common.String
	Dy     common.String
	Width  common.String
	Height common.String
}
