package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribePurchasedApiGroupsRequest struct {
	requests.RpcRequest
	PageSize   int `position:"Query" name:"PageSize"`
	PageNumber int `position:"Query" name:"PageNumber"`
}

func (req *DescribePurchasedApiGroupsRequest) Invoke(client *sdk.Client) (resp *DescribePurchasedApiGroupsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribePurchasedApiGroups", "apigateway", "")
	resp = &DescribePurchasedApiGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePurchasedApiGroupsResponse struct {
	responses.BaseResponse
	RequestId                   common.String
	TotalCount                  common.Integer
	PageSize                    common.Integer
	PageNumber                  common.Integer
	PurchasedApiGroupAttributes DescribePurchasedApiGroupsPurchasedApiGroupAttributeList
}

type DescribePurchasedApiGroupsPurchasedApiGroupAttribute struct {
	GroupId        common.String
	GroupName      common.String
	Description    common.String
	PurchasedTime  common.String
	ExpireTime     common.String
	RegionId       common.String
	BillingType    common.String
	InvokeTimesMax common.Long
	InvokeTimesNow common.Long
}

type DescribePurchasedApiGroupsPurchasedApiGroupAttributeList []DescribePurchasedApiGroupsPurchasedApiGroupAttribute

func (list *DescribePurchasedApiGroupsPurchasedApiGroupAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePurchasedApiGroupsPurchasedApiGroupAttribute)
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
