package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RegisterNoticeRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r RegisterNoticeRequest) Invoke(client *sdk.Client) (response *RegisterNoticeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RegisterNoticeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "RegisterNotice", "", "")

	resp := struct {
		*responses.BaseResponse
		RegisterNoticeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RegisterNoticeResponse

	err = client.DoAction(&req, &resp)
	return
}

type RegisterNoticeResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
}
