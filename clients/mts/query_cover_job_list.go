package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryCoverJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CoverJobIds          string `position:"Query" name:"CoverJobIds"`
}

func (req *QueryCoverJobListRequest) Invoke(client *sdk.Client) (resp *QueryCoverJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryCoverJobList", "mts", "")
	resp = &QueryCoverJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryCoverJobListResponse struct {
	responses.BaseResponse
	RequestId    string
	CoverJobList QueryCoverJobListCoverJobList
	NonExistIds  QueryCoverJobListNonExistIdList
}

type QueryCoverJobListCoverJob struct {
	Id             string
	UserData       string
	PipelineId     string
	State          string
	Code           string
	Message        string
	CreationTime   string
	CoverImageList QueryCoverJobListCoverImageList
	Input          QueryCoverJobListInput
	CoverConfig    QueryCoverJobListCoverConfig
}

type QueryCoverJobListCoverImage struct {
	Score string
	Url   string
	Time  string
}

type QueryCoverJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryCoverJobListCoverConfig struct {
	OutputFile QueryCoverJobListOutputFile
}

type QueryCoverJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryCoverJobListCoverJobList []QueryCoverJobListCoverJob

func (list *QueryCoverJobListCoverJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCoverJobListCoverJob)
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

type QueryCoverJobListNonExistIdList []string

func (list *QueryCoverJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryCoverJobListCoverImageList []QueryCoverJobListCoverImage

func (list *QueryCoverJobListCoverImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCoverJobListCoverImage)
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
