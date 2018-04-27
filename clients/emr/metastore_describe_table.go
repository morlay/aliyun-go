package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MetastoreDescribeTableRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	TableName       string `position:"Query" name:"TableName"`
}

func (req *MetastoreDescribeTableRequest) Invoke(client *sdk.Client) (resp *MetastoreDescribeTableResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "MetastoreDescribeTable", "", "")
	resp = &MetastoreDescribeTableResponse{}
	err = client.DoAction(req, resp)
	return
}

type MetastoreDescribeTableResponse struct {
	responses.BaseResponse
	RequestId             string
	CreateTime            int
	LastAccessTime        int
	LocationUri           string
	InputFormat           string
	OutputFormat          string
	Compressed            bool
	SerializationLib      string
	TableName             string
	DbName                string
	Owner                 string
	TableType             string
	Columns               MetastoreDescribeTableColumnList
	SerdeParameters       MetastoreDescribeTableSerdeParameterList
	StorageDescParameters MetastoreDescribeTableStorageDescParameterList
}

type MetastoreDescribeTableColumn struct {
	Name    string
	Type    string
	Comment string
}

type MetastoreDescribeTableSerdeParameter struct {
	Key   string
	Value string
}

type MetastoreDescribeTableStorageDescParameter struct {
	Key   string
	Value string
}

type MetastoreDescribeTableColumnList []MetastoreDescribeTableColumn

func (list *MetastoreDescribeTableColumnList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MetastoreDescribeTableColumn)
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

type MetastoreDescribeTableSerdeParameterList []MetastoreDescribeTableSerdeParameter

func (list *MetastoreDescribeTableSerdeParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MetastoreDescribeTableSerdeParameter)
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

type MetastoreDescribeTableStorageDescParameterList []MetastoreDescribeTableStorageDescParameter

func (list *MetastoreDescribeTableStorageDescParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MetastoreDescribeTableStorageDescParameter)
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
