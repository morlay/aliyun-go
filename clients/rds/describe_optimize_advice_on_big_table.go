package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeOptimizeAdviceOnBigTableRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeOptimizeAdviceOnBigTableRequest) Invoke(client *sdk.Client) (response *DescribeOptimizeAdviceOnBigTableResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeOptimizeAdviceOnBigTableRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceOnBigTable", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeOptimizeAdviceOnBigTableResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeOptimizeAdviceOnBigTableResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeOptimizeAdviceOnBigTableResponse struct {
	RequestId         string
	TotalRecordsCount int
	PageNumber        int
	PageRecordCount   int
	Items             DescribeOptimizeAdviceOnBigTableAdviceOnBigTableList
}

type DescribeOptimizeAdviceOnBigTableAdviceOnBigTable struct {
	DBName    string
	TableName string
	TableSize int64
	DataSize  int64
	IndexSize int64
}

type DescribeOptimizeAdviceOnBigTableAdviceOnBigTableList []DescribeOptimizeAdviceOnBigTableAdviceOnBigTable

func (list *DescribeOptimizeAdviceOnBigTableAdviceOnBigTableList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnBigTableAdviceOnBigTable)
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
