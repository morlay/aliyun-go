package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCasterComponentsRequest struct {
	requests.RpcRequest
	ComponentId   string `position:"Query" name:"ComponentId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *DescribeCasterComponentsRequest) Invoke(client *sdk.Client) (resp *DescribeCasterComponentsResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterComponents", "", "")
	resp = &DescribeCasterComponentsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCasterComponentsResponse struct {
	responses.BaseResponse
	RequestId  string
	Total      int
	Components DescribeCasterComponentsComponentList
}

type DescribeCasterComponentsComponent struct {
	ComponentId       string
	ComponentName     string
	LocationId        string
	ComponentType     string
	Effect            string
	ComponentLayer    DescribeCasterComponentsComponentLayer
	TextLayerContent  DescribeCasterComponentsTextLayerContent
	ImageLayerContent DescribeCasterComponentsImageLayerContent
}

type DescribeCasterComponentsComponentLayer struct {
	HeightNormalized    float32
	WidthNormalized     float32
	PositionRefer       string
	PositionNormalizeds DescribeCasterComponentsPositionNormalizedList
}

type DescribeCasterComponentsTextLayerContent struct {
	Text                  string
	Color                 string
	FontName              string
	SizeNormalized        float32
	BorderWidthNormalized float32
	BorderColor           string
}

type DescribeCasterComponentsImageLayerContent struct {
	MaterialId string
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

type DescribeCasterComponentsPositionNormalizedList []string

func (list *DescribeCasterComponentsPositionNormalizedList) UnmarshalJSON(data []byte) error {
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
