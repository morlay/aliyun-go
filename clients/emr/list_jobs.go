package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListJobsRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	QueryString     string `position:"Query" name:"QueryString"`
	IsDesc          string `position:"Query" name:"IsDesc"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
	QueryType       string `position:"Query" name:"QueryType"`
}

func (req *ListJobsRequest) Invoke(client *sdk.Client) (resp *ListJobsResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListJobs", "", "")
	resp = &ListJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListJobsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Jobs       ListJobsJobInfoList
}

type ListJobsJobInfo struct {
	Id            common.String
	Name          common.String
	Type          common.String
	RunParameter  common.String
	FailAct       common.String
	MaxRetry      common.Integer
	RetryInterval common.Integer
}

type ListJobsJobInfoList []ListJobsJobInfo

func (list *ListJobsJobInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobsJobInfo)
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
