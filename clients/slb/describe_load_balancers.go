package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLoadBalancersRequest struct {
	requests.RpcRequest
	Access_key_id         string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	Address               string `position:"Query" name:"Address"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	NetworkType           string `position:"Query" name:"NetworkType"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	ServerId              string `position:"Query" name:"ServerId"`
	MasterZoneId          string `position:"Query" name:"MasterZoneId"`
	PageNumber            int    `position:"Query" name:"PageNumber"`
	Tags                  string `position:"Query" name:"Tags"`
	ServerIntranetAddress string `position:"Query" name:"ServerIntranetAddress"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	ResourceGroupId       string `position:"Query" name:"ResourceGroupId"`
	LoadBalancerName      string `position:"Query" name:"LoadBalancerName"`
	LoadBalancerId        string `position:"Query" name:"LoadBalancerId"`
	InternetChargeType    string `position:"Query" name:"InternetChargeType"`
	VpcId                 string `position:"Query" name:"VpcId"`
	PageSize              int    `position:"Query" name:"PageSize"`
	AddressType           string `position:"Query" name:"AddressType"`
	SlaveZoneId           string `position:"Query" name:"SlaveZoneId"`
	PayType               string `position:"Query" name:"PayType"`
}

func (req *DescribeLoadBalancersRequest) Invoke(client *sdk.Client) (resp *DescribeLoadBalancersResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancers", "slb", "")
	resp = &DescribeLoadBalancersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLoadBalancersResponse struct {
	responses.BaseResponse
	RequestId     common.String
	PageNumber    common.Integer
	PageSize      common.Integer
	TotalCount    common.Integer
	LoadBalancers DescribeLoadBalancersLoadBalancerList
}

type DescribeLoadBalancersLoadBalancer struct {
	LoadBalancerId     common.String
	LoadBalancerName   common.String
	LoadBalancerStatus common.String
	Address            common.String
	AddressType        common.String
	RegionId           common.String
	RegionIdAlias      common.String
	VSwitchId          common.String
	VpcId              common.String
	NetworkType        common.String
	MasterZoneId       common.String
	SlaveZoneId        common.String
	InternetChargeType common.String
	CreateTime         common.String
	CreateTimeStamp    common.Long
	PayType            common.String
	ResourceGroupId    common.String
}

type DescribeLoadBalancersLoadBalancerList []DescribeLoadBalancersLoadBalancer

func (list *DescribeLoadBalancersLoadBalancerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersLoadBalancer)
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
