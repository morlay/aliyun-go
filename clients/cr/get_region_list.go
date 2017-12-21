package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRegionListRequest struct {
}

func (r GetRegionListRequest) Invoke(client *sdk.Client) (response *GetRegionListResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetRegionListRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetRegionList", "/regions", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetRegionListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRegionListResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRegionListResponse struct {
}
