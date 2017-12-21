package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySecurityGroupRuleRequest struct {
	requests.RpcRequest
	NicType                 string `position:"Query" name:"NicType"`
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	SourcePortRange         string `position:"Query" name:"SourcePortRange"`
	ClientToken             string `position:"Query" name:"ClientToken"`
	SecurityGroupId         string `position:"Query" name:"SecurityGroupId"`
	Description             string `position:"Query" name:"Description"`
	SourceGroupOwnerId      int64  `position:"Query" name:"SourceGroupOwnerId"`
	SourceGroupOwnerAccount string `position:"Query" name:"SourceGroupOwnerAccount"`
	Policy                  string `position:"Query" name:"Policy"`
	PortRange               string `position:"Query" name:"PortRange"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol              string `position:"Query" name:"IpProtocol"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	SourceCidrIp            string `position:"Query" name:"SourceCidrIp"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	Priority                string `position:"Query" name:"Priority"`
	DestCidrIp              string `position:"Query" name:"DestCidrIp"`
	SourceGroupId           string `position:"Query" name:"SourceGroupId"`
}

func (req *ModifySecurityGroupRuleRequest) Invoke(client *sdk.Client) (resp *ModifySecurityGroupRuleResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifySecurityGroupRule", "ecs", "")
	resp = &ModifySecurityGroupRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySecurityGroupRuleResponse struct {
	responses.BaseResponse
	RequestId string
}
