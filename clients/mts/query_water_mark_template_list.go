package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryWaterMarkTemplateListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	WaterMarkTemplateIds string `position:"Query" name:"WaterMarkTemplateIds"`
}

func (req *QueryWaterMarkTemplateListRequest) Invoke(client *sdk.Client) (resp *QueryWaterMarkTemplateListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryWaterMarkTemplateList", "", "")
	resp = &QueryWaterMarkTemplateListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryWaterMarkTemplateListResponse struct {
	responses.BaseResponse
	RequestId             string
	WaterMarkTemplateList QueryWaterMarkTemplateListWaterMarkTemplateList
	NonExistWids          QueryWaterMarkTemplateListNonExistWidList
}

type QueryWaterMarkTemplateListWaterMarkTemplate struct {
	Id         string
	Name       string
	Width      string
	Height     string
	Dx         string
	Dy         string
	ReferPos   string
	Type       string
	State      string
	Timeline   QueryWaterMarkTemplateListTimeline
	RatioRefer QueryWaterMarkTemplateListRatioRefer
}

type QueryWaterMarkTemplateListTimeline struct {
	Start    string
	Duration string
}

type QueryWaterMarkTemplateListRatioRefer struct {
	Dx     string
	Dy     string
	Width  string
	Height string
}

type QueryWaterMarkTemplateListWaterMarkTemplateList []QueryWaterMarkTemplateListWaterMarkTemplate

func (list *QueryWaterMarkTemplateListWaterMarkTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryWaterMarkTemplateListWaterMarkTemplate)
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

type QueryWaterMarkTemplateListNonExistWidList []string

func (list *QueryWaterMarkTemplateListNonExistWidList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
