package entity

type Category struct {
	ID       int
	Name     string
	CourseID []int
}

type Course struct {
	ID         int
	Name       string
	CategoryID int
}

func (c *Category) AddCouse(id int) {
	c.CourseID = append(c.CourseID, id)
}
