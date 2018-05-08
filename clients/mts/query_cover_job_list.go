package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	CoverJobList QueryCoverJobListCoverJobList
	NonExistIds  QueryCoverJobListNonExistIdList
}

type QueryCoverJobListCoverJob struct {
	Id             common.String
	UserData       common.String
	PipelineId     common.String
	State          common.String
	Code           common.String
	Message        common.String
	CreationTime   common.String
	CoverImageList QueryCoverJobListCoverImageList
	Input          QueryCoverJobListInput
	CoverConfig    QueryCoverJobListCoverConfig
}

type QueryCoverJobListCoverImage struct {
	Score common.String
	Url   common.String
	Time  common.String
}

type QueryCoverJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryCoverJobListCoverConfig struct {
	OutputFile QueryCoverJobListOutputFile
}

type QueryCoverJobListOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
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

type QueryCoverJobListNonExistIdList []common.String

func (list *QueryCoverJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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
