package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateMediaRequest struct {
	requests.RpcRequest
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	Title                string `position:"Query" name:"Title"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *UpdateMediaRequest) Invoke(client *sdk.Client) (resp *UpdateMediaResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateMedia", "mts", "")
	resp = &UpdateMediaResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateMediaResponse struct {
	responses.BaseResponse
	RequestId common.String
	Media     UpdateMediaMedia
}

type UpdateMediaMedia struct {
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
	Tags         UpdateMediaTagList
	RunIdList    UpdateMediaRunIdListList
	File         UpdateMediaFile
}

type UpdateMediaFile struct {
	URL   common.String
	State common.String
}

type UpdateMediaTagList []common.String

func (list *UpdateMediaTagList) UnmarshalJSON(data []byte) error {
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

type UpdateMediaRunIdListList []common.String

func (list *UpdateMediaRunIdListList) UnmarshalJSON(data []byte) error {
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
