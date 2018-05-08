package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeEipAddressesRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Filter2Value           string `position:"Query" name:"Filter.2.Value"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	AllocationId           string `position:"Query" name:"AllocationId"`
	Filter1Value           string `position:"Query" name:"Filter.1.Value"`
	Filter2Key             string `position:"Query" name:"Filter.2.Key"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	EipAddress             string `position:"Query" name:"EipAddress"`
	PageNumber             int    `position:"Query" name:"PageNumber"`
	LockReason             string `position:"Query" name:"LockReason"`
	Filter1Key             string `position:"Query" name:"Filter.1.Key"`
	AssociatedInstanceType string `position:"Query" name:"AssociatedInstanceType"`
	PageSize               int    `position:"Query" name:"PageSize"`
	ChargeType             string `position:"Query" name:"ChargeType"`
	AssociatedInstanceId   string `position:"Query" name:"AssociatedInstanceId"`
	Status                 string `position:"Query" name:"Status"`
}

func (req *DescribeEipAddressesRequest) Invoke(client *sdk.Client) (resp *DescribeEipAddressesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeEipAddresses", "ecs", "")
	resp = &DescribeEipAddressesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeEipAddressesResponse struct {
	responses.BaseResponse
	RequestId    common.String
	TotalCount   common.Integer
	PageNumber   common.Integer
	PageSize     common.Integer
	EipAddresses DescribeEipAddressesEipAddressList
}

type DescribeEipAddressesEipAddress struct {
	RegionId           common.String
	IpAddress          common.String
	AllocationId       common.String
	Status             common.String
	InstanceId         common.String
	Bandwidth          common.String
	EipBandwidth       common.String
	InternetChargeType common.String
	AllocationTime     common.String
	InstanceType       common.String
	ChargeType         common.String
	ExpiredTime        common.String
	OperationLocks     DescribeEipAddressesLockReasonList
}

type DescribeEipAddressesLockReason struct {
	LockReason common.String
}

type DescribeEipAddressesEipAddressList []DescribeEipAddressesEipAddress

func (list *DescribeEipAddressesEipAddressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEipAddressesEipAddress)
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

type DescribeEipAddressesLockReasonList []DescribeEipAddressesLockReason

func (list *DescribeEipAddressesLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEipAddressesLockReason)
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
