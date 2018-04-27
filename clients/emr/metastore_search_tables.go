package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MetastoreSearchTablesRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	TableName       string `position:"Query" name:"TableName"`
}

func (req *MetastoreSearchTablesRequest) Invoke(client *sdk.Client) (resp *MetastoreSearchTablesResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "MetastoreSearchTables", "", "")
	resp = &MetastoreSearchTablesResponse{}
	err = client.DoAction(req, resp)
	return
}

type MetastoreSearchTablesResponse struct {
	responses.BaseResponse
	RequestId  string
	TableNames MetastoreSearchTablesTableNameList
}

type MetastoreSearchTablesTableNameList []string

func (list *MetastoreSearchTablesTableNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
