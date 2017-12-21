package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateMediaRequest struct {
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

func (r UpdateMediaRequest) Invoke(client *sdk.Client) (response *UpdateMediaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateMediaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateMedia", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateMediaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateMediaResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateMediaResponse struct {
	RequestId string
	Media     UpdateMediaMedia
}

type UpdateMediaMedia struct {
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
	Tags         UpdateMediaTagList
	RunIdList    UpdateMediaRunIdListList
	File         UpdateMediaFile
}

type UpdateMediaFile struct {
	URL   string
	State string
}

type UpdateMediaTagList []string

func (list *UpdateMediaTagList) UnmarshalJSON(data []byte) error {
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

type UpdateMediaRunIdListList []string

func (list *UpdateMediaRunIdListList) UnmarshalJSON(data []byte) error {
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
