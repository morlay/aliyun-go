package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MetastoreListDatabasesRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
}

func (req *MetastoreListDatabasesRequest) Invoke(client *sdk.Client) (resp *MetastoreListDatabasesResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "MetastoreListDatabases", "", "")
	resp = &MetastoreListDatabasesResponse{}
	err = client.DoAction(req, resp)
	return
}

type MetastoreListDatabasesResponse struct {
	responses.BaseResponse
	RequestId   string
	Description string
	DbNames     MetastoreListDatabasesDbNameList
}

type MetastoreListDatabasesDbNameList []string

func (list *MetastoreListDatabasesDbNameList) UnmarshalJSON(data []byte) error {
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
