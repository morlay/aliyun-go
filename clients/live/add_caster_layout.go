package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddCasterLayoutRequest struct {
	requests.RpcRequest
	BlendLists  *AddCasterLayoutBlendListList  `position:"Query" type:"Repeated" name:"BlendList"`
	AudioLayers *AddCasterLayoutAudioLayerList `position:"Query" type:"Repeated" name:"AudioLayer"`
	VideoLayers *AddCasterLayoutVideoLayerList `position:"Query" type:"Repeated" name:"VideoLayer"`
	CasterId    string                         `position:"Query" name:"CasterId"`
	MixLists    *AddCasterLayoutMixListList    `position:"Query" type:"Repeated" name:"MixList"`
	OwnerId     int64                          `position:"Query" name:"OwnerId"`
}

func (req *AddCasterLayoutRequest) Invoke(client *sdk.Client) (resp *AddCasterLayoutResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddCasterLayout", "live", "")
	resp = &AddCasterLayoutResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCasterLayoutAudioLayer struct {
	VolumeRate         float32 `name:"VolumeRate"`
	ValidChannel       string  `name:"ValidChannel"`
	FixedDelayDuration int     `name:"FixedDelayDuration"`
}

type AddCasterLayoutVideoLayer struct {
	HeightNormalized    float32                                `name:"HeightNormalized"`
	WidthNormalized     float32                                `name:"WidthNormalized"`
	PositionRefer       string                                 `name:"PositionRefer"`
	PositionNormalizeds *AddCasterLayoutPositionNormalizedList `type:"Repeated" name:"PositionNormalized"`
	FixedDelayDuration  int                                    `name:"FixedDelayDuration"`
}

type AddCasterLayoutResponse struct {
	responses.BaseResponse
	RequestId common.String
	LayoutId  common.String
}

type AddCasterLayoutBlendListList []string

func (list *AddCasterLayoutBlendListList) UnmarshalJSON(data []byte) error {
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

type AddCasterLayoutAudioLayerList []AddCasterLayoutAudioLayer

func (list *AddCasterLayoutAudioLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterLayoutAudioLayer)
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

type AddCasterLayoutVideoLayerList []AddCasterLayoutVideoLayer

func (list *AddCasterLayoutVideoLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterLayoutVideoLayer)
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

type AddCasterLayoutMixListList []string

func (list *AddCasterLayoutMixListList) UnmarshalJSON(data []byte) error {
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

type AddCasterLayoutPositionNormalizedList []float32

func (list *AddCasterLayoutPositionNormalizedList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]float32)
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
