package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryFacerecogJobList", "mts", "")
	resp = &QueryFacerecogJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryFacerecogJobListResponse struct {
	responses.BaseResponse
	RequestId        common.String
	FacerecogJobList QueryFacerecogJobListFacerecogJobList
	NonExistIds      QueryFacerecogJobListNonExistIdList
}

type QueryFacerecogJobListFacerecogJob struct {
	Id                   common.String
	UserData             common.String
	PipelineId           common.String
	State                common.String
	Code                 common.String
	Message              common.String
	CreationTime         common.String
	Input                QueryFacerecogJobListInput
	VideoFacerecogResult QueryFacerecogJobListVideoFacerecogResult
}

type QueryFacerecogJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryFacerecogJobListVideoFacerecogResult struct {
	Facerecogs QueryFacerecogJobListFacerecogList
}

type QueryFacerecogJobListFacerecog struct {
	Time  common.String
	Faces QueryFacerecogJobListFaceList
}

type QueryFacerecogJobListFace struct {
	Name   common.String
	Score  common.String
	Target common.String
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

type QueryFacerecogJobListNonExistIdList []common.String

func (list *QueryFacerecogJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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
