package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FileAsyncScanResultsRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *FileAsyncScanResultsRequest) Invoke(client *sdk.Client) (resp *FileAsyncScanResultsResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "FileAsyncScanResults", "/green/file/results", "green", "")
	req.Method = "POST"

	resp = &FileAsyncScanResultsResponse{}
	err = client.DoAction(req, resp)
	return
}

type FileAsyncScanResultsResponse struct {
	responses.BaseResponse
}
