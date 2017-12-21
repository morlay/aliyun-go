package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ActivateLicenseRequest struct {
	requests.RpcRequest
	LicenseCode    string `position:"Query" name:"LicenseCode"`
	Identification string `position:"Query" name:"Identification"`
}

func (req *ActivateLicenseRequest) Invoke(client *sdk.Client) (resp *ActivateLicenseResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "ActivateLicense", "", "")
	resp = &ActivateLicenseResponse{}
	err = client.DoAction(req, resp)
	return
}

type ActivateLicenseResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
