package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "SearchWaterMarkTemplate", "", "")
	resp = &SearchWaterMarkTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchWaterMarkTemplateResponse struct {
	responses.BaseResponse
	RequestId             string
	TotalCount            int64
	PageNumber            int64
	PageSize              int64
	WaterMarkTemplateList SearchWaterMarkTemplateWaterMarkTemplateList
}

type SearchWaterMarkTemplateWaterMarkTemplate struct {
	Id         string
	Name       string
	Width      string
	Height     string
	Dx         string
	Dy         string
	ReferPos   string
	Type       string
	State      string
	Timeline   SearchWaterMarkTemplateTimeline
	RatioRefer SearchWaterMarkTemplateRatioRefer
}

type SearchWaterMarkTemplateTimeline struct {
	Start    string
	Duration string
}

type SearchWaterMarkTemplateRatioRefer struct {
	Dx     string
	Dy     string
	Width  string
	Height string
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
