package models

import (
	"reflect"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
	"errors"
)

type MentorTopic struct {
	Id int `orm:"column(mentor_topic_id);pk;auto"`
	User      *User  `orm:"column(user_id);rel(fk)"`
	Topic     *Topic `orm:"column(topic_id);rel(fk)"`
	Level       int    `orm:"column(level);null"`
	Description string `orm:"column(description);null"`
}

func (t *MentorTopic) TableName() string {
	return "mentor_topic"
}

func init() {
	orm.RegisterModel(new(MentorTopic))
}

// AddTopic insert a new MentorTopic into database and returns
// last inserted Id on MentorTopic.
func AddMentorTopic(m *MentorTopic) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTopicById retrieves MentorTopic by UserId and TopicId. Returns error if
// Id doesn't exist
func GetMentorTopicByIds(userId int, topicId int) (v *MentorTopic, err error) {
	o := orm.NewOrm()
	user, err := GetUserById(userId)
	if err != nil {
		return nil, err
	}
	topic, err := GetTopicById(topicId)
	if err != nil {
		return nil, err
	}
	v = &MentorTopic{User: user, Topic: topic}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTopic retrieves all MentorTopic matches certain condition. Returns empty list if
// no records exist
func GetAllMentorTopic(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MentorTopic))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []MentorTopic
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTopic updates MentorTopic by Id and returns error if
// the record to be updated doesn't exist
func UpdateMentorTopicById(m *MentorTopic) (err error) {
	o := orm.NewOrm()
	v := MentorTopic{User: m.User, Topic: m.Topic}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTopic deletes MentorTopic by UserId and TopicId and returns error if
// the record to be deleted doesn't exist
func DeleteMentorTopic(userId int, topicId int) (err error) {
	o := orm.NewOrm()
	user, err := GetUserById(userId)
	if err != nil {
		return err
	}
	topic, err := GetTopicById(topicId)
	if err != nil {
		return err
	}
	v := MentorTopic{User: user, Topic: topic}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MentorTopic{User: user, Topic: topic}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

