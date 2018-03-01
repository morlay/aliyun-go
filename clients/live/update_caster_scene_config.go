package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateCasterSceneConfigRequest struct {
	requests.RpcRequest
	ComponentIds  *UpdateCasterSceneConfigComponentIdList `position:"Query" type:"Repeated" name:"ComponentId"`
	SecurityToken string                                  `position:"Query" name:"SecurityToken"`
	CasterId      string                                  `position:"Query" name:"CasterId"`
	SceneId       string                                  `position:"Query" name:"SceneId"`
	OwnerId       int64                                   `position:"Query" name:"OwnerId"`
	Version       string                                  `position:"Query" name:"Version"`
	LayoutId      string                                  `position:"Query" name:"LayoutId"`
}

func (req *UpdateCasterSceneConfigRequest) Invoke(client *sdk.Client) (resp *UpdateCasterSceneConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "UpdateCasterSceneConfig", "live", "")
	resp = &UpdateCasterSceneConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateCasterSceneConfigResponse struct {
	responses.BaseResponse
	RequestId string
}

type UpdateCasterSceneConfigComponentIdList []string

func (list *UpdateCasterSceneConfigComponentIdList) UnmarshalJSON(data []byte) error {
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
