package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UserDataDeleteRequest struct {
	requests.RpcRequest
	Iid int64 `position:"Query" name:"Iid"`
	Bid int64 `position:"Query" name:"Bid"`
}

func (req *UserDataDeleteRequest) Invoke(client *sdk.Client) (resp *UserDataDeleteResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "UserDataDelete", "", "")
	resp = &UserDataDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type UserDataDeleteResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
