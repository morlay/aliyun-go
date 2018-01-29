package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetLicenseRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Data                 string `position:"Query" name:"Data"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	LicenseUrl           string `position:"Query" name:"LicenseUrl"`
}

func (req *GetLicenseRequest) Invoke(client *sdk.Client) (resp *GetLicenseResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "GetLicense", "mts", "")
	resp = &GetLicenseResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetLicenseResponse struct {
	responses.BaseResponse
	RequestId string
	License   string
}
