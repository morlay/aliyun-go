package pvtz

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateZoneRemarkRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	ZoneId       string `position:"Query" name:"ZoneId"`
	Remark       string `position:"Query" name:"Remark"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *UpdateZoneRemarkRequest) Invoke(client *sdk.Client) (resp *UpdateZoneRemarkResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "UpdateZoneRemark", "pvtz", "")
	resp = &UpdateZoneRemarkResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateZoneRemarkResponse struct {
	responses.BaseResponse
	RequestId common.String
	ZoneId    common.String
}
