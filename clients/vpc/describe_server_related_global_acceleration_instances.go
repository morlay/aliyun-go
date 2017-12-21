package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeServerRelatedGlobalAccelerationInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ServerType           string `position:"Query" name:"ServerType"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ServerId             string `position:"Query" name:"ServerId"`
}

func (r DescribeServerRelatedGlobalAccelerationInstancesRequest) Invoke(client *sdk.Client) (response *DescribeServerRelatedGlobalAccelerationInstancesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeServerRelatedGlobalAccelerationInstancesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeServerRelatedGlobalAccelerationInstances", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeServerRelatedGlobalAccelerationInstancesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeServerRelatedGlobalAccelerationInstancesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeServerRelatedGlobalAccelerationInstancesResponse struct {
	RequestId                   string
	GlobalAccelerationInstances DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstanceList
}

type DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstance struct {
	RegionId                     string
	GlobalAccelerationInstanceId string
	IpAddress                    string
	ServerIpAddress              string
}

type DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstanceList []DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstance

func (list *DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstance)
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
