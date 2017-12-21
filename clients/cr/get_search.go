package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetSearchRequest struct {
	requests.RoaRequest
	Keyword  string `position:"Query" name:"Keyword"`
	Page     int    `position:"Query" name:"Page"`
	PageSize int    `position:"Query" name:"PageSize"`
}

func (req *GetSearchRequest) Invoke(client *sdk.Client) (resp *GetSearchResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetSearch", "/search", "", "")
	req.Method = "GET"

	resp = &GetSearchResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetSearchResponse struct {
	responses.BaseResponse
}
