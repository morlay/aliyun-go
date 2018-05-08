package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCasterLayoutRequest struct {
	requests.RpcRequest
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	LayoutId string `position:"Query" name:"LayoutId"`
}

func (req *DeleteCasterLayoutRequest) Invoke(client *sdk.Client) (resp *DeleteCasterLayoutResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterLayout", "live", "")
	resp = &DeleteCasterLayoutResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterLayoutResponse struct {
	responses.BaseResponse
	RequestId common.String
	CasterId  common.String
	LayoutId  common.String
}
