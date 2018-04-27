package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartCasterRequest struct {
	requests.RpcRequest
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
}

func (req *StartCasterRequest) Invoke(client *sdk.Client) (resp *StartCasterResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "StartCaster", "live", "")
	resp = &StartCasterResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartCasterResponse struct {
	responses.BaseResponse
	RequestId     string
	PvwSceneInfos StartCasterSceneInfoList
	PgmSceneInfos StartCasterSceneInfo1List
}

type StartCasterSceneInfo struct {
	SceneId   string
	StreamUrl string
}

type StartCasterSceneInfo1 struct {
	SceneId     string
	StreamUrl   string
	StreamInfos StartCasterStreamInfoList
}

type StartCasterStreamInfo struct {
	TranscodeConfig string
	VideoFormat     string
	OutputStreamUrl string
}

type StartCasterSceneInfoList []StartCasterSceneInfo

func (list *StartCasterSceneInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterSceneInfo)
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

type StartCasterSceneInfo1List []StartCasterSceneInfo1

func (list *StartCasterSceneInfo1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterSceneInfo1)
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

type StartCasterStreamInfoList []StartCasterStreamInfo

func (list *StartCasterStreamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterStreamInfo)
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
