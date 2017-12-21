package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpLoadMapRequest struct {
	requests.RpcRequest
	FileName   string `position:"Query" name:"FileName"`
	UploadId   string `position:"Query" name:"UploadId"`
	ObjectName string `position:"Query" name:"ObjectName"`
	ChunkIndex int    `position:"Query" name:"ChunkIndex"`
	ChunkCnt   int    `position:"Query" name:"ChunkCnt"`
}

func (req *UpLoadMapRequest) Invoke(client *sdk.Client) (resp *UpLoadMapResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "UpLoadMap", "", "")
	resp = &UpLoadMapResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpLoadMapResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
