package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Jobs       ListJobsJobInfoList
}

type ListJobsJobInfo struct {
	Id             string
	Name           string
	Owner          string
	Priority       int
	State          string
	SubmitTime     string
	StartTime      string
	LastModifyTime string
	Stdout         string
	Stderr         string
	Comment        string
	ArrayRequest   string
	Resources      ListJobsResources
}

type ListJobsResources struct {
	Nodes int
	Cores int
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
