package hpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstancesInSecurityGroupRequest struct {
	requests.RpcRequest
	TOKEN string `position:"Query" name:"TOKEN"`
}

func (req *DescribeInstancesInSecurityGroupRequest) Invoke(client *sdk.Client) (resp *DescribeInstancesInSecurityGroupResponse, err error) {
	req.InitWithApiInfo("HPC", "2016-12-13", "DescribeInstancesInSecurityGroup", "hpc", "")
	resp = &DescribeInstancesInSecurityGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstancesInSecurityGroupResponse struct {
	responses.BaseResponse
	RequestId string
	Records   DescribeInstancesInSecurityGroupRecordList
}

type DescribeInstancesInSecurityGroupRecord struct {
	RegionId       string
	InstanceId     string
	InnerIpAddress string
}

type DescribeInstancesInSecurityGroupRecordList []DescribeInstancesInSecurityGroupRecord

func (list *DescribeInstancesInSecurityGroupRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesInSecurityGroupRecord)
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
