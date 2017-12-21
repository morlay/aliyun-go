package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetResultCountRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r GetResultCountRequest) Invoke(client *sdk.Client) (response *GetResultCountResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetResultCountRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetResultCount", "", "")

	resp := struct {
		*responses.BaseResponse
		GetResultCountResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetResultCountResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetResultCountResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      int
}
