package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryFpShotJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryFpShotJobListRequest) Invoke(client *sdk.Client) (resp *QueryFpShotJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryFpShotJobList", "mts", "")
	resp = &QueryFpShotJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryFpShotJobListResponse struct {
	responses.BaseResponse
	RequestId     string
	FpShotJobList QueryFpShotJobListFpShotJobList
	NonExistIds   QueryFpShotJobListNonExistIdList
}

type QueryFpShotJobListFpShotJob struct {
	Id           string
	UserData     string
	PipelineId   string
	State        string
	Code         string
	Message      string
	CreationTime string
	InputFile    QueryFpShotJobListInputFile
	FpShotResult QueryFpShotJobListFpShotResult
}

type QueryFpShotJobListInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryFpShotJobListFpShotResult struct {
	FpShots QueryFpShotJobListFpShotList
}

type QueryFpShotJobListFpShot struct {
	PrimaryKey   string
	Similarity   string
	FpShotSlices QueryFpShotJobListFpShotSliceList
}

type QueryFpShotJobListFpShotSlice struct {
	Input       QueryFpShotJobListInput
	Duplication QueryFpShotJobListDuplication
}

type QueryFpShotJobListInput struct {
	Start    string
	Duration string
}

type QueryFpShotJobListDuplication struct {
	Start    string
	Duration string
}

type QueryFpShotJobListFpShotJobList []QueryFpShotJobListFpShotJob

func (list *QueryFpShotJobListFpShotJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFpShotJobListFpShotJob)
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

type QueryFpShotJobListNonExistIdList []string

func (list *QueryFpShotJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryFpShotJobListFpShotList []QueryFpShotJobListFpShot

func (list *QueryFpShotJobListFpShotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFpShotJobListFpShot)
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

type QueryFpShotJobListFpShotSliceList []QueryFpShotJobListFpShotSlice

func (list *QueryFpShotJobListFpShotSliceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFpShotJobListFpShotSlice)
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
