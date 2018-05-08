package polardb

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
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeAccountsRequest) Invoke(client *sdk.Client) (resp *DescribeAccountsResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeAccounts", "polardb", "")
	resp = &DescribeAccountsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAccountsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Accounts  DescribeAccountsDBInstanceAccountList
}

type DescribeAccountsDBInstanceAccount struct {
	DBClusterId        common.String
	AccountName        common.String
	AccountStatus      common.String
	AccountDescription common.String
	AccountType        common.String
}

type DescribeAccountsDBInstanceAccountList []DescribeAccountsDBInstanceAccount

func (list *DescribeAccountsDBInstanceAccountList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccountsDBInstanceAccount)
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
