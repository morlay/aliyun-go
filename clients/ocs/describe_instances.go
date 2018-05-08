package ocs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInstancesRequest struct {
	requests.RpcRequest
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

func (req *DescribeInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeInstancesResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "DescribeInstances", "", "")
	resp = &DescribeInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstancesResponse struct {
	responses.BaseResponse
	RequestId               common.String
	GetOcsInstancesResponse DescribeInstancesGetOcsInstancesResponse
}

type DescribeInstancesGetOcsInstancesResponse struct {
	Total        common.Integer
	PageNo       common.Integer
	PageSize     common.Integer
	OcsInstances DescribeInstancesOcsInstanceList
}

type DescribeInstancesOcsInstance struct {
	OcsInstanceId     common.String
	OcsInstanceName   common.String
	Capacity          common.Long
	Qps               common.Long
	Bandwidth         common.Long
	Connections       common.Long
	ConnectionDomain  common.String
	Port              common.Integer
	UserName          common.String
	RegionId          common.String
	OcsInstanceStatus common.String
	GmtCreated        common.String
	EndTime           common.String
	ChargeType        common.String
	IzId              common.String
	NetworkType       common.String
	VpcId             common.String
	VSwitchId         common.String
	PrivateIp         common.String
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
