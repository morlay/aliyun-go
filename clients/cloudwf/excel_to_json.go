package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ExcelToJsonRequest struct {
	requests.RpcRequest
	UploadData string `position:"Query" name:"UploadData"`
}

func (req *ExcelToJsonRequest) Invoke(client *sdk.Client) (resp *ExcelToJsonResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ExcelToJson", "", "")
	resp = &ExcelToJsonResponse{}
	err = client.DoAction(req, resp)
	return
}

type ExcelToJsonResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
