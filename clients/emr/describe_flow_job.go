package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	Id            common.String
	GmtCreate     common.Long
	GmtModified   common.Long
	Name          common.String
	Description   common.String
	FailAct       common.String
	MaxRetry      common.Integer
	RetryInterval common.Long
	Params        common.String
	ParamConf     common.String
	EnvConf       common.String
	RunConf       common.String
	CategoryId    common.String
	Mode          common.String
	Resource      DescribeFlowJobResourceList
}

type DescribeFlowJobResourceList []common.String

func (list *DescribeFlowJobResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
