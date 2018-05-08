package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryAnnotationJobList", "mts", "")
	resp = &QueryAnnotationJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAnnotationJobListResponse struct {
	responses.BaseResponse
	RequestId         common.String
	AnnotationJobList QueryAnnotationJobListAnnotationJobList
	NonExistIds       QueryAnnotationJobListNonExistIdList
}

type QueryAnnotationJobListAnnotationJob struct {
	Id                    common.String
	UserData              common.String
	PipelineId            common.String
	State                 common.String
	Code                  common.String
	Message               common.String
	CreationTime          common.String
	Input                 QueryAnnotationJobListInput
	VideoAnnotationResult QueryAnnotationJobListVideoAnnotationResult
}

type QueryAnnotationJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryAnnotationJobListVideoAnnotationResult struct {
	Details     common.String
	Annotations QueryAnnotationJobListAnnotationList
}

type QueryAnnotationJobListAnnotation struct {
	Label common.String
	Score common.String
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

type QueryAnnotationJobListNonExistIdList []common.String

func (list *QueryAnnotationJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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
