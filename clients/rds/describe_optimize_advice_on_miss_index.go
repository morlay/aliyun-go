package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeOptimizeAdviceOnMissIndexRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeOptimizeAdviceOnMissIndexRequest) Invoke(client *sdk.Client) (response *DescribeOptimizeAdviceOnMissIndexResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeOptimizeAdviceOnMissIndexRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceOnMissIndex", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeOptimizeAdviceOnMissIndexResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeOptimizeAdviceOnMissIndexResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeOptimizeAdviceOnMissIndexResponse struct {
	RequestId         string
	DBInstanceId      string
	TotalRecordsCount int
	PageNumber        int
	PageRecordCount   int
	Items             DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList
}

type DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex struct {
	DBName      string
	TableName   string
	QueryColumn string
	SQLText     string
}

type DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList []DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex

func (list *DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex)
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
