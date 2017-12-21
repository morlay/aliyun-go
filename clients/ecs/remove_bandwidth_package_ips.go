package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveBandwidthPackageIpsRequest struct {
	RemovedIpAddressess  *RemoveBandwidthPackageIpsRemovedIpAddressesList `position:"Query" type:"Repeated" name:"RemovedIpAddresses"`
	ResourceOwnerId      int64                                            `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string                                           `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string                                           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string                                           `position:"Query" name:"ClientToken"`
	OwnerAccount         string                                           `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                            `position:"Query" name:"OwnerId"`
}

func (r RemoveBandwidthPackageIpsRequest) Invoke(client *sdk.Client) (response *RemoveBandwidthPackageIpsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveBandwidthPackageIpsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "RemoveBandwidthPackageIps", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		RemoveBandwidthPackageIpsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveBandwidthPackageIpsResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveBandwidthPackageIpsResponse struct {
	RequestId string
}

type RemoveBandwidthPackageIpsRemovedIpAddressesList []string

func (list *RemoveBandwidthPackageIpsRemovedIpAddressesList) UnmarshalJSON(data []byte) error {
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
