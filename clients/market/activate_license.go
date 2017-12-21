package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ActivateLicenseRequest struct {
	LicenseCode    string `position:"Query" name:"LicenseCode"`
	Identification string `position:"Query" name:"Identification"`
}

func (r ActivateLicenseRequest) Invoke(client *sdk.Client) (response *ActivateLicenseResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ActivateLicenseRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Market", "2015-11-01", "ActivateLicense", "", "")

	resp := struct {
		*responses.BaseResponse
		ActivateLicenseResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ActivateLicenseResponse

	err = client.DoAction(&req, &resp)
	return
}

type ActivateLicenseResponse struct {
	RequestId string
	Success   bool
}
