package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "AddMedia", "mts", "")
	resp = &AddMediaResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddMediaResponse struct {
	responses.BaseResponse
	RequestId common.String
	Media     AddMediaMedia
}

type AddMediaMedia struct {
	MediaId      common.String
	Title        common.String
	Description  common.String
	CoverURL     common.String
	CateId       common.Long
	Duration     common.String
	Format       common.String
	Size         common.String
	Bitrate      common.String
	Width        common.String
	Height       common.String
	Fps          common.String
	PublishState common.String
	CreationTime common.String
	Tags         AddMediaTagList
	RunIdList    AddMediaRunIdListList
	File         AddMediaFile
}

type AddMediaFile struct {
	URL   common.String
	State common.String
}

type AddMediaTagList []common.String

func (list *AddMediaTagList) UnmarshalJSON(data []byte) error {
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

type AddMediaRunIdListList []common.String

func (list *AddMediaRunIdListList) UnmarshalJSON(data []byte) error {
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
