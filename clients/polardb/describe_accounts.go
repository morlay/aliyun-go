package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAccountsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
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
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeAccounts", "polardb", "")

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
	Accounts  DescribeAccountsDBInstanceAccountList
}

type DescribeAccountsDBInstanceAccount struct {
	DBClusterId        string
	AccountName        string
	AccountStatus      string
	AccountDescription string
	AccountType        string
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
