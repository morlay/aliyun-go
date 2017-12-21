package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UserDataShowListRequest struct {
	requests.RpcRequest
	Iid  int64  `position:"Query" name:"Iid"`
	Name string `position:"Query" name:"Name"`
	Page int    `position:"Query" name:"Page"`
	Bid  int64  `position:"Query" name:"Bid"`
	Per  int    `position:"Query" name:"Per"`
}

func (req *UserDataShowListRequest) Invoke(client *sdk.Client) (resp *UserDataShowListResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "UserDataShowList", "", "")
	resp = &UserDataShowListResponse{}
	err = client.DoAction(req, resp)
	return
}

type UserDataShowListResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
