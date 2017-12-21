package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileHistoryListRequest struct {
	requests.RpcRequest
	Idtype int64 `position:"Query" name:"Idtype"`
	Page   int   `position:"Query" name:"Page"`
	Per    int   `position:"Query" name:"Per"`
	Agsid  int64 `position:"Query" name:"Agsid"`
}

func (req *ProfileHistoryListRequest) Invoke(client *sdk.Client) (resp *ProfileHistoryListResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileHistoryList", "", "")
	resp = &ProfileHistoryListResponse{}
	err = client.DoAction(req, resp)
	return
}

type ProfileHistoryListResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
