package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyNetworkInterfaceAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64                                               `position:"Query" name:"ResourceOwnerId"`
	SecurityGroupIds     *ModifyNetworkInterfaceAttributeSecurityGroupIdList `position:"Query" type:"Repeated" name:"SecurityGroupId"`
	Description          string                                              `position:"Query" name:"Description"`
	NetworkInterfaceName string                                              `position:"Query" name:"NetworkInterfaceName"`
	ResourceOwnerAccount string                                              `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                              `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                               `position:"Query" name:"OwnerId"`
	NetworkInterfaceId   string                                              `position:"Query" name:"NetworkInterfaceId"`
}

func (req *ModifyNetworkInterfaceAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyNetworkInterfaceAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyNetworkInterfaceAttribute", "ecs", "")
	resp = &ModifyNetworkInterfaceAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyNetworkInterfaceAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type ModifyNetworkInterfaceAttributeSecurityGroupIdList []string

func (list *ModifyNetworkInterfaceAttributeSecurityGroupIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
