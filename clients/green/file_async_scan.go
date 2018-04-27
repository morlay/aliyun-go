package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FileAsyncScanRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *FileAsyncScanRequest) Invoke(client *sdk.Client) (resp *FileAsyncScanResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "FileAsyncScan", "/green/file/asyncscan", "green", "")
	req.Method = "POST"

	resp = &FileAsyncScanResponse{}
	err = client.DoAction(req, resp)
	return
}

type FileAsyncScanResponse struct {
	responses.BaseResponse
}
