package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLicenseRequest struct {
	requests.RpcRequest
	LicenseCode string `position:"Query" name:"LicenseCode"`
}

func (req *DescribeLicenseRequest) Invoke(client *sdk.Client) (resp *DescribeLicenseResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "DescribeLicense", "", "")
	resp = &DescribeLicenseResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLicenseResponse struct {
	responses.BaseResponse
	RequestId string
	License   DescribeLicenseLicense
}

type DescribeLicenseLicense struct {
	LicenseStatus string
	LicenseCode   string
	InstanceId    string
	CreateTime    string
	ExpiredTime   string
	ActivateTime  string
	ProductSkuId  string
	ProductCode   string
	ProductName   string
	ExtendInfo    DescribeLicenseExtendInfo
}

type DescribeLicenseExtendInfo struct {
	AliUid          int64
	Email           string
	Mobile          string
	AccountQuantity int64
}
