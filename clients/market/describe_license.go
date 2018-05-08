package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	License   DescribeLicenseLicense
}

type DescribeLicenseLicense struct {
	LicenseStatus common.String
	LicenseCode   common.String
	InstanceId    common.String
	CreateTime    common.String
	ExpiredTime   common.String
	ActivateTime  common.String
	ProductSkuId  common.String
	ProductCode   common.String
	ProductName   common.String
	ExtendArray   DescribeLicenseLicenseAttributeList
	ExtendInfo    DescribeLicenseExtendInfo
}

type DescribeLicenseLicenseAttribute struct {
	Code  common.String
	Value common.String
}

type DescribeLicenseExtendInfo struct {
	AliUid          common.Long
	Email           common.String
	Mobile          common.String
	AccountQuantity common.Long
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
