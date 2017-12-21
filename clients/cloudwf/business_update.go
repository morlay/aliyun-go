package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BusinessUpdateRequest struct {
	Warn             int    `position:"Query" name:"Warn"`
	BusinessCity     string `position:"Query" name:"BusinessCity"`
	WarnEmail        string `position:"Query" name:"WarnEmail"`
	BusinessAddress  string `position:"Query" name:"BusinessAddress"`
	Bid              int64  `position:"Query" name:"Bid"`
	BusinessManager  string `position:"Query" name:"BusinessManager"`
	BusinessProvince string `position:"Query" name:"BusinessProvince"`
}

func (r BusinessUpdateRequest) Invoke(client *sdk.Client) (response *BusinessUpdateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BusinessUpdateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BusinessUpdate", "", "")

	resp := struct {
		*responses.BaseResponse
		BusinessUpdateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BusinessUpdateResponse

	err = client.DoAction(&req, &resp)
	return
}

type BusinessUpdateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
