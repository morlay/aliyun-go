package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileDistrictRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r ProfileDistrictRequest) Invoke(client *sdk.Client) (response *ProfileDistrictResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ProfileDistrictRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileDistrict", "", "")

	resp := struct {
		*responses.BaseResponse
		ProfileDistrictResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ProfileDistrictResponse

	err = client.DoAction(&req, &resp)
	return
}

type ProfileDistrictResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
