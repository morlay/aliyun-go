package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListJobsRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	IsDesc          string `position:"Query" name:"IsDesc"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
	PageSize        int    `position:"Query" name:"PageSize"`
}

func (req *ListJobsRequest) Invoke(client *sdk.Client) (resp *ListJobsResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListJobs", "", "")
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
	Id           string
	Name         string
	Type         string
	RunParameter string
	FailAct      string
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
