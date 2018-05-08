package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifySecurityGroupConfigurationRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifySecurityGroupConfigurationRequest) Invoke(client *sdk.Client) (resp *ModifySecurityGroupConfigurationResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifySecurityGroupConfiguration", "rds", "")
	resp = &ModifySecurityGroupConfigurationResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySecurityGroupConfigurationResponse struct {
	responses.BaseResponse
	RequestId      common.String
	DBInstanceName common.String
	Items          ModifySecurityGroupConfigurationEcsSecurityGroupRelationList
}

type ModifySecurityGroupConfigurationEcsSecurityGroupRelation struct {
	RegionId        common.String
	SecurityGroupId common.String
	NetworkType     common.String
}

type ModifySecurityGroupConfigurationEcsSecurityGroupRelationList []ModifySecurityGroupConfigurationEcsSecurityGroupRelation

func (list *ModifySecurityGroupConfigurationEcsSecurityGroupRelationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifySecurityGroupConfigurationEcsSecurityGroupRelation)
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
