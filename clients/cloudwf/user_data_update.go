package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UserDataUpdateRequest struct {
	requests.RpcRequest
	Iid        int64  `position:"Query" name:"Iid"`
	UploadFile string `position:"Query" name:"UploadFile"`
	Name       string `position:"Query" name:"Name"`
	Bid        int64  `position:"Query" name:"Bid"`
	Type       string `position:"Query" name:"Type"`
}

func (req *UserDataUpdateRequest) Invoke(client *sdk.Client) (resp *UserDataUpdateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "UserDataUpdate", "", "")
	resp = &UserDataUpdateResponse{}
	err = client.DoAction(req, resp)
	return
}

type UserDataUpdateResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
