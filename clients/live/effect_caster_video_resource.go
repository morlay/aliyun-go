package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EffectCasterVideoResourceRequest struct {
	requests.RpcRequest
	ResourceId    string `position:"Query" name:"ResourceId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	SceneId       string `position:"Query" name:"SceneId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *EffectCasterVideoResourceRequest) Invoke(client *sdk.Client) (resp *EffectCasterVideoResourceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "EffectCasterVideoResource", "live", "")
	resp = &EffectCasterVideoResourceResponse{}
	err = client.DoAction(req, resp)
	return
}

type EffectCasterVideoResourceResponse struct {
	responses.BaseResponse
	RequestId string
}
