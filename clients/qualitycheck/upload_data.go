package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadDataRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r UploadDataRequest) Invoke(client *sdk.Client) (response *UploadDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UploadDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadData", "", "")

	resp := struct {
		*responses.BaseResponse
		UploadDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UploadDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type UploadDataResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
