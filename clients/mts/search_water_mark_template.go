package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SearchWaterMarkTemplateRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
}

func (req *SearchWaterMarkTemplateRequest) Invoke(client *sdk.Client) (resp *SearchWaterMarkTemplateResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SearchWaterMarkTemplate", "mts", "")
	resp = &SearchWaterMarkTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchWaterMarkTemplateResponse struct {
	responses.BaseResponse
	RequestId             common.String
	TotalCount            common.Long
	PageNumber            common.Long
	PageSize              common.Long
	WaterMarkTemplateList SearchWaterMarkTemplateWaterMarkTemplateList
}

type SearchWaterMarkTemplateWaterMarkTemplate struct {
	Id         common.String
	Name       common.String
	Width      common.String
	Height     common.String
	Dx         common.String
	Dy         common.String
	ReferPos   common.String
	Type       common.String
	State      common.String
	Timeline   SearchWaterMarkTemplateTimeline
	RatioRefer SearchWaterMarkTemplateRatioRefer
}

type SearchWaterMarkTemplateTimeline struct {
	Start    common.String
	Duration common.String
}

type SearchWaterMarkTemplateRatioRefer struct {
	Dx     common.String
	Dy     common.String
	Width  common.String
	Height common.String
}

type SearchWaterMarkTemplateWaterMarkTemplateList []SearchWaterMarkTemplateWaterMarkTemplate

func (list *SearchWaterMarkTemplateWaterMarkTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchWaterMarkTemplateWaterMarkTemplate)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
