package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryTagJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	TagJobIds            string `position:"Query" name:"TagJobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryTagJobListRequest) Invoke(client *sdk.Client) (resp *QueryTagJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryTagJobList", "mts", "")
	resp = &QueryTagJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTagJobListResponse struct {
	responses.BaseResponse
	RequestId   string
	TagJobList  QueryTagJobListTagJobList
	NonExistIds QueryTagJobListNonExistIdList
}

type QueryTagJobListTagJob struct {
	Id             string
	UserData       string
	PipelineId     string
	State          string
	Code           string
	Message        string
	CreationTime   string
	Input          QueryTagJobListInput
	VideoTagResult QueryTagJobListVideoTagResult
}

type QueryTagJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryTagJobListVideoTagResult struct {
	Details      string
	TagAnResults QueryTagJobListTagAnResultList
	TagFrResults QueryTagJobListTagFrResultList
}

type QueryTagJobListTagAnResult struct {
	Label string
	Score string
}

type QueryTagJobListTagFrResult struct {
	Time     string
	TagFaces QueryTagJobListTagFaceList
}

type QueryTagJobListTagFace struct {
	Name   string
	Score  string
	Target string
}

type QueryTagJobListTagJobList []QueryTagJobListTagJob

func (list *QueryTagJobListTagJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagJob)
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

type QueryTagJobListNonExistIdList []string

func (list *QueryTagJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryTagJobListTagAnResultList []QueryTagJobListTagAnResult

func (list *QueryTagJobListTagAnResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagAnResult)
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

type QueryTagJobListTagFrResultList []QueryTagJobListTagFrResult

func (list *QueryTagJobListTagFrResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagFrResult)
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

type QueryTagJobListTagFaceList []QueryTagJobListTagFace

func (list *QueryTagJobListTagFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagFace)
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
