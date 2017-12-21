package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ExcelToJsonRequest struct {
	UploadData string `position:"Query" name:"UploadData"`
}

func (r ExcelToJsonRequest) Invoke(client *sdk.Client) (response *ExcelToJsonResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ExcelToJsonRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ExcelToJson", "", "")

	resp := struct {
		*responses.BaseResponse
		ExcelToJsonResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ExcelToJsonResponse

	err = client.DoAction(&req, &resp)
	return
}

type ExcelToJsonResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
