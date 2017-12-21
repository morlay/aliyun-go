package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UserDataCreateRequest struct {
	requests.RpcRequest
	UploadFile string `position:"Query" name:"UploadFile"`
	Name       string `position:"Query" name:"Name"`
	Bid        int64  `position:"Query" name:"Bid"`
	Type       string `position:"Query" name:"Type"`
}

func (req *UserDataCreateRequest) Invoke(client *sdk.Client) (resp *UserDataCreateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "UserDataCreate", "", "")
	resp = &UserDataCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type UserDataCreateResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
