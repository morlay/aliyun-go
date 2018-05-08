package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryWaterMarkTemplateList", "mts", "")
	resp = &QueryWaterMarkTemplateListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryWaterMarkTemplateListResponse struct {
	responses.BaseResponse
	RequestId             common.String
	WaterMarkTemplateList QueryWaterMarkTemplateListWaterMarkTemplateList
	NonExistWids          QueryWaterMarkTemplateListNonExistWidList
}

type QueryWaterMarkTemplateListWaterMarkTemplate struct {
	Id         common.String
	Name       common.String
	Width      common.String
	Height     common.String
	Dx         common.String
	Dy         common.String
	ReferPos   common.String
	Type       common.String
	State      common.String
	Timeline   QueryWaterMarkTemplateListTimeline
	RatioRefer QueryWaterMarkTemplateListRatioRefer
}

type QueryWaterMarkTemplateListTimeline struct {
	Start    common.String
	Duration common.String
}

type QueryWaterMarkTemplateListRatioRefer struct {
	Dx     common.String
	Dy     common.String
	Width  common.String
	Height common.String
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

type QueryWaterMarkTemplateListNonExistWidList []common.String

func (list *QueryWaterMarkTemplateListNonExistWidList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
