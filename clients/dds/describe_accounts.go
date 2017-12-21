package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAccountsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeAccountsRequest) Invoke(client *sdk.Client) (response *DescribeAccountsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAccountsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeAccounts", "dds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAccountsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAccountsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAccountsResponse struct {
	RequestId string
	Accounts  DescribeAccountsAccountList
}

type DescribeAccountsAccount struct {
	DBInstanceId       string
	AccountName        string
	AccountStatus      string
	AccountDescription string
}

type DescribeAccountsAccountList []DescribeAccountsAccount

func (list *DescribeAccountsAccountList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccountsAccount)
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
