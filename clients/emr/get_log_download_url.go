package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetLogDownloadUrlRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostName        string `position:"Query" name:"HostName"`
	LogstoreName    string `position:"Query" name:"LogstoreName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	LogFileName     string `position:"Query" name:"LogFileName"`
}

func (req *GetLogDownloadUrlRequest) Invoke(client *sdk.Client) (resp *GetLogDownloadUrlResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetLogDownloadUrl", "", "")
	resp = &GetLogDownloadUrlResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetLogDownloadUrlResponse struct {
	responses.BaseResponse
	RequestId                string
	DownloadUrlBase64Encoded string
}
