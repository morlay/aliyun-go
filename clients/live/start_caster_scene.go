package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartCasterSceneRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	SceneId       string `position:"Query" name:"SceneId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *StartCasterSceneRequest) Invoke(client *sdk.Client) (resp *StartCasterSceneResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "StartCasterScene", "live", "")
	resp = &StartCasterSceneResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartCasterSceneResponse struct {
	responses.BaseResponse
	RequestId string
	StreamUrl string
}
