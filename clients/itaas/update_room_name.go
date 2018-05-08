package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateRoomNameRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Screencode  string `position:"Query" name:"Screencode"`
	Drname      string `position:"Query" name:"Drname"`
}

func (req *UpdateRoomNameRequest) Invoke(client *sdk.Client) (resp *UpdateRoomNameResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "UpdateRoomName", "", "")
	resp = &UpdateRoomNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateRoomNameResponse struct {
	responses.BaseResponse
	RequestId common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
	Success   bool
	ErrorList UpdateRoomNameErrorMessageList
}

type UpdateRoomNameErrorMessage struct {
	ErrorMessage common.String
}

type UpdateRoomNameErrorMessageList []UpdateRoomNameErrorMessage

func (list *UpdateRoomNameErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]UpdateRoomNameErrorMessage)
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
