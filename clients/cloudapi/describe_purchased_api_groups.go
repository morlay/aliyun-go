package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId                   string
	TotalCount                  int
	PageSize                    int
	PageNumber                  int
	PurchasedApiGroupAttributes DescribePurchasedApiGroupsPurchasedApiGroupAttributeList
}

type DescribePurchasedApiGroupsPurchasedApiGroupAttribute struct {
	GroupId        string
	GroupName      string
	Description    string
	PurchasedTime  string
	ExpireTime     string
	RegionId       string
	BillingType    string
	InvokeTimesMax int64
	InvokeTimesNow int64
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
