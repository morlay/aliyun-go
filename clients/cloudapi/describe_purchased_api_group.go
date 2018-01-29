package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribePurchasedApiGroupRequest struct {
	requests.RpcRequest
	GroupId string `position:"Query" name:"GroupId"`
}

func (req *DescribePurchasedApiGroupRequest) Invoke(client *sdk.Client) (resp *DescribePurchasedApiGroupResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribePurchasedApiGroup", "apigateway", "")
	resp = &DescribePurchasedApiGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePurchasedApiGroupResponse struct {
	responses.BaseResponse
	RequestId     string
	GroupId       string
	GroupName     string
	Description   string
	PurchasedTime string
	RegionId      string
	Status        string
	Domains       DescribePurchasedApiGroupDomainItemList
}

type DescribePurchasedApiGroupDomainItem struct {
	DomainName string
}

type DescribePurchasedApiGroupDomainItemList []DescribePurchasedApiGroupDomainItem

func (list *DescribePurchasedApiGroupDomainItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePurchasedApiGroupDomainItem)
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
