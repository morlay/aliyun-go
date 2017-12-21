package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceAutoRenewAttributeRequest struct {
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

func (r DescribeInstanceAutoRenewAttributeRequest) Invoke(client *sdk.Client) (response *DescribeInstanceAutoRenewAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceAutoRenewAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeInstanceAutoRenewAttribute", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceAutoRenewAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstanceAutoRenewAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceAutoRenewAttributeResponse struct {
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeInstanceAutoRenewAttributeItemList
}

type DescribeInstanceAutoRenewAttributeItem struct {
	DBInstanceId string
	RegionId     string
	Duration     int
	Status       string
	AutoRenew    string
}

type DescribeInstanceAutoRenewAttributeItemList []DescribeInstanceAutoRenewAttributeItem

func (list *DescribeInstanceAutoRenewAttributeItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAutoRenewAttributeItem)
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
