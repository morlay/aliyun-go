package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCasterSceneConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	SceneId       string `position:"Query" name:"SceneId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Type          string `position:"Query" name:"Type"`
	Version       string `position:"Query" name:"Version"`
}

func (req *DeleteCasterSceneConfigRequest) Invoke(client *sdk.Client) (resp *DeleteCasterSceneConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterSceneConfig", "", "")
	resp = &DeleteCasterSceneConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterSceneConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
