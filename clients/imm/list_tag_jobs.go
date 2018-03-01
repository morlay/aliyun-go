package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListTagJobsRequest struct {
	requests.RpcRequest
	Condition string `position:"Query" name:"Condition"`
	MaxKeys   int    `position:"Query" name:"MaxKeys"`
	Marker    string `position:"Query" name:"Marker"`
	Project   string `position:"Query" name:"Project"`
}

func (req *ListTagJobsRequest) Invoke(client *sdk.Client) (resp *ListTagJobsResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListTagJobs", "imm", "")
	resp = &ListTagJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTagJobsResponse struct {
	responses.BaseResponse
	RequestId  string
	NextMarker string
	Jobs       ListTagJobsJobsItemList
}

type ListTagJobsJobsItem struct {
	JobId      string
	SetId      string
	SrcUri     string
	Status     string
	Percent    int
	CreateTime string
	FinishTime string
}

type ListTagJobsJobsItemList []ListTagJobsJobsItem

func (list *ListTagJobsJobsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagJobsJobsItem)
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
