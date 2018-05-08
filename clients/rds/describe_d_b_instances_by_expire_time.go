package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBInstancesByExpireTimeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Tags                 string `position:"Query" name:"Tags"`
	Expired              string `position:"Query" name:"Expired"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ExpirePeriod         int    `position:"Query" name:"ExpirePeriod"`
	ProxyId              string `position:"Query" name:"ProxyId"`
}

func (req *DescribeDBInstancesByExpireTimeRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstancesByExpireTimeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstancesByExpireTime", "rds", "")
	resp = &DescribeDBInstancesByExpireTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstancesByExpireTimeResponse struct {
	responses.BaseResponse
	RequestId        common.String
	PageNumber       common.Integer
	TotalRecordCount common.Integer
	PageRecordCount  common.Integer
	Items            DescribeDBInstancesByExpireTimeDBInstanceExpireTimeList
}

type DescribeDBInstancesByExpireTimeDBInstanceExpireTime struct {
	DBInstanceId          common.String
	DBInstanceDescription common.String
	ExpireTime            common.String
	DBInstanceStatus      common.String
	LockMode              common.String
}

type DescribeDBInstancesByExpireTimeDBInstanceExpireTimeList []DescribeDBInstancesByExpireTimeDBInstanceExpireTime

func (list *DescribeDBInstancesByExpireTimeDBInstanceExpireTimeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesByExpireTimeDBInstanceExpireTime)
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
