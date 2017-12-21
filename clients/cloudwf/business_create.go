package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BusinessCreateRequest struct {
	BusinessCity     string `position:"Query" name:"BusinessCity"`
	Combo            string `position:"Query" name:"Combo"`
	WarnEmail        string `position:"Query" name:"WarnEmail"`
	BusinessManager  string `position:"Query" name:"BusinessManager"`
	BusinessType     int    `position:"Query" name:"BusinessType"`
	Warn             int    `position:"Query" name:"Warn"`
	BusinessName     string `position:"Query" name:"BusinessName"`
	BusinessTopType  int    `position:"Query" name:"BusinessTopType"`
	BusinessAddress  string `position:"Query" name:"BusinessAddress"`
	BusinessTel      string `position:"Query" name:"BusinessTel"`
	BusinessProvince string `position:"Query" name:"BusinessProvince"`
	BusinessSubtype  int    `position:"Query" name:"BusinessSubtype"`
}

func (r BusinessCreateRequest) Invoke(client *sdk.Client) (response *BusinessCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BusinessCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BusinessCreate", "", "")

	resp := struct {
		*responses.BaseResponse
		BusinessCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BusinessCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type BusinessCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
