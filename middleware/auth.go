package middleware

import "github.com/tgo-team/tgo-talkapi/cmd"

func Auth(c *cmd.Context)  {
	openId,err := c.Auth()
	if err!=nil {
		c.Abort()
		c.Error("认证失败！-> %v",err)
		c.ReplyError(401,"认证失败！")
		return
	}
	c.SetCurrentOpenId(openId)
}
