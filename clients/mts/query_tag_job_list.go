package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	TagJobList  QueryTagJobListTagJobList
	NonExistIds QueryTagJobListNonExistIdList
}

type QueryTagJobListTagJob struct {
	Id             common.String
	UserData       common.String
	PipelineId     common.String
	State          common.String
	Code           common.String
	Message        common.String
	CreationTime   common.String
	Input          QueryTagJobListInput
	VideoTagResult QueryTagJobListVideoTagResult
}

type QueryTagJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryTagJobListVideoTagResult struct {
	Details      common.String
	TagAnResults QueryTagJobListTagAnResultList
	TagFrResults QueryTagJobListTagFrResultList
}

type QueryTagJobListTagAnResult struct {
	Label common.String
	Score common.String
}

type QueryTagJobListTagFrResult struct {
	Time     common.String
	TagFaces QueryTagJobListTagFaceList
}

type QueryTagJobListTagFace struct {
	Name   common.String
	Score  common.String
	Target common.String
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

type QueryTagJobListNonExistIdList []common.String

func (list *QueryTagJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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
