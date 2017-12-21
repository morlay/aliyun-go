package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySecurityGroupEgressRuleRequest struct {
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

func (r ModifySecurityGroupEgressRuleRequest) Invoke(client *sdk.Client) (response *ModifySecurityGroupEgressRuleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifySecurityGroupEgressRuleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifySecurityGroupEgressRule", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifySecurityGroupEgressRuleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifySecurityGroupEgressRuleResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifySecurityGroupEgressRuleResponse struct {
	RequestId string
}
