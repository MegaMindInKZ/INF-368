package main

import (
	"fmt"
	"strconv"
)

var dayOfTheWeeks = map[int]string{
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
	7: "Sunday",
}

type Students []Student

func (s *Students) addStudent(firstName, lastName string) {
	stud := Student{id: len(*s), firstName: firstName, lastName: lastName}
	*s = append(*s, stud)
}

type Teacher struct {
	id        int
	firstName string
	lastName  string
}

type Student struct {
	id        int
	firstName string
	lastName  string
}

type Lesson struct {
	code      string
	title     string
	startTime Time
	endTime   Time
	dayOfWeek int
	teacher   Teacher
	students  []Student
}

type Time struct {
	hour   int
	minute int
}

func (s Student) String() string {
	return s.firstName + " " + s.lastName
}

func (s Teacher) String() string {
	return s.firstName + " " + s.lastName
}

func (t Time) String() string {
	var hour string
	if t.hour < 10 {
		hour = "0" + strconv.Itoa(t.hour)
	} else {
		hour = strconv.Itoa(t.hour)
	}
	var minute string
	if t.minute < 10 {
		minute = "0" + strconv.Itoa(t.minute)
	} else {
		minute = strconv.Itoa(t.minute)
	}
	return hour + ":" + minute
}

func (l Lesson) String() string {
	return l.code + " " + l.title + " " + dayOfTheWeeks[l.dayOfWeek] + " " + l.startTime.String() + "-" + l.endTime.String() + " " + l.teacher.String()
}

func (l *Lesson) addStudent(s Student) {
	l.students = append(l.students, s)
}

func lessonsOfStudent(s Student, l []Lesson) []Lesson {
	ans := []Lesson{}
	for _, les := range l {
		for _, student := range les.students {
			if student.id == s.id {
				ans = append(ans, les)
			}
		}
	}
	return ans
}

func main() {
	var students Students
	students.addStudent("Zanggar", "Zhumagaliyev")
	students.addStudent("Alibi", "Zhumagaliyev")
	students.addStudent("Duman", "Zhumagaliyev")
	students.addStudent("Kanat", "Zhumagaliyev")
	students.addStudent("Daryn", "Zhumagaliyev")
	teacher1 := Teacher{id: 1, firstName: "Azamat", lastName: "Zhumagaliyev"}
	l := Lesson{code: "INF-368", title: "Go", startTime: Time{11, 0}, endTime: Time{12, 0}, dayOfWeek: 1, teacher: teacher1}
	l.addStudent(students[0])
	fmt.Println(lessonsOfStudent(students[0], []Lesson{l}))
}
