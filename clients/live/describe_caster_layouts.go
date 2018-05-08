package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCasterLayoutsRequest struct {
	requests.RpcRequest
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	LayoutId string `position:"Query" name:"LayoutId"`
}

func (req *DescribeCasterLayoutsRequest) Invoke(client *sdk.Client) (resp *DescribeCasterLayoutsResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterLayouts", "live", "")
	resp = &DescribeCasterLayoutsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCasterLayoutsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Total     common.Integer
	Layouts   DescribeCasterLayoutsLayoutList
}

type DescribeCasterLayoutsLayout struct {
	LayoutId    common.String
	VideoLayers DescribeCasterLayoutsVideoLayerList
	AudioLayers DescribeCasterLayoutsAudioLayerList
	BlendList   DescribeCasterLayoutsBlendListList
	MixList     DescribeCasterLayoutsMixListList
}

type DescribeCasterLayoutsVideoLayer struct {
	HeightNormalized    common.Float
	WidthNormalized     common.Float
	PositionRefer       common.String
	FixedDelayDuration  common.Integer
	PositionNormalizeds DescribeCasterLayoutsPositionNormalizedList
}

type DescribeCasterLayoutsAudioLayer struct {
	VolumeRate         common.Float
	ValidChannel       common.String
	FixedDelayDuration common.Integer
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

type DescribeCasterLayoutsBlendListList []common.String

func (list *DescribeCasterLayoutsBlendListList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterLayoutsMixListList []common.String

func (list *DescribeCasterLayoutsMixListList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterLayoutsPositionNormalizedList []common.String

func (list *DescribeCasterLayoutsPositionNormalizedList) UnmarshalJSON(data []byte) error {
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
