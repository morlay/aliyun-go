package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDatabasesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBStatus             string `position:"Query" name:"DBStatus"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeDatabasesRequest) Invoke(client *sdk.Client) (resp *DescribeDatabasesResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDatabases", "rds", "")
	resp = &DescribeDatabasesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDatabasesResponse struct {
	responses.BaseResponse
	RequestId common.String
	Databases DescribeDatabasesDatabaseList
}

type DescribeDatabasesDatabase struct {
	DBName           common.String
	DBInstanceId     common.String
	Engine           common.String
	DBStatus         common.String
	CharacterSetName common.String
	DBDescription    common.String
	Accounts         DescribeDatabasesAccountPrivilegeInfoList
}

type DescribeDatabasesAccountPrivilegeInfo struct {
	Account                common.String
	AccountPrivilege       common.String
	AccountPrivilegeDetail common.String
}

type DescribeDatabasesDatabaseList []DescribeDatabasesDatabase

func (list *DescribeDatabasesDatabaseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDatabasesDatabase)
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

type DescribeDatabasesAccountPrivilegeInfoList []DescribeDatabasesAccountPrivilegeInfo

func (list *DescribeDatabasesAccountPrivilegeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDatabasesAccountPrivilegeInfo)
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
