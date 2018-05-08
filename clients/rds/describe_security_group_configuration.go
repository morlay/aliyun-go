package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSecurityGroupConfigurationRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeSecurityGroupConfigurationRequest) Invoke(client *sdk.Client) (resp *DescribeSecurityGroupConfigurationResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSecurityGroupConfiguration", "rds", "")
	resp = &DescribeSecurityGroupConfigurationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSecurityGroupConfigurationResponse struct {
	responses.BaseResponse
	RequestId      common.String
	DBInstanceName common.String
	Items          DescribeSecurityGroupConfigurationEcsSecurityGroupRelationList
}

type DescribeSecurityGroupConfigurationEcsSecurityGroupRelation struct {
	RegionId        common.String
	SecurityGroupId common.String
	NetworkType     common.String
}

type DescribeSecurityGroupConfigurationEcsSecurityGroupRelationList []DescribeSecurityGroupConfigurationEcsSecurityGroupRelation

func (list *DescribeSecurityGroupConfigurationEcsSecurityGroupRelationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupConfigurationEcsSecurityGroupRelation)
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
