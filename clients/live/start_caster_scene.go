package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartCasterSceneRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	SceneId       string `position:"Query" name:"SceneId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r StartCasterSceneRequest) Invoke(client *sdk.Client) (response *StartCasterSceneResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StartCasterSceneRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "StartCasterScene", "", "")

	resp := struct {
		*responses.BaseResponse
		StartCasterSceneResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StartCasterSceneResponse

	err = client.DoAction(&req, &resp)
	return
}

type StartCasterSceneResponse struct {
	RequestId string
	StreamUrl string
}
