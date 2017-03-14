package models

import (
	"time"
	"reflect"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
	"errors"
)

type MentorRequest struct {
	Id int `orm:"column(mentor_request_id);pk"`
	Mentor      *User  `orm:"column(mentor_id);rel(fk)"`
	Mentee     *User `orm:"column(mentee_id);rel(fk)"`
	Requested time.Time `orm:"column(requested);type(timestamp without time zone)"`
	Accepted  time.Time `orm:"column(accepted);type(timestamp without time zone);null"`
	Rejected  time.Time `orm:"column(rejected);type(timestamp without time zone);null"`
}

func (t *MentorRequest) TableName() string {
	return "mentor_request"
}

func init() {
	orm.RegisterModel(new(MentorRequest))
}

// AddMentorRequest insert a new MentorRequest into database and returns
// last inserted Id on success.
func AddMentorRequest(m *MentorRequest) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMentorRequestByBothIds retrieves MentorRequest by MenteeId. Returns error if
// MenteeId doesn't exist
func GetMentorRequestByBothIds(mentorId int, menteeId int) (v *MentorRequest, err error) {
	o := orm.NewOrm()
	mentor, err := GetUserById(mentorId)
	if err != nil {
		return nil, err
	}
	mentee, err := GetUserById(menteeId)
	if err != nil {
		return nil, err
	}
	v = &MentorRequest{Mentor: mentor, Mentee: mentee}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetMentorRequestByMentorId retrieves MentorRequest by MentorId. Returns error if
// MentorId doesn't exist
func GetMentorRequestsByMentorId(id int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MentorRequest))
	qs.Filter("mentor_id", id)
	var l []*MentorRequest
	if _, err = qs.All(&l); err == nil {
		for _, v := range l {
			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}

// GetMentorRequestByMenteeId retrieves MentorRequest by MenteeId. Returns error if
// MenteeId doesn't exist
func GetMentorRequestsByMenteeId(id int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MentorRequest))
	qs.Filter("mentee_id", id)
	var l []*MentorRequest
	if _, err = qs.All(&l); err == nil {
		for _, v := range l {
			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}

// GetAllMentorRequest retrieves all MentorRequest matches certain condition. Returns empty list if
// no records exist
func GetAllMentorRequest(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MentorRequest))
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

	var l []MentorRequest
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

// UpdateMentorRequest updates MentorRequest by Id and returns error if
// the record to be updated doesn't exist
func UpdateMentorRequest(m *MentorRequest) (err error) {
	o := orm.NewOrm()
	v := MentorRequest{Mentor: m.Mentor, Mentee: m.Mentee}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMentorRequest deletes MentorRequest by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMentorRequest(mentorId int, menteeId int) (err error) {
	o := orm.NewOrm()
	mentor, err := GetUserById(mentorId)
	if err != nil {
		return err
	}
	mentee, err := GetUserById(menteeId)
	if err != nil {
		return err
	}
	v := MentorRequest{Mentor: mentor, Mentee: mentee}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MentorRequest{Mentor: mentor, Mentee: mentee}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
