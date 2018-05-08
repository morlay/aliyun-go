package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeOptimizeAdviceOnMissPKRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeOptimizeAdviceOnMissPKRequest) Invoke(client *sdk.Client) (resp *DescribeOptimizeAdviceOnMissPKResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceOnMissPK", "rds", "")
	resp = &DescribeOptimizeAdviceOnMissPKResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeOptimizeAdviceOnMissPKResponse struct {
	responses.BaseResponse
	RequestId         common.String
	TotalRecordsCount common.Integer
	PageNumber        common.Integer
	PageRecordCount   common.Integer
	Items             DescribeOptimizeAdviceOnMissPKAdviceOnMissPKList
}

type DescribeOptimizeAdviceOnMissPKAdviceOnMissPK struct {
	DBName    common.String
	TableName common.String
}

type DescribeOptimizeAdviceOnMissPKAdviceOnMissPKList []DescribeOptimizeAdviceOnMissPKAdviceOnMissPK

func (list *DescribeOptimizeAdviceOnMissPKAdviceOnMissPKList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnMissPKAdviceOnMissPK)
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
