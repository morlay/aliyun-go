package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaDetailJobList", "mts", "")
	resp = &QueryMediaDetailJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMediaDetailJobListResponse struct {
	responses.BaseResponse
	RequestId   common.String
	JobList     QueryMediaDetailJobListJobList
	NonExistIds QueryMediaDetailJobListNonExistIdList
}

type QueryMediaDetailJobListJob struct {
	Id                common.String
	UserData          common.String
	PipelineId        common.String
	State             common.String
	Code              common.String
	Message           common.String
	CreationTime      common.String
	Input             QueryMediaDetailJobListInput
	MediaDetailConfig QueryMediaDetailJobListMediaDetailConfig
	MediaDetailResult QueryMediaDetailJobListMediaDetailResult
}

type QueryMediaDetailJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryMediaDetailJobListMediaDetailConfig struct {
	Scenario   common.String
	DetailType common.String
	OutputFile QueryMediaDetailJobListOutputFile
}

type QueryMediaDetailJobListOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryMediaDetailJobListMediaDetailResult struct {
	Status                 common.String
	MediaDetailRecgResults QueryMediaDetailJobListMediaDetailRecgResultList
	Tags                   QueryMediaDetailJobListTagList
}

type QueryMediaDetailJobListMediaDetailRecgResult struct {
	ImageUrl      common.String
	Time          common.String
	OcrText       common.String
	Celebrities   QueryMediaDetailJobListCelebrityList
	Sensitives    QueryMediaDetailJobListSensitiveList
	Politicians   QueryMediaDetailJobListPoliticianList
	FrameTagInfos QueryMediaDetailJobListFrameTagInfoList
	FrameTags     QueryMediaDetailJobListFrameTagList
}

type QueryMediaDetailJobListCelebrity struct {
	Name   common.String
	Score  common.String
	Target common.String
}

type QueryMediaDetailJobListSensitive struct {
	Name   common.String
	Score  common.String
	Target common.String
}

type QueryMediaDetailJobListPolitician struct {
	Name   common.String
	Score  common.String
	Target common.String
}

type QueryMediaDetailJobListFrameTagInfo struct {
	Tag      common.String
	Score    common.String
	Category common.String
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

type QueryMediaDetailJobListNonExistIdList []common.String

func (list *QueryMediaDetailJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryMediaDetailJobListTagList []common.String

func (list *QueryMediaDetailJobListTagList) UnmarshalJSON(data []byte) error {
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

type QueryMediaDetailJobListFrameTagList []common.String

func (list *QueryMediaDetailJobListFrameTagList) UnmarshalJSON(data []byte) error {
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
