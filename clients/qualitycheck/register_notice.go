package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Code      string
	Message   string
}
