package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetBatchSaveApAssetProgressRequest struct {
}

func (r GetBatchSaveApAssetProgressRequest) Invoke(client *sdk.Client) (response *GetBatchSaveApAssetProgressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetBatchSaveApAssetProgressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetBatchSaveApAssetProgress", "", "")

	resp := struct {
		*responses.BaseResponse
		GetBatchSaveApAssetProgressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetBatchSaveApAssetProgressResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetBatchSaveApAssetProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
