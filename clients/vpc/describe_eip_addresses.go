package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	ResourceGroupId        string `position:"Query" name:"ResourceGroupId"`
	LockReason             string `position:"Query" name:"LockReason"`
	Filter1Key             string `position:"Query" name:"Filter.1.Key"`
	AssociatedInstanceType string `position:"Query" name:"AssociatedInstanceType"`
	PageSize               int    `position:"Query" name:"PageSize"`
	ChargeType             string `position:"Query" name:"ChargeType"`
	AssociatedInstanceId   string `position:"Query" name:"AssociatedInstanceId"`
	Status                 string `position:"Query" name:"Status"`
}

func (req *DescribeEipAddressesRequest) Invoke(client *sdk.Client) (resp *DescribeEipAddressesResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeEipAddresses", "vpc", "")
	resp = &DescribeEipAddressesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeEipAddressesResponse struct {
	responses.BaseResponse
	RequestId    string
	TotalCount   int
	PageNumber   int
	PageSize     int
	EipAddresses DescribeEipAddressesEipAddressList
}

type DescribeEipAddressesEipAddress struct {
	RegionId             string
	IpAddress            string
	AllocationId         string
	Status               string
	InstanceId           string
	Bandwidth            string
	EipBandwidth         string
	InternetChargeType   string
	AllocationTime       string
	InstanceType         string
	InstanceRegionId     string
	ChargeType           string
	ExpiredTime          string
	Name                 string
	ISP                  string
	Descritpion          string
	BandwidthPackageId   string
	BandwidthPackageType string
	ResourceGroupId      string
	OperationLocks       DescribeEipAddressesLockReasonList
	AvailableRegions     DescribeEipAddressesAvailableRegionList
}

type DescribeEipAddressesLockReason struct {
	LockReason string
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

type DescribeEipAddressesAvailableRegionList []string

func (list *DescribeEipAddressesAvailableRegionList) UnmarshalJSON(data []byte) error {
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
