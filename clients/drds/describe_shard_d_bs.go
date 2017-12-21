package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeShardDBsRequest struct {
	requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *DescribeShardDBsRequest) Invoke(client *sdk.Client) (resp *DescribeShardDBsResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeShardDBs", "", "")
	resp = &DescribeShardDBsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeShardDBsResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Data      DescribeShardDBsDbIntancePairList
}

type DescribeShardDBsDbIntancePair struct {
	SubDbName    string
	InstanceName string
}

type DescribeShardDBsDbIntancePairList []DescribeShardDBsDbIntancePair

func (list *DescribeShardDBsDbIntancePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeShardDBsDbIntancePair)
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
