package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListFaceJobsRequest struct {
	requests.RpcRequest
	Condition string `position:"Query" name:"Condition"`
	MaxKeys   int    `position:"Query" name:"MaxKeys"`
	Marker    string `position:"Query" name:"Marker"`
	Project   string `position:"Query" name:"Project"`
}

func (req *ListFaceJobsRequest) Invoke(client *sdk.Client) (resp *ListFaceJobsResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListFaceJobs", "imm", "")
	resp = &ListFaceJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFaceJobsResponse struct {
	responses.BaseResponse
	RequestId  string
	NextMarker string
	Jobs       ListFaceJobsJobsItemList
}

type ListFaceJobsJobsItem struct {
	JobId      string
	SetId      string
	SrcUri     string
	Status     string
	Percent    int
	CreateTime string
	FinishTime string
}

type ListFaceJobsJobsItemList []ListFaceJobsJobsItem

func (list *ListFaceJobsJobsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFaceJobsJobsItem)
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
