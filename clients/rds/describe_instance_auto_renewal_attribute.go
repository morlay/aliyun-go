package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceAutoRenewalAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	ProxyId              string `position:"Query" name:"ProxyId"`
}

func (r DescribeInstanceAutoRenewalAttributeRequest) Invoke(client *sdk.Client) (response *DescribeInstanceAutoRenewalAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceAutoRenewalAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeInstanceAutoRenewalAttribute", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceAutoRenewalAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstanceAutoRenewalAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceAutoRenewalAttributeResponse struct {
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeInstanceAutoRenewalAttributeItemList
}

type DescribeInstanceAutoRenewalAttributeItem struct {
	DBInstanceId string
	RegionId     string
	Duration     int
	Status       string
	AutoRenew    string
}

type DescribeInstanceAutoRenewalAttributeItemList []DescribeInstanceAutoRenewalAttributeItem

func (list *DescribeInstanceAutoRenewalAttributeItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAutoRenewalAttributeItem)
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
