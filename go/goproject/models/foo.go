package models

//Foo type is json rep
type Foo struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

//GetFoo returns f
func GetFoo() (Foo, error) {
	f := &Foo{Name: "Default foo object", ID: 0}
	return *f, nil
}

//GetAltFoo return modified Foo instance
func GetAltFoo(id int) (Foo, error) {
	f := &Foo{Name: "New ID", ID: id}
	return *f, nil
}

//PostFoo handles Foo posts
func PostFoo(f Foo) (Foo, error) {
	f.Name = "Name modifed by PostFoo"
	return f, nil
}
