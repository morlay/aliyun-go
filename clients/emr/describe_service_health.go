package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeServiceHealthRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DescribeServiceHealthRequest) Invoke(client *sdk.Client) (resp *DescribeServiceHealthResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeServiceHealth", "", "")
	resp = &DescribeServiceHealthResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeServiceHealthResponse struct {
	responses.BaseResponse
	RequestId                 string
	Name                      string
	ComponentHealthResultList DescribeServiceHealthComponentHealthResultList
	HealthResult              DescribeServiceHealthHealthResult
}

type DescribeServiceHealthComponentHealthResult struct {
	Key           string
	PassNumber    int
	ErrorNumber   int
	WarningNumber int
	UnKnownNumber int
}

type DescribeServiceHealthHealthResult struct {
	Key           string
	PassNumber    int
	ErrorNumber   int
	WarningNumber int
	UnKnownNumber int
}

type DescribeServiceHealthComponentHealthResultList []DescribeServiceHealthComponentHealthResult

func (list *DescribeServiceHealthComponentHealthResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeServiceHealthComponentHealthResult)
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
