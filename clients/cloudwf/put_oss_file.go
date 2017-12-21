package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PutOssFileRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
}

func (r PutOssFileRequest) Invoke(client *sdk.Client) (response *PutOssFileResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PutOssFileRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "PutOssFile", "", "")

	resp := struct {
		*responses.BaseResponse
		PutOssFileResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PutOssFileResponse

	err = client.DoAction(&req, &resp)
	return
}

type PutOssFileResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
