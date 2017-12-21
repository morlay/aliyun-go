package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateRoomNameRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Screencode  string `position:"Query" name:"Screencode"`
	Drname      string `position:"Query" name:"Drname"`
}

func (r UpdateRoomNameRequest) Invoke(client *sdk.Client) (response *UpdateRoomNameResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateRoomNameRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("ITaaS", "2017-05-05", "UpdateRoomName", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateRoomNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateRoomNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateRoomNameResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList UpdateRoomNameErrorMessageList
}

type UpdateRoomNameErrorMessage struct {
	ErrorMessage string
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
