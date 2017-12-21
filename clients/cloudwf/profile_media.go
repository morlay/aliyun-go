package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileMediaRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *ProfileMediaRequest) Invoke(client *sdk.Client) (resp *ProfileMediaResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileMedia", "", "")
	resp = &ProfileMediaResponse{}
	err = client.DoAction(req, resp)
	return
}

type ProfileMediaResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
