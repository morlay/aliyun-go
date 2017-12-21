package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersRankingRequest struct {
	requests.RpcRequest
	Bid int64 `position:"Query" name:"Bid"`
}

func (req *HeadquartersRankingRequest) Invoke(client *sdk.Client) (resp *HeadquartersRankingResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersRanking", "", "")
	resp = &HeadquartersRankingResponse{}
	err = client.DoAction(req, resp)
	return
}

type HeadquartersRankingResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
