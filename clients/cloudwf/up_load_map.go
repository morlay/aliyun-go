package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpLoadMapRequest struct {
	FileName   string `position:"Query" name:"FileName"`
	UploadId   string `position:"Query" name:"UploadId"`
	ObjectName string `position:"Query" name:"ObjectName"`
	ChunkIndex int    `position:"Query" name:"ChunkIndex"`
	ChunkCnt   int    `position:"Query" name:"ChunkCnt"`
}

func (r UpLoadMapRequest) Invoke(client *sdk.Client) (response *UpLoadMapResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpLoadMapRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "UpLoadMap", "", "")

	resp := struct {
		*responses.BaseResponse
		UpLoadMapResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpLoadMapResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpLoadMapResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
