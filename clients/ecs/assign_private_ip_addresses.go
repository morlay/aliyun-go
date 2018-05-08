package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AssignPrivateIpAddressesRequest struct {
	requests.RpcRequest
	ResourceOwnerId                int64                                         `position:"Query" name:"ResourceOwnerId"`
	SecondaryPrivateIpAddressCount int                                           `position:"Query" name:"SecondaryPrivateIpAddressCount"`
	ResourceOwnerAccount           string                                        `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                   string                                        `position:"Query" name:"OwnerAccount"`
	OwnerId                        int64                                         `position:"Query" name:"OwnerId"`
	PrivateIpAddresss              *AssignPrivateIpAddressesPrivateIpAddressList `position:"Query" type:"Repeated" name:"PrivateIpAddress"`
	NetworkInterfaceId             string                                        `position:"Query" name:"NetworkInterfaceId"`
}

func (req *AssignPrivateIpAddressesRequest) Invoke(client *sdk.Client) (resp *AssignPrivateIpAddressesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "AssignPrivateIpAddresses", "ecs", "")
	resp = &AssignPrivateIpAddressesResponse{}
	err = client.DoAction(req, resp)
	return
}

type AssignPrivateIpAddressesResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type AssignPrivateIpAddressesPrivateIpAddressList []string

func (list *AssignPrivateIpAddressesPrivateIpAddressList) UnmarshalJSON(data []byte) error {
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
