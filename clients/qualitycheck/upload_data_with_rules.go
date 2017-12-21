package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadDataWithRulesRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r UploadDataWithRulesRequest) Invoke(client *sdk.Client) (response *UploadDataWithRulesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UploadDataWithRulesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadDataWithRules", "", "")

	resp := struct {
		*responses.BaseResponse
		UploadDataWithRulesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UploadDataWithRulesResponse

	err = client.DoAction(&req, &resp)
	return
}

type UploadDataWithRulesResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
