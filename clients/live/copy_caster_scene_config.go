package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CopyCasterSceneConfigRequest struct {
	requests.RpcRequest
	FromSceneId   string `position:"Query" name:"FromSceneId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	ToSceneId     string `position:"Query" name:"ToSceneId"`
}

func (req *CopyCasterSceneConfigRequest) Invoke(client *sdk.Client) (resp *CopyCasterSceneConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "CopyCasterSceneConfig", "", "")
	resp = &CopyCasterSceneConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type CopyCasterSceneConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
