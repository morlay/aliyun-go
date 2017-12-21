package ocs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstancesRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	OcsInstanceStatus    string `position:"Query" name:"OcsInstanceStatus"`
	PageNo               int    `position:"Query" name:"PageNo"`
	PageSize             int    `position:"Query" name:"PageSize"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	VpcId                string `position:"Query" name:"VpcId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	PrivateIps           string `position:"Query" name:"PrivateIps"`
}

func (r DescribeInstancesRequest) Invoke(client *sdk.Client) (response *DescribeInstancesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstancesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "DescribeInstances", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstancesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstancesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstancesResponse struct {
	RequestId               string
	GetOcsInstancesResponse DescribeInstancesGetOcsInstancesResponse
}

type DescribeInstancesGetOcsInstancesResponse struct {
	Total        int
	PageNo       int
	PageSize     int
	OcsInstances DescribeInstancesOcsInstanceList
}

type DescribeInstancesOcsInstance struct {
	OcsInstanceId     string
	OcsInstanceName   string
	Capacity          int64
	Qps               int64
	Bandwidth         int64
	Connections       int64
	ConnectionDomain  string
	Port              int
	UserName          string
	RegionId          string
	OcsInstanceStatus string
	GmtCreated        string
	EndTime           string
	ChargeType        string
	IzId              string
	NetworkType       string
	VpcId             string
	VSwitchId         string
	PrivateIp         string
}

type DescribeInstancesOcsInstanceList []DescribeInstancesOcsInstance

func (list *DescribeInstancesOcsInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesOcsInstance)
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
