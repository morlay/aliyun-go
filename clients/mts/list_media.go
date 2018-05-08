package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListMediaRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	NextPageToken        string `position:"Query" name:"NextPageToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MaximumPageSize      int64  `position:"Query" name:"MaximumPageSize"`
	From                 string `position:"Query" name:"From"`
	To                   string `position:"Query" name:"To"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ListMediaRequest) Invoke(client *sdk.Client) (resp *ListMediaResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ListMedia", "mts", "")
	resp = &ListMediaResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListMediaResponse struct {
	responses.BaseResponse
	RequestId     common.String
	NextPageToken common.String
	MediaList     ListMediaMediaList
}

type ListMediaMedia struct {
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
	Tags         ListMediaTagList
	RunIdList    ListMediaRunIdListList
	File         ListMediaFile
}

type ListMediaFile struct {
	URL   common.String
	State common.String
}

type ListMediaMediaList []ListMediaMedia

func (list *ListMediaMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMediaMedia)
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

type ListMediaTagList []common.String

func (list *ListMediaTagList) UnmarshalJSON(data []byte) error {
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

type ListMediaRunIdListList []common.String

func (list *ListMediaRunIdListList) UnmarshalJSON(data []byte) error {
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
