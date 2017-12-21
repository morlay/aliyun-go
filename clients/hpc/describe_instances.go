package hpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstancesRequest struct {
	TOKEN        string `position:"Query" name:"TOKEN"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (r DescribeInstancesRequest) Invoke(client *sdk.Client) (response *DescribeInstancesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstancesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("HPC", "2016-12-13", "DescribeInstances", "hpc", "")

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
