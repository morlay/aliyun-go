package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLicenseRequest struct {
	LicenseCode string `position:"Query" name:"LicenseCode"`
}

func (r DescribeLicenseRequest) Invoke(client *sdk.Client) (response *DescribeLicenseResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLicenseRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Market", "2015-11-01", "DescribeLicense", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLicenseResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLicenseResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLicenseResponse struct {
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
