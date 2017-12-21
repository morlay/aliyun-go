package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRegionRequest struct {
	Domain string `position:"Query" name:"Domain"`
}

func (r GetRegionRequest) Invoke(client *sdk.Client) (response *GetRegionResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetRegionRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetRegion", "/regions", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetRegionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRegionResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRegionResponse struct {
}
