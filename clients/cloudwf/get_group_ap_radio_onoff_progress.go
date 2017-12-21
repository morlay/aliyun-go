package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupApRadioOnoffProgressRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetGroupApRadioOnoffProgressRequest) Invoke(client *sdk.Client) (resp *GetGroupApRadioOnoffProgressResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetGroupApRadioOnoffProgress", "", "")
	resp = &GetGroupApRadioOnoffProgressResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetGroupApRadioOnoffProgressResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
