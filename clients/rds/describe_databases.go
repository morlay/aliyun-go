package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDatabasesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBStatus             string `position:"Query" name:"DBStatus"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDatabasesRequest) Invoke(client *sdk.Client) (response *DescribeDatabasesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDatabasesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDatabases", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDatabasesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDatabasesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDatabasesResponse struct {
	RequestId string
	Databases DescribeDatabasesDatabaseList
}

type DescribeDatabasesDatabase struct {
	DBName           string
	DBInstanceId     string
	Engine           string
	DBStatus         string
	CharacterSetName string
	DBDescription    string
	Accounts         DescribeDatabasesAccountPrivilegeInfoList
}

type DescribeDatabasesAccountPrivilegeInfo struct {
	Account          string
	AccountPrivilege string
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
