package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type GroupsScanlation struct {
	Name         string `orm:"column(name);pk"`
	Description  string `orm:"column(description)"`
	ReleaseDelay int    `orm:"column(release_delay);null"`
}

func (t *GroupsScanlation) TableName() string {
	return "groups_scanlation"
}

func init() {
	orm.RegisterModel(new(GroupsScanlation))
}

// GroupsScanlationNameExists checks to see if name exists
// returns bool
func GroupsScanlationNameExists(name string) (exists bool) {
	o := orm.NewOrm()
	exists = o.QueryTable("groups_scanlation").Filter("Name", name).Exist()
	return
}

// AddGroupsScanlation insert a new GroupsScanlation into database and returns
// last inserted Name on success.
func AddGroupsScanlation(m *GroupsScanlation) ( err error) {
	o := orm.NewOrm()
	_, err = o.Insert(m)
	return
}

// GetGroupsScanlationByName retrieves GroupsScanlation by Name. Returns error if
// Name doesn't exist
func GetGroupsScanlationByName(name string) (v *GroupsScanlation, err error) {
	o := orm.NewOrm()
	v = &GroupsScanlation{Name: name}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGroupsScanlation retrieves all GroupsScanlation matches certain condition. Returns empty list if
// no records exist
func GetAllGroupsScanlation(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GroupsScanlation))
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

	var l []GroupsScanlation
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

// UpdateGroupsScanlation updates GroupsScanlation by Name and returns error if
// the record to be updated doesn't exist
func UpdateGroupsScanlationByName(m *GroupsScanlation) (err error) {
	o := orm.NewOrm()
	v := GroupsScanlation{Name: m.Name}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGroupsScanlation deletes GroupsScanlation by Name and returns error if
// the record to be deleted doesn't exist
func DeleteGroupsScanlation(name string) (err error) {
	o := orm.NewOrm()
	v := GroupsScanlation{Name: name}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GroupsScanlation{Name: name}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
