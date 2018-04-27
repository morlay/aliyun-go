package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId     string
	PageNumber    int
	PageSize      int
	TotalCount    int
	LoadBalancers DescribeLoadBalancersLoadBalancerList
}

type DescribeLoadBalancersLoadBalancer struct {
	LoadBalancerId     string
	LoadBalancerName   string
	LoadBalancerStatus string
	Address            string
	AddressType        string
	RegionId           string
	RegionIdAlias      string
	VSwitchId          string
	VpcId              string
	NetworkType        string
	MasterZoneId       string
	SlaveZoneId        string
	InternetChargeType string
	CreateTime         string
	CreateTimeStamp    int64
	PayType            string
	ResourceGroupId    string
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
