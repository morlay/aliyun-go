package pvtz

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CheckZoneNameRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	ZoneName     string `position:"Query" name:"ZoneName"`
}

func (req *CheckZoneNameRequest) Invoke(client *sdk.Client) (resp *CheckZoneNameResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "CheckZoneName", "pvtz", "")
	resp = &CheckZoneNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckZoneNameResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Check     bool
}
