package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListJobsRequest struct {
	requests.RpcRequest
	Owner      string `position:"Query" name:"Owner"`
	PageSize   int    `position:"Query" name:"PageSize"`
	ClusterId  string `position:"Query" name:"ClusterId"`
	State      string `position:"Query" name:"State"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListJobsRequest) Invoke(client *sdk.Client) (resp *ListJobsResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListJobs", "ehs", "")
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
	Id             common.String
	Name           common.String
	Owner          common.String
	Priority       common.Integer
	State          common.String
	SubmitTime     common.String
	StartTime      common.String
	LastModifyTime common.String
	Stdout         common.String
	Stderr         common.String
	Comment        common.String
	ArrayRequest   common.String
	Resources      ListJobsResources
}

type ListJobsResources struct {
	Nodes common.Integer
	Cores common.Integer
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
