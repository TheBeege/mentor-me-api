package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:mentorid/:menteeid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:mentorid/:menteeid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"],
		beego.ControllerComments{
			Method: "GetManyByMenteeId",
			Router: `/mentee/:menteeid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorRequestController"],
		beego.ControllerComments{
			Method: "GetManyByMentorId",
			Router: `/mentor/:mentorid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorTopicController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorTopicController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorTopicController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorTopicController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorTopicController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorTopicController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorTopicController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorTopicController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorTopicController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:MentorTopicController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:TopicController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:TopicController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/TheBeege/mentor-me-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
