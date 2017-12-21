package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeShardDBsRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r DescribeShardDBsRequest) Invoke(client *sdk.Client) (response *DescribeShardDBsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeShardDBsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeShardDBs", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeShardDBsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeShardDBsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeShardDBsResponse struct {
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
