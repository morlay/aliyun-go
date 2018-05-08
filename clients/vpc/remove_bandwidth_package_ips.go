package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveBandwidthPackageIpsRequest struct {
	requests.RpcRequest
	RemovedIpAddressess  *RemoveBandwidthPackageIpsRemovedIpAddressesList `position:"Query" type:"Repeated" name:"RemovedIpAddresses"`
	ResourceOwnerId      int64                                            `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string                                           `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string                                           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string                                           `position:"Query" name:"ClientToken"`
	OwnerAccount         string                                           `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                            `position:"Query" name:"OwnerId"`
}

func (req *RemoveBandwidthPackageIpsRequest) Invoke(client *sdk.Client) (resp *RemoveBandwidthPackageIpsResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "RemoveBandwidthPackageIps", "vpc", "")
	resp = &RemoveBandwidthPackageIpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveBandwidthPackageIpsResponse struct {
	responses.BaseResponse
	RequestId common.String
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
