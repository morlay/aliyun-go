package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	GroupId       common.String
	GroupName     common.String
	Description   common.String
	PurchasedTime common.String
	RegionId      common.String
	Status        common.String
	Domains       DescribePurchasedApiGroupDomainItemList
}

type DescribePurchasedApiGroupDomainItem struct {
	DomainName common.String
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
