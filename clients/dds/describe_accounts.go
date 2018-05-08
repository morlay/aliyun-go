package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAccountsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeAccountsRequest) Invoke(client *sdk.Client) (resp *DescribeAccountsResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeAccounts", "dds", "")
	resp = &DescribeAccountsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAccountsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Accounts  DescribeAccountsAccountList
}

type DescribeAccountsAccount struct {
	DBInstanceId       common.String
	AccountName        common.String
	AccountStatus      common.String
	AccountDescription common.String
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
