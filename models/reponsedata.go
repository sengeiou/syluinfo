package models

type ReqLogin struct {
	Username string       `json:"username"`
	Email    string       `json:"email"`
	Token    string       `json:"token"`
	SyluInfo *ReqSyluInfo `json:"syluinfo"`
}

type ReqBind struct {
	Cookie   string       `json:"cookie"`
	SyluInfo *ReqSyluInfo `json:"syluinfo"`
}

type ReqSyluInfo struct {
	ReUsername string `json:"reusername" default:"肖嘉兴"`       //真实姓名
	StudentID  string `json:"studentID" default:"2203050212"` //学号
	Grade      string `json:"grade" default:"2022"`           //年级
	College    string `json:"college" default:"信息科学与工程学院"`    //学院
	Major      string `json:"major" default:"计算机科学与技术(0305)"` //专业
}

type ReqCourse struct {
	StartTime Time         `json:"starttime"`
	Courses   []JsonCourse `json:"courses"`
}

type Time struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type JsonCourse struct {
	Name     string `json:"name"`
	Teacher  string `json:"teacher"`
	Location string `json:"location"`
	Time     int    `json:"time"`
	WeekDay  int    `json:"weekday"`
	WeekS    []int  `json:"weeks"`
}

type JsonGrades struct {
	Name        string  `json:"name"`
	ClassID     string  `json:"classid"`
	Teacher     string  `json:"teacher"`
	IsDegree    bool    `json:"isdegree"`
	Credits     float64 `json:"credits"`     //学分
	GPA         float64 `json:"gpa"`         //绩点
	GradePoints float64 `json:"gradepoints"` //学分绩点
	Fraction    float64 `json:"fraction"`
	Grade       string  `json:"grade"`
}
