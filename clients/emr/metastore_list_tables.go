package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MetastoreListTablesRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
}

func (req *MetastoreListTablesRequest) Invoke(client *sdk.Client) (resp *MetastoreListTablesResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "MetastoreListTables", "", "")
	resp = &MetastoreListTablesResponse{}
	err = client.DoAction(req, resp)
	return
}

type MetastoreListTablesResponse struct {
	responses.BaseResponse
	RequestId  string
	TableNames MetastoreListTablesTableNameList
}

type MetastoreListTablesTableNameList []string

func (list *MetastoreListTablesTableNameList) UnmarshalJSON(data []byte) error {
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
