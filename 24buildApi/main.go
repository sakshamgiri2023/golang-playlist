package main

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:author`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// temp. db
var courses []Course

//middleware

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
}
