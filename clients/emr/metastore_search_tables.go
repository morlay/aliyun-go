package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	TableNames MetastoreSearchTablesTableNameList
}

type MetastoreSearchTablesTableNameList []common.String

func (list *MetastoreSearchTablesTableNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
