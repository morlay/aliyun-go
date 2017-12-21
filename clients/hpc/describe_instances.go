package hpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Instances DescribeInstancesInstanceList
}

type DescribeInstancesInstance struct {
	InstanceId                string
	RegionId                  string
	InstanceType              DescribeInstancesInstanceType
	PackageId                 DescribeInstancesPackageId
	Status                    DescribeInstancesStatus
	InnerIpAddress            string
	JumpserverStatus          DescribeInstancesJumpserverStatus
	JumpserverInnerIpAddress  string
	JumpServerPublicIpAddress string
	CreationTime              string
	ExpireTime                string
}

type DescribeInstancesInstanceType struct {
	StringValue string
}

type DescribeInstancesPackageId struct {
	StringValue string
}

type DescribeInstancesStatus struct {
	StringValue string
}

type DescribeInstancesJumpserverStatus struct {
	StringValue string
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
