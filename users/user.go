package users

import(
	"fmt"
	"time"
	"github.com/cajami/godesde0/modelos"
)

func AltaUsuario(){

	u := new (modelos.User)

	u.AddUser(10, "Javier", time.Now(), true)
	fmt.Println(u)
}