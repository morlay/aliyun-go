package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateNatGatewayRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64                                 `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string                                `position:"Query" name:"ClientToken"`
	OwnerAccount         string                                `position:"Query" name:"OwnerAccount"`
	VpcId                string                                `position:"Query" name:"VpcId"`
	Name                 string                                `position:"Query" name:"Name"`
	Description          string                                `position:"Query" name:"Description"`
	OwnerId              int64                                 `position:"Query" name:"OwnerId"`
	BandwidthPackages    *CreateNatGatewayBandwidthPackageList `position:"Query" type:"Repeated" name:"BandwidthPackage"`
	Spec                 string                                `position:"Query" name:"Spec"`
}

func (req *CreateNatGatewayRequest) Invoke(client *sdk.Client) (resp *CreateNatGatewayResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateNatGateway", "vpc", "")
	resp = &CreateNatGatewayResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateNatGatewayBandwidthPackage struct {
	IpCount            int    `name:"IpCount"`
	Bandwidth          int    `name:"Bandwidth"`
	Zone               string `name:"Zone"`
	ISP                string `name:"ISP"`
	InternetChargeType string `name:"InternetChargeType"`
}

type CreateNatGatewayResponse struct {
	responses.BaseResponse
	RequestId           common.String
	NatGatewayId        common.String
	ForwardTableIds     CreateNatGatewayForwardTableIdList
	SnatTableIds        CreateNatGatewaySnatTableIdList
	BandwidthPackageIds CreateNatGatewayBandwidthPackageIdList
}

type CreateNatGatewayBandwidthPackageList []CreateNatGatewayBandwidthPackage

func (list *CreateNatGatewayBandwidthPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateNatGatewayBandwidthPackage)
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

type CreateNatGatewayForwardTableIdList []common.String

func (list *CreateNatGatewayForwardTableIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type CreateNatGatewaySnatTableIdList []common.String

func (list *CreateNatGatewaySnatTableIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type CreateNatGatewayBandwidthPackageIdList []common.String

func (list *CreateNatGatewayBandwidthPackageIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
