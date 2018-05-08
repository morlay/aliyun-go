package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	NextMarker common.String
	Jobs       ListTagJobsJobsItemList
}

type ListTagJobsJobsItem struct {
	JobId      common.String
	SetId      common.String
	SrcUri     common.String
	Status     common.String
	Percent    common.Integer
	CreateTime common.String
	FinishTime common.String
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
