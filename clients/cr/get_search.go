package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetSearchRequest struct {
	Keyword  string `position:"Query" name:"Keyword"`
	Page     int    `position:"Query" name:"Page"`
	PageSize int    `position:"Query" name:"PageSize"`
}

func (r GetSearchRequest) Invoke(client *sdk.Client) (response *GetSearchResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetSearchRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetSearch", "/search", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetSearchResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetSearchResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetSearchResponse struct {
}
