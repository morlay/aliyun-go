package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RegisterNoticeRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *RegisterNoticeRequest) Invoke(client *sdk.Client) (resp *RegisterNoticeResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "RegisterNotice", "", "")
	resp = &RegisterNoticeResponse{}
	err = client.DoAction(req, resp)
	return
}

type RegisterNoticeResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
}
