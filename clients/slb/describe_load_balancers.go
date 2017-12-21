package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLoadBalancersRequest struct {
	Access_key_id         string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	NetworkType           string `position:"Query" name:"NetworkType"`
	MasterZoneId          string `position:"Query" name:"MasterZoneId"`
	PageNumber            int    `position:"Query" name:"PageNumber"`
	ResourceGroupId       string `position:"Query" name:"ResourceGroupId"`
	LoadBalancerName      string `position:"Query" name:"LoadBalancerName"`
	PageSize              int    `position:"Query" name:"PageSize"`
	AddressType           string `position:"Query" name:"AddressType"`
	SlaveZoneId           string `position:"Query" name:"SlaveZoneId"`
	Address               string `position:"Query" name:"Address"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	ServerId              string `position:"Query" name:"ServerId"`
	Tags                  string `position:"Query" name:"Tags"`
	ServerIntranetAddress string `position:"Query" name:"ServerIntranetAddress"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	LoadBalancerId        string `position:"Query" name:"LoadBalancerId"`
	InternetChargeType    string `position:"Query" name:"InternetChargeType"`
	VpcId                 string `position:"Query" name:"VpcId"`
	PayType               string `position:"Query" name:"PayType"`
}

func (r DescribeLoadBalancersRequest) Invoke(client *sdk.Client) (response *DescribeLoadBalancersResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLoadBalancersRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancers", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLoadBalancersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLoadBalancersResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLoadBalancersResponse struct {
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
