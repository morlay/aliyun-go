package pvtz

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddZoneRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	ZoneName     string `position:"Query" name:"ZoneName"`
}

func (req *AddZoneRequest) Invoke(client *sdk.Client) (resp *AddZoneResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "AddZone", "pvtz", "")
	resp = &AddZoneResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddZoneResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	ZoneId    common.String
	ZoneName  common.String
}
