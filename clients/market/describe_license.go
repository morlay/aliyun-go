package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLicenseRequest struct {
	requests.RpcRequest
	LicenseCode string `position:"Query" name:"LicenseCode"`
}

func (req *DescribeLicenseRequest) Invoke(client *sdk.Client) (resp *DescribeLicenseResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "DescribeLicense", "yunmarket", "")
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
	ExtendArray   DescribeLicenseLicenseAttributeList
	ExtendInfo    DescribeLicenseExtendInfo
}

type DescribeLicenseLicenseAttribute struct {
	Code  string
	Value string
}

type DescribeLicenseExtendInfo struct {
	AliUid          int64
	Email           string
	Mobile          string
	AccountQuantity int64
}

type DescribeLicenseLicenseAttributeList []DescribeLicenseLicenseAttribute

func (list *DescribeLicenseLicenseAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLicenseLicenseAttribute)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
