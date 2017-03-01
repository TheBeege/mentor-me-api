package models

type MentorTopic struct {
	UserId      *User  `orm:"column(user_id);rel(fk)"`
	TopicId     *Topic `orm:"column(topic_id);rel(fk)"`
	Level       int    `orm:"column(level);null"`
	Description string `orm:"column(description);null"`
}
