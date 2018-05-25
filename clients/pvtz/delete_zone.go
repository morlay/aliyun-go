package pvtz

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteZoneRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	ZoneId       string `position:"Query" name:"ZoneId"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *DeleteZoneRequest) Invoke(client *sdk.Client) (resp *DeleteZoneResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "DeleteZone", "pvtz", "")
	resp = &DeleteZoneResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteZoneResponse struct {
	responses.BaseResponse
	RequestId common.String
	ZoneId    common.String
}
