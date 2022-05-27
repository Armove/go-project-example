package cotroller

import (
	"strconv"

	//"github.com/Moonlight-Zhao/go-project-example/service"
	"github.com/Armove/go-project-example/service"
)

func PublishPost(topicIDStr, content string) *PageData {
	//参数转换
	topicId, _ := strconv.ParseInt(topicIDStr, 10, 64)
	//获取service层结果
	postId, err := service.PublishPost(topicId, content)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: map[string]int64{
			"post_id": postId,
		},
	}
}
