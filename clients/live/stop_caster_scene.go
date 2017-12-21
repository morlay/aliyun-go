package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopCasterSceneRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	SceneId       string `position:"Query" name:"SceneId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r StopCasterSceneRequest) Invoke(client *sdk.Client) (response *StopCasterSceneResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StopCasterSceneRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "StopCasterScene", "", "")

	resp := struct {
		*responses.BaseResponse
		StopCasterSceneResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StopCasterSceneResponse

	err = client.DoAction(&req, &resp)
	return
}

type StopCasterSceneResponse struct {
	RequestId string
}
