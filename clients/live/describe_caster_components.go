package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCasterComponentsRequest struct {
	requests.RpcRequest
	ComponentId string `position:"Query" name:"ComponentId"`
	CasterId    string `position:"Query" name:"CasterId"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCasterComponentsRequest) Invoke(client *sdk.Client) (resp *DescribeCasterComponentsResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterComponents", "live", "")
	resp = &DescribeCasterComponentsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCasterComponentsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	Total      common.Integer
	Components DescribeCasterComponentsComponentList
}

type DescribeCasterComponentsComponent struct {
	ComponentId         common.String
	ComponentName       common.String
	LocationId          common.String
	ComponentType       common.String
	Effect              common.String
	ComponentLayer      DescribeCasterComponentsComponentLayer
	TextLayerContent    DescribeCasterComponentsTextLayerContent
	ImageLayerContent   DescribeCasterComponentsImageLayerContent
	CaptionLayerContent DescribeCasterComponentsCaptionLayerContent
}

type DescribeCasterComponentsComponentLayer struct {
	HeightNormalized    common.Float
	WidthNormalized     common.Float
	PositionRefer       common.String
	PositionNormalizeds DescribeCasterComponentsPositionNormalizedList
}

type DescribeCasterComponentsTextLayerContent struct {
	Text                  common.String
	Color                 common.String
	FontName              common.String
	SizeNormalized        common.Float
	BorderWidthNormalized common.Float
	BorderColor           common.String
}

type DescribeCasterComponentsImageLayerContent struct {
	MaterialId common.String
}

type DescribeCasterComponentsCaptionLayerContent struct {
	LocationId            common.String
	PtsOffset             common.Integer
	WordsCount            common.Integer
	Color                 common.String
	FontName              common.String
	SizeNormalized        common.Float
	BorderWidthNormalized common.Float
	BorderColor           common.String
}

type DescribeCasterComponentsComponentList []DescribeCasterComponentsComponent

func (list *DescribeCasterComponentsComponentList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterComponentsComponent)
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

type DescribeCasterComponentsPositionNormalizedList []common.String

func (list *DescribeCasterComponentsPositionNormalizedList) UnmarshalJSON(data []byte) error {
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
