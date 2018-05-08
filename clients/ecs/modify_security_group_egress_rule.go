package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifySecurityGroupEgressRuleRequest struct {
	requests.RpcRequest
	NicType               string `position:"Query" name:"NicType"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	SourcePortRange       string `position:"Query" name:"SourcePortRange"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	SecurityGroupId       string `position:"Query" name:"SecurityGroupId"`
	Description           string `position:"Query" name:"Description"`
	Policy                string `position:"Query" name:"Policy"`
	PortRange             string `position:"Query" name:"PortRange"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol            string `position:"Query" name:"IpProtocol"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	SourceCidrIp          string `position:"Query" name:"SourceCidrIp"`
	DestGroupId           string `position:"Query" name:"DestGroupId"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	DestGroupOwnerAccount string `position:"Query" name:"DestGroupOwnerAccount"`
	Priority              string `position:"Query" name:"Priority"`
	DestCidrIp            string `position:"Query" name:"DestCidrIp"`
	DestGroupOwnerId      int64  `position:"Query" name:"DestGroupOwnerId"`
}

func (req *ModifySecurityGroupEgressRuleRequest) Invoke(client *sdk.Client) (resp *ModifySecurityGroupEgressRuleResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifySecurityGroupEgressRule", "ecs", "")
	resp = &ModifySecurityGroupEgressRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySecurityGroupEgressRuleResponse struct {
	responses.BaseResponse
	RequestId common.String
}
