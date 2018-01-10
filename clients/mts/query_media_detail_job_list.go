package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMediaDetailJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryMediaDetailJobListRequest) Invoke(client *sdk.Client) (resp *QueryMediaDetailJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaDetailJobList", "", "")
	resp = &QueryMediaDetailJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMediaDetailJobListResponse struct {
	responses.BaseResponse
	RequestId   string
	JobList     QueryMediaDetailJobListJobList
	NonExistIds QueryMediaDetailJobListNonExistIdList
}

type QueryMediaDetailJobListJob struct {
	Id                string
	UserData          string
	PipelineId        string
	State             string
	Code              string
	Message           string
	CreationTime      string
	Input             QueryMediaDetailJobListInput
	MediaDetailConfig QueryMediaDetailJobListMediaDetailConfig
	MediaDetailResult QueryMediaDetailJobListMediaDetailResult
}

type QueryMediaDetailJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryMediaDetailJobListMediaDetailConfig struct {
	Scenario   string
	DetailType string
	OutputFile QueryMediaDetailJobListOutputFile
}

type QueryMediaDetailJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryMediaDetailJobListMediaDetailResult struct {
	Status                 string
	MediaDetailRecgResults QueryMediaDetailJobListMediaDetailRecgResultList
	Tags                   QueryMediaDetailJobListTagList
}

type QueryMediaDetailJobListMediaDetailRecgResult struct {
	ImageUrl      string
	Time          string
	OcrText       string
	Celebrities   QueryMediaDetailJobListCelebrityList
	Sensitives    QueryMediaDetailJobListSensitiveList
	Politicians   QueryMediaDetailJobListPoliticianList
	FrameTagInfos QueryMediaDetailJobListFrameTagInfoList
	FrameTags     QueryMediaDetailJobListFrameTagList
}

type QueryMediaDetailJobListCelebrity struct {
	Name   string
	Score  string
	Target string
}

type QueryMediaDetailJobListSensitive struct {
	Name   string
	Score  string
	Target string
}

type QueryMediaDetailJobListPolitician struct {
	Name   string
	Score  string
	Target string
}

type QueryMediaDetailJobListFrameTagInfo struct {
	Tag      string
	Score    string
	Category string
}

type QueryMediaDetailJobListJobList []QueryMediaDetailJobListJob

func (list *QueryMediaDetailJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListJob)
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

type QueryMediaDetailJobListNonExistIdList []string

func (list *QueryMediaDetailJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryMediaDetailJobListMediaDetailRecgResultList []QueryMediaDetailJobListMediaDetailRecgResult

func (list *QueryMediaDetailJobListMediaDetailRecgResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListMediaDetailRecgResult)
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

type QueryMediaDetailJobListTagList []string

func (list *QueryMediaDetailJobListTagList) UnmarshalJSON(data []byte) error {
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

type QueryMediaDetailJobListCelebrityList []QueryMediaDetailJobListCelebrity

func (list *QueryMediaDetailJobListCelebrityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListCelebrity)
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

type QueryMediaDetailJobListSensitiveList []QueryMediaDetailJobListSensitive

func (list *QueryMediaDetailJobListSensitiveList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListSensitive)
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

type QueryMediaDetailJobListPoliticianList []QueryMediaDetailJobListPolitician

func (list *QueryMediaDetailJobListPoliticianList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListPolitician)
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

type QueryMediaDetailJobListFrameTagInfoList []QueryMediaDetailJobListFrameTagInfo

func (list *QueryMediaDetailJobListFrameTagInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListFrameTagInfo)
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

type QueryMediaDetailJobListFrameTagList []string

func (list *QueryMediaDetailJobListFrameTagList) UnmarshalJSON(data []byte) error {
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
