package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetCasterSceneConfigRequest struct {
	ComponentIds  *SetCasterSceneConfigComponentIdList `position:"Query" type:"Repeated" name:"ComponentId"`
	SecurityToken string                               `position:"Query" name:"SecurityToken"`
	CasterId      string                               `position:"Query" name:"CasterId"`
	SceneId       string                               `position:"Query" name:"SceneId"`
	OwnerId       int64                                `position:"Query" name:"OwnerId"`
	Version       string                               `position:"Query" name:"Version"`
	LayoutId      string                               `position:"Query" name:"LayoutId"`
}

func (r SetCasterSceneConfigRequest) Invoke(client *sdk.Client) (response *SetCasterSceneConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetCasterSceneConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "SetCasterSceneConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetCasterSceneConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetCasterSceneConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetCasterSceneConfigResponse struct {
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
