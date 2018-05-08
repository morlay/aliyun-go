package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeOptimizeAdviceOnStorageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeOptimizeAdviceOnStorageRequest) Invoke(client *sdk.Client) (resp *DescribeOptimizeAdviceOnStorageResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceOnStorage", "rds", "")
	resp = &DescribeOptimizeAdviceOnStorageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeOptimizeAdviceOnStorageResponse struct {
	responses.BaseResponse
	RequestId         common.String
	DBInstanceId      common.String
	TotalRecordsCount common.Integer
	PageNumber        common.Integer
	PageRecordCount   common.Integer
	Items             DescribeOptimizeAdviceOnStorageAdviceOnStorageList
}

type DescribeOptimizeAdviceOnStorageAdviceOnStorage struct {
	DBName        common.String
	TableName     common.String
	CurrentEngine common.String
	AdviseEngine  common.String
}

type DescribeOptimizeAdviceOnStorageAdviceOnStorageList []DescribeOptimizeAdviceOnStorageAdviceOnStorage

func (list *DescribeOptimizeAdviceOnStorageAdviceOnStorageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnStorageAdviceOnStorage)
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
