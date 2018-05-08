package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId             common.String
	CreateTime            common.Integer
	LastAccessTime        common.Integer
	LocationUri           common.String
	InputFormat           common.String
	OutputFormat          common.String
	Compressed            bool
	SerializationLib      common.String
	TableName             common.String
	DbName                common.String
	Owner                 common.String
	TableType             common.String
	Columns               MetastoreDescribeTableColumnList
	SerdeParameters       MetastoreDescribeTableSerdeParameterList
	StorageDescParameters MetastoreDescribeTableStorageDescParameterList
}

type MetastoreDescribeTableColumn struct {
	Name    common.String
	Type    common.String
	Comment common.String
}

type MetastoreDescribeTableSerdeParameter struct {
	Key   common.String
	Value common.String
}

type MetastoreDescribeTableStorageDescParameter struct {
	Key   common.String
	Value common.String
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
