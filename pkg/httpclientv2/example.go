package httpclientv2

import "fmt"

func Example() {
	profile, err := New("http://k8s.goscl.ru/okarpova/my-app/api/v1")
	if err != nil {
		panic(err)
	}

	id, err := profile.Create("John", 25, "john@gmail.com", "+73003002020")
	if err != nil {
		panic(err)
	}

	p, err := profile.Get(id.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(p.ID)
	fmt.Println(p.Age)
	fmt.Println(p.Name)
	fmt.Println(p.Contacts.Email)
	fmt.Println(p.Contacts.Phone)

	var (
		name  = "John Doe"
		age   = 26
		email = "new-john@gmail.com"
		phone = "+73003004000"
	)

	err = profile.Update(id.String(), &name, &age, &email, &phone)
	if err != nil {
		panic(err)
	}

	p, err = profile.Get(id.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(p.ID)
	fmt.Println(p.Age)
	fmt.Println(p.Name)
	fmt.Println(p.Contacts.Email)
	fmt.Println(p.Contacts.Phone)

	err = profile.Delete(id.String())
	if err != nil {
		panic(err)
	}

	_, err = profile.Get(id.String())

	fmt.Println("Get request:", err)
}
