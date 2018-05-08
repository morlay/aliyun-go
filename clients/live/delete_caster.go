package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCasterRequest struct {
	requests.RpcRequest
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteCasterRequest) Invoke(client *sdk.Client) (resp *DeleteCasterResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCaster", "live", "")
	resp = &DeleteCasterResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterResponse struct {
	responses.BaseResponse
	RequestId common.String
	CasterId  common.String
}
