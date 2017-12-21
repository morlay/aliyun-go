package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryAnnotationJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AnnotationJobIds     string `position:"Query" name:"AnnotationJobIds"`
}

func (req *QueryAnnotationJobListRequest) Invoke(client *sdk.Client) (resp *QueryAnnotationJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryAnnotationJobList", "", "")
	resp = &QueryAnnotationJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAnnotationJobListResponse struct {
	responses.BaseResponse
	RequestId         string
	AnnotationJobList QueryAnnotationJobListAnnotationJobList
	NonExistIds       QueryAnnotationJobListNonExistIdList
}

type QueryAnnotationJobListAnnotationJob struct {
	Id                    string
	UserData              string
	PipelineId            string
	State                 string
	Code                  string
	Message               string
	CreationTime          string
	Input                 QueryAnnotationJobListInput
	VideoAnnotationResult QueryAnnotationJobListVideoAnnotationResult
}

type QueryAnnotationJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryAnnotationJobListVideoAnnotationResult struct {
	Details     string
	Annotations QueryAnnotationJobListAnnotationList
}

type QueryAnnotationJobListAnnotation struct {
	Label string
	Score string
}

type QueryAnnotationJobListAnnotationJobList []QueryAnnotationJobListAnnotationJob

func (list *QueryAnnotationJobListAnnotationJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnnotationJobListAnnotationJob)
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

type QueryAnnotationJobListNonExistIdList []string

func (list *QueryAnnotationJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryAnnotationJobListAnnotationList []QueryAnnotationJobListAnnotation

func (list *QueryAnnotationJobListAnnotationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnnotationJobListAnnotation)
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
