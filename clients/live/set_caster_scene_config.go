package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetCasterSceneConfigRequest struct {
	requests.RpcRequest
	ComponentIds *SetCasterSceneConfigComponentIdList `position:"Query" type:"Repeated" name:"ComponentId"`
	CasterId     string                               `position:"Query" name:"CasterId"`
	SceneId      string                               `position:"Query" name:"SceneId"`
	OwnerId      int64                                `position:"Query" name:"OwnerId"`
	LayoutId     string                               `position:"Query" name:"LayoutId"`
}

func (req *SetCasterSceneConfigRequest) Invoke(client *sdk.Client) (resp *SetCasterSceneConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "SetCasterSceneConfig", "live", "")
	resp = &SetCasterSceneConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetCasterSceneConfigResponse struct {
	responses.BaseResponse
	RequestId string
}

type SetCasterSceneConfigComponentIdList []string

func (list *SetCasterSceneConfigComponentIdList) UnmarshalJSON(data []byte) error {
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
