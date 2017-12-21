package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCasterLayoutRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	LayoutId      string `position:"Query" name:"LayoutId"`
}

func (req *DeleteCasterLayoutRequest) Invoke(client *sdk.Client) (resp *DeleteCasterLayoutResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterLayout", "", "")
	resp = &DeleteCasterLayoutResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterLayoutResponse struct {
	responses.BaseResponse
	RequestId string
	CasterId  string
	LayoutId  string
}
