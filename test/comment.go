package main

import (
	"github.com/airdb/aa/models"
)

func main() {
	var aa models.Comment

	aa.UUID = "82ab3721-a986-44eb-83c7-61112e8f7595"
	aa.AvatarUrl = "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTK476rqoy0zicVAxv1fJ9v9EaltE42WibKS7btmicQBkbwBFkeB6XVSFHF4CGmFlB5OaPeVft7WJcn9w/0"
	aa.NickName = "阿正_Dean"
	aa.AvatarUrl = "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTK476rqoy0zicVAxv1fJ9v9EaltE42WibKS7btmicQBkbwBFkeB6XVSFHF4CGmFlB5OaPeVft7WJcn9w/0"
	aa.Content = "测试评论4"
	aa.Like = 100000
	aa.Blow = 1

	models.AddComment(aa)
	// models.GetComment()

}
