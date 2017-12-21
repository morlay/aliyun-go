package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCasterScenesRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	SceneId       string `position:"Query" name:"SceneId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *DescribeCasterScenesRequest) Invoke(client *sdk.Client) (resp *DescribeCasterScenesResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterScenes", "", "")
	resp = &DescribeCasterScenesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCasterScenesResponse struct {
	responses.BaseResponse
	RequestId string
	Total     int
	SceneList DescribeCasterScenesSceneList
}

type DescribeCasterScenesScene struct {
	SceneId      string
	SceneName    string
	OutputType   string
	LayoutId     string
	StreamUrl    string
	Status       int
	StreamInfos  DescribeCasterScenesStreamInfoList
	ComponentIds DescribeCasterScenesComponentIdList
}

type DescribeCasterScenesStreamInfo struct {
	TranscodeConfig string
	VideoFormat     string
	OutputStreamUrl string
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

type DescribeCasterScenesComponentIdList []string

func (list *DescribeCasterScenesComponentIdList) UnmarshalJSON(data []byte) error {
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
