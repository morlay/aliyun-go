package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryFacerecogJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	FacerecogJobIds      string `position:"Query" name:"FacerecogJobIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryFacerecogJobListRequest) Invoke(client *sdk.Client) (resp *QueryFacerecogJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryFacerecogJobList", "", "")
	resp = &QueryFacerecogJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryFacerecogJobListResponse struct {
	responses.BaseResponse
	RequestId        string
	FacerecogJobList QueryFacerecogJobListFacerecogJobList
	NonExistIds      QueryFacerecogJobListNonExistIdList
}

type QueryFacerecogJobListFacerecogJob struct {
	Id                   string
	UserData             string
	PipelineId           string
	State                string
	Code                 string
	Message              string
	CreationTime         string
	Input                QueryFacerecogJobListInput
	VideoFacerecogResult QueryFacerecogJobListVideoFacerecogResult
}

type QueryFacerecogJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryFacerecogJobListVideoFacerecogResult struct {
	Facerecogs QueryFacerecogJobListFacerecogList
}

type QueryFacerecogJobListFacerecog struct {
	Time  string
	Faces QueryFacerecogJobListFaceList
}

type QueryFacerecogJobListFace struct {
	Name   string
	Score  string
	Target string
}

type QueryFacerecogJobListFacerecogJobList []QueryFacerecogJobListFacerecogJob

func (list *QueryFacerecogJobListFacerecogJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFacerecogJobListFacerecogJob)
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

type QueryFacerecogJobListNonExistIdList []string

func (list *QueryFacerecogJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryFacerecogJobListFacerecogList []QueryFacerecogJobListFacerecog

func (list *QueryFacerecogJobListFacerecogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFacerecogJobListFacerecog)
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

type QueryFacerecogJobListFaceList []QueryFacerecogJobListFace

func (list *QueryFacerecogJobListFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFacerecogJobListFace)
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
