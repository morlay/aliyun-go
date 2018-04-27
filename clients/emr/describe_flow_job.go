package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeFlowJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DescribeFlowJobRequest) Invoke(client *sdk.Client) (resp *DescribeFlowJobResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowJob", "", "")
	resp = &DescribeFlowJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFlowJobResponse struct {
	responses.BaseResponse
	RequestId     string
	Id            string
	GmtCreate     int64
	GmtModified   int64
	Name          string
	Description   string
	FailAct       string
	MaxRetry      int
	RetryInterval int64
	Params        string
	ParamConf     string
	EnvConf       string
	RunConf       string
	CategoryId    string
	Mode          string
	Resource      DescribeFlowJobResourceList
}

type DescribeFlowJobResourceList []string

func (list *DescribeFlowJobResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
