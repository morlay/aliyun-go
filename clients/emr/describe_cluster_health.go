package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterHealthRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DescribeClusterHealthRequest) Invoke(client *sdk.Client) (resp *DescribeClusterHealthResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterHealth", "", "")
	resp = &DescribeClusterHealthResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterHealthResponse struct {
	responses.BaseResponse
	RequestId             string
	ClusterId             int64
	ServiceHealthInfoList DescribeClusterHealthServiceHealthInfoList
	HealthResult          DescribeClusterHealthHealthResult
}

type DescribeClusterHealthServiceHealthInfo struct {
	Key           string
	PassNumber    int
	ErrorNumber   int
	WarningNumber int
	UnKnownNumber int
}

type DescribeClusterHealthHealthResult struct {
	Key           string
	PassNumber    int
	ErrorNumber   int
	WarningNumber int
	UnKnownNumber int
}

type DescribeClusterHealthServiceHealthInfoList []DescribeClusterHealthServiceHealthInfo

func (list *DescribeClusterHealthServiceHealthInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterHealthServiceHealthInfo)
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
