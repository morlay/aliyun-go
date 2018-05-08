package hpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInstancesRequest struct {
	requests.RpcRequest
	TOKEN        string `position:"Query" name:"TOKEN"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (req *DescribeInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeInstancesResponse, err error) {
	req.InitWithApiInfo("HPC", "2016-12-13", "DescribeInstances", "hpc", "")
	resp = &DescribeInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstancesResponse struct {
	responses.BaseResponse
	RequestId common.String
	Instances DescribeInstancesInstanceList
}

type DescribeInstancesInstance struct {
	InstanceId                common.String
	RegionId                  common.String
	InstanceType              DescribeInstancesInstanceType
	PackageId                 DescribeInstancesPackageId
	Status                    DescribeInstancesStatus
	InnerIpAddress            common.String
	JumpserverStatus          DescribeInstancesJumpserverStatus
	JumpserverInnerIpAddress  common.String
	JumpServerPublicIpAddress common.String
	CreationTime              common.String
	ExpireTime                common.String
}

type DescribeInstancesInstanceType struct {
	StringValue common.String
}

type DescribeInstancesPackageId struct {
	StringValue common.String
}

type DescribeInstancesStatus struct {
	StringValue common.String
}

type DescribeInstancesJumpserverStatus struct {
	StringValue common.String
}

type DescribeInstancesInstanceList []DescribeInstancesInstance

func (list *DescribeInstancesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesInstance)
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
