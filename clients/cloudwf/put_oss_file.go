package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PutOssFileRequest struct {
	requests.RpcRequest
	JsonData string `position:"Query" name:"JsonData"`
}

func (req *PutOssFileRequest) Invoke(client *sdk.Client) (resp *PutOssFileResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "PutOssFile", "", "")
	resp = &PutOssFileResponse{}
	err = client.DoAction(req, resp)
	return
}

type PutOssFileResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
