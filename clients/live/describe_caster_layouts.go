package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCasterLayoutsRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	LayoutId      string `position:"Query" name:"LayoutId"`
}

func (req *DescribeCasterLayoutsRequest) Invoke(client *sdk.Client) (resp *DescribeCasterLayoutsResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterLayouts", "live", "")
	resp = &DescribeCasterLayoutsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCasterLayoutsResponse struct {
	responses.BaseResponse
	RequestId string
	Total     int
	Layouts   DescribeCasterLayoutsLayoutList
}

type DescribeCasterLayoutsLayout struct {
	LayoutId    string
	VideoLayers DescribeCasterLayoutsVideoLayerList
	AudioLayers DescribeCasterLayoutsAudioLayerList
	BlendList   DescribeCasterLayoutsBlendListList
	MixList     DescribeCasterLayoutsMixListList
}

type DescribeCasterLayoutsVideoLayer struct {
	HeightNormalized    float32
	WidthNormalized     float32
	PositionRefer       string
	PositionNormalizeds DescribeCasterLayoutsPositionNormalizedList
}

type DescribeCasterLayoutsAudioLayer struct {
	VolumeRate   float32
	ValidChannel string
}

type DescribeCasterLayoutsLayoutList []DescribeCasterLayoutsLayout

func (list *DescribeCasterLayoutsLayoutList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterLayoutsLayout)
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

type DescribeCasterLayoutsVideoLayerList []DescribeCasterLayoutsVideoLayer

func (list *DescribeCasterLayoutsVideoLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterLayoutsVideoLayer)
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

type DescribeCasterLayoutsAudioLayerList []DescribeCasterLayoutsAudioLayer

func (list *DescribeCasterLayoutsAudioLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterLayoutsAudioLayer)
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

type DescribeCasterLayoutsBlendListList []string

func (list *DescribeCasterLayoutsBlendListList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterLayoutsMixListList []string

func (list *DescribeCasterLayoutsMixListList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterLayoutsPositionNormalizedList []string

func (list *DescribeCasterLayoutsPositionNormalizedList) UnmarshalJSON(data []byte) error {
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
