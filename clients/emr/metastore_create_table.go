package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type MetastoreCreateTableRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                           `position:"Query" name:"ResourceOwnerId"`
	FieldDelimiter  string                          `position:"Query" name:"FieldDelimiter"`
	DbName          string                          `position:"Query" name:"DbName"`
	Columns         *MetastoreCreateTableColumnList `position:"Query" type:"Repeated" name:"Column"`
	LocationUri     string                          `position:"Query" name:"LocationUri"`
	TableName       string                          `position:"Query" name:"TableName"`
}

func (req *MetastoreCreateTableRequest) Invoke(client *sdk.Client) (resp *MetastoreCreateTableResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "MetastoreCreateTable", "", "")
	resp = &MetastoreCreateTableResponse{}
	err = client.DoAction(req, resp)
	return
}

type MetastoreCreateTableColumn struct {
	Name    string `name:"Name"`
	Type    string `name:"Type"`
	Comment string `name:"Comment"`
}

type MetastoreCreateTableResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type MetastoreCreateTableColumnList []MetastoreCreateTableColumn

func (list *MetastoreCreateTableColumnList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MetastoreCreateTableColumn)
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
