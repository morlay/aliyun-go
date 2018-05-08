package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyCasterLayoutRequest struct {
	requests.RpcRequest
	BlendLists  *ModifyCasterLayoutBlendListList  `position:"Query" type:"Repeated" name:"BlendList"`
	AudioLayers *ModifyCasterLayoutAudioLayerList `position:"Query" type:"Repeated" name:"AudioLayer"`
	VideoLayers *ModifyCasterLayoutVideoLayerList `position:"Query" type:"Repeated" name:"VideoLayer"`
	CasterId    string                            `position:"Query" name:"CasterId"`
	MixLists    *ModifyCasterLayoutMixListList    `position:"Query" type:"Repeated" name:"MixList"`
	OwnerId     int64                             `position:"Query" name:"OwnerId"`
	LayoutId    string                            `position:"Query" name:"LayoutId"`
}

func (req *ModifyCasterLayoutRequest) Invoke(client *sdk.Client) (resp *ModifyCasterLayoutResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "ModifyCasterLayout", "live", "")
	resp = &ModifyCasterLayoutResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCasterLayoutAudioLayer struct {
	VolumeRate         float32 `name:"VolumeRate"`
	ValidChannel       string  `name:"ValidChannel"`
	FixedDelayDuration int     `name:"FixedDelayDuration"`
}

type ModifyCasterLayoutVideoLayer struct {
	HeightNormalized    float32                                   `name:"HeightNormalized"`
	WidthNormalized     float32                                   `name:"WidthNormalized"`
	PositionRefer       string                                    `name:"PositionRefer"`
	PositionNormalizeds *ModifyCasterLayoutPositionNormalizedList `type:"Repeated" name:"PositionNormalized"`
	FixedDelayDuration  int                                       `name:"FixedDelayDuration"`
}

type ModifyCasterLayoutResponse struct {
	responses.BaseResponse
	RequestId common.String
	LayoutId  common.String
}

type ModifyCasterLayoutBlendListList []string

func (list *ModifyCasterLayoutBlendListList) UnmarshalJSON(data []byte) error {
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

type ModifyCasterLayoutAudioLayerList []ModifyCasterLayoutAudioLayer

func (list *ModifyCasterLayoutAudioLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyCasterLayoutAudioLayer)
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

type ModifyCasterLayoutVideoLayerList []ModifyCasterLayoutVideoLayer

func (list *ModifyCasterLayoutVideoLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyCasterLayoutVideoLayer)
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

type ModifyCasterLayoutMixListList []string

func (list *ModifyCasterLayoutMixListList) UnmarshalJSON(data []byte) error {
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

type ModifyCasterLayoutPositionNormalizedList []float32

func (list *ModifyCasterLayoutPositionNormalizedList) UnmarshalJSON(data []byte) error {
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
