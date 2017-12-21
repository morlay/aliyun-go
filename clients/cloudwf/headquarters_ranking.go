package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersRankingRequest struct {
	Bid int64 `position:"Query" name:"Bid"`
}

func (r HeadquartersRankingRequest) Invoke(client *sdk.Client) (response *HeadquartersRankingResponse, err error) {
	req := struct {
		*requests.RpcRequest
		HeadquartersRankingRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersRanking", "", "")

	resp := struct {
		*responses.BaseResponse
		HeadquartersRankingResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.HeadquartersRankingResponse

	err = client.DoAction(&req, &resp)
	return
}

type HeadquartersRankingResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
