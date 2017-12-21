package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeOptimizeAdviceOnMissPKRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeOptimizeAdviceOnMissPKRequest) Invoke(client *sdk.Client) (response *DescribeOptimizeAdviceOnMissPKResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeOptimizeAdviceOnMissPKRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceOnMissPK", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeOptimizeAdviceOnMissPKResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeOptimizeAdviceOnMissPKResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeOptimizeAdviceOnMissPKResponse struct {
	RequestId         string
	TotalRecordsCount int
	PageNumber        int
	PageRecordCount   int
	Items             DescribeOptimizeAdviceOnMissPKAdviceOnMissPKList
}

type DescribeOptimizeAdviceOnMissPKAdviceOnMissPK struct {
	DBName    string
	TableName string
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
