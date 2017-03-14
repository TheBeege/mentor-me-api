// @APIVersion 1.0.0
// @Title MentorMe API
// @Description Do stuff! Make things!
// @Contact bryanberry235711@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License MIT License
// @LicenseUrl https://opensource.org/licenses/MIT
package routers

import (
	"github.com/TheBeege/mentor-me-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/mentor_topic",
			beego.NSInclude(
				&controllers.MentorTopicController{},
			),
		),

		beego.NSNamespace("/mentor_request",
			beego.NSInclude(
				&controllers.MentorRequestController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/topic",
			beego.NSInclude(
				&controllers.TopicController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
