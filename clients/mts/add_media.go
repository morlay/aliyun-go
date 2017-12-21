package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddMediaRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	Description           string `position:"Query" name:"Description"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	Title                 string `position:"Query" name:"Title"`
	Tags                  string `position:"Query" name:"Tags"`
	CoverURL              string `position:"Query" name:"CoverURL"`
	CateId                int64  `position:"Query" name:"CateId"`
	FileURL               string `position:"Query" name:"FileURL"`
	MediaWorkflowId       string `position:"Query" name:"MediaWorkflowId"`
	MediaWorkflowUserData string `position:"Query" name:"MediaWorkflowUserData"`
}

func (req *AddMediaRequest) Invoke(client *sdk.Client) (resp *AddMediaResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddMedia", "", "")
	resp = &AddMediaResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddMediaResponse struct {
	responses.BaseResponse
	RequestId string
	Media     AddMediaMedia
}

type AddMediaMedia struct {
	MediaId      string
	Title        string
	Description  string
	CoverURL     string
	CateId       int64
	Duration     string
	Format       string
	Size         string
	Bitrate      string
	Width        string
	Height       string
	Fps          string
	PublishState string
	CreationTime string
	Tags         AddMediaTagList
	RunIdList    AddMediaRunIdListList
	File         AddMediaFile
}

type AddMediaFile struct {
	URL   string
	State string
}

type AddMediaTagList []string

func (list *AddMediaTagList) UnmarshalJSON(data []byte) error {
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

type AddMediaRunIdListList []string

func (list *AddMediaRunIdListList) UnmarshalJSON(data []byte) error {
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
