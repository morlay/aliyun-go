package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopCasterSceneRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	SceneId       string `position:"Query" name:"SceneId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *StopCasterSceneRequest) Invoke(client *sdk.Client) (resp *StopCasterSceneResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "StopCasterScene", "live", "")
	resp = &StopCasterSceneResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopCasterSceneResponse struct {
	responses.BaseResponse
	RequestId string
}
