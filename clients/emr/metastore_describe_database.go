package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MetastoreDescribeDatabaseRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
}

func (req *MetastoreDescribeDatabaseRequest) Invoke(client *sdk.Client) (resp *MetastoreDescribeDatabaseResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "MetastoreDescribeDatabase", "", "")
	resp = &MetastoreDescribeDatabaseResponse{}
	err = client.DoAction(req, resp)
	return
}

type MetastoreDescribeDatabaseResponse struct {
	responses.BaseResponse
	RequestId   string
	DbName      string
	Description string
	LocationUri string
	Parameters  MetastoreDescribeDatabaseParameterList
}

type MetastoreDescribeDatabaseParameter struct {
	Key   string
	Value string
}

type MetastoreDescribeDatabaseParameterList []MetastoreDescribeDatabaseParameter

func (list *MetastoreDescribeDatabaseParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MetastoreDescribeDatabaseParameter)
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
