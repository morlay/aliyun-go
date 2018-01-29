package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId         string
	WaterMarkTemplate AddWaterMarkTemplateWaterMarkTemplate
}

type AddWaterMarkTemplateWaterMarkTemplate struct {
	Id         string
	Name       string
	Width      string
	Height     string
	Dx         string
	Dy         string
	ReferPos   string
	Type       string
	State      string
	Timeline   AddWaterMarkTemplateTimeline
	RatioRefer AddWaterMarkTemplateRatioRefer
}

type AddWaterMarkTemplateTimeline struct {
	Start    string
	Duration string
}

type AddWaterMarkTemplateRatioRefer struct {
	Dx     string
	Dy     string
	Width  string
	Height string
}
