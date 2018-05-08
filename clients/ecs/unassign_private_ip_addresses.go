package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UnassignPrivateIpAddressesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64                                           `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                          `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                          `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                           `position:"Query" name:"OwnerId"`
	PrivateIpAddresss    *UnassignPrivateIpAddressesPrivateIpAddressList `position:"Query" type:"Repeated" name:"PrivateIpAddress"`
	NetworkInterfaceId   string                                          `position:"Query" name:"NetworkInterfaceId"`
}

func (req *UnassignPrivateIpAddressesRequest) Invoke(client *sdk.Client) (resp *UnassignPrivateIpAddressesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "UnassignPrivateIpAddresses", "ecs", "")
	resp = &UnassignPrivateIpAddressesResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnassignPrivateIpAddressesResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type UnassignPrivateIpAddressesPrivateIpAddressList []string

func (list *UnassignPrivateIpAddressesPrivateIpAddressList) UnmarshalJSON(data []byte) error {
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
