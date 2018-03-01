package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCasterRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *DeleteCasterRequest) Invoke(client *sdk.Client) (resp *DeleteCasterResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCaster", "live", "")
	resp = &DeleteCasterResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterResponse struct {
	responses.BaseResponse
	RequestId string
	CasterId  string
}
