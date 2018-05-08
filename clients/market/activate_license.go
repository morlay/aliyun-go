package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ActivateLicenseRequest struct {
	requests.RpcRequest
	Identification string `position:"Query" name:"Identification"`
	LicenseCode    string `position:"Query" name:"LicenseCode"`
}

func (req *ActivateLicenseRequest) Invoke(client *sdk.Client) (resp *ActivateLicenseResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "ActivateLicense", "yunmarket", "")
	resp = &ActivateLicenseResponse{}
	err = client.DoAction(req, resp)
	return
}

type ActivateLicenseResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
