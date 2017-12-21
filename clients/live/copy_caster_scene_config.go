package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CopyCasterSceneConfigRequest struct {
	FromSceneId   string `position:"Query" name:"FromSceneId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	ToSceneId     string `position:"Query" name:"ToSceneId"`
}

func (r CopyCasterSceneConfigRequest) Invoke(client *sdk.Client) (response *CopyCasterSceneConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CopyCasterSceneConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "CopyCasterSceneConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		CopyCasterSceneConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CopyCasterSceneConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type CopyCasterSceneConfigResponse struct {
	RequestId string
}
