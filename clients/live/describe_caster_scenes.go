package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCasterScenesRequest struct {
	requests.RpcRequest
	CasterId string `position:"Query" name:"CasterId"`
	SceneId  string `position:"Query" name:"SceneId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCasterScenesRequest) Invoke(client *sdk.Client) (resp *DescribeCasterScenesResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterScenes", "live", "")
	resp = &DescribeCasterScenesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCasterScenesResponse struct {
	responses.BaseResponse
	RequestId common.String
	Total     common.Integer
	SceneList DescribeCasterScenesSceneList
}

type DescribeCasterScenesScene struct {
	SceneId      common.String
	SceneName    common.String
	OutputType   common.String
	LayoutId     common.String
	StreamUrl    common.String
	Status       common.Integer
	StreamInfos  DescribeCasterScenesStreamInfoList
	ComponentIds DescribeCasterScenesComponentIdList
}

type DescribeCasterScenesStreamInfo struct {
	TranscodeConfig common.String
	VideoFormat     common.String
	OutputStreamUrl common.String
}

type DescribeCasterScenesSceneList []DescribeCasterScenesScene

func (list *DescribeCasterScenesSceneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterScenesScene)
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

type DescribeCasterScenesStreamInfoList []DescribeCasterScenesStreamInfo

func (list *DescribeCasterScenesStreamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterScenesStreamInfo)
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

type DescribeCasterScenesComponentIdList []common.String

func (list *DescribeCasterScenesComponentIdList) UnmarshalJSON(data []byte) error {
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
