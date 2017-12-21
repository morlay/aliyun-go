package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetBatchSaveApAssetProgressRequest struct {
	requests.RpcRequest
}

func (req *GetBatchSaveApAssetProgressRequest) Invoke(client *sdk.Client) (resp *GetBatchSaveApAssetProgressResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetBatchSaveApAssetProgress", "", "")
	resp = &GetBatchSaveApAssetProgressResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetBatchSaveApAssetProgressResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
