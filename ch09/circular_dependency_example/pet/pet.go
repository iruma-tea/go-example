package pet

// 循環参照になるためエラー。コメント化

// import "circular_dependency_example/person"

type Pet struct {
	Name      string
	Type      string
	OwnerName string
}

// var owners = map[string]person.Person{ //liststart1
// 	"Bob":   {"Bob", 30, "Fluffy"},
// 	"Julia": {"Julia", 40, "Rex"},
// } //listend1

// func (p Pet) Owner() person.Person {
// 	return owners[p.OwnerName]
// }
