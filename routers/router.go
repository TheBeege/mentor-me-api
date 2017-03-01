// @APIVersion 1.0.0
// @Title Mentor-Me API
// @Description Mentor-Me is a simple mentor matchmaking app. Users can sign up as a mentor and/or mentee. Mentors can list a skill they'd like to mentor in along with their level of expertise in said skill. Mentees can search for mentors by skill and level. A mentee can submit a request for mentorship upon finding a desired mentor. The mentor can then accept or reject the request. Upon acceptance, the mentor and mentee will receive each other's contact information.
// @Contact bryanberry235711@gmail.com
// @TermsOfServiceUrl #
// @License MIT
// @LicenseUrl https://opensource.org/licenses/MIT
package routers

import (
	"github.com/TheBeege/mentor-me-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/topic",
			beego.NSInclude(
				&controllers.TopicController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
