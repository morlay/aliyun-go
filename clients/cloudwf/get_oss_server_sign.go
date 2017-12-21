package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetOssServerSignRequest struct {
	requests.RpcRequest
}

func (req *GetOssServerSignRequest) Invoke(client *sdk.Client) (resp *GetOssServerSignResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetOssServerSign", "", "")
	resp = &GetOssServerSignResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetOssServerSignResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
