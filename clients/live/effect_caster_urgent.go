package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EffectCasterUrgentRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	SceneId       string `position:"Query" name:"SceneId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r EffectCasterUrgentRequest) Invoke(client *sdk.Client) (response *EffectCasterUrgentResponse, err error) {
	req := struct {
		*requests.RpcRequest
		EffectCasterUrgentRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "EffectCasterUrgent", "", "")

	resp := struct {
		*responses.BaseResponse
		EffectCasterUrgentResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.EffectCasterUrgentResponse

	err = client.DoAction(&req, &resp)
	return
}

type EffectCasterUrgentResponse struct {
	RequestId string
}
