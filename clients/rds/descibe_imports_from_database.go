package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescibeImportsFromDatabaseRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	ImportId             int    `position:"Query" name:"ImportId"`
	Engine               string `position:"Query" name:"Engine"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
}

func (req *DescibeImportsFromDatabaseRequest) Invoke(client *sdk.Client) (resp *DescibeImportsFromDatabaseResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescibeImportsFromDatabase", "rds", "")
	resp = &DescibeImportsFromDatabaseResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescibeImportsFromDatabaseResponse struct {
	responses.BaseResponse
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescibeImportsFromDatabaseImportResultFromDBList
}

type DescibeImportsFromDatabaseImportResultFromDB struct {
	ImportId                    int
	ImportDataType              string
	ImportDataStatus            string
	ImportDataStatusDescription string
	IncrementalImportingTime    string
}

type DescibeImportsFromDatabaseImportResultFromDBList []DescibeImportsFromDatabaseImportResultFromDB

func (list *DescibeImportsFromDatabaseImportResultFromDBList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescibeImportsFromDatabaseImportResultFromDB)
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
