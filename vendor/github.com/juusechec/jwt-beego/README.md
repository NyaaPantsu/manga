[![Build Status](https://travis-ci.org/juusechec/jwt-beego.svg?branch=master)](https://travis-ci.org/juusechec/jwt-beego)
[![Go Report Card](https://goreportcard.com/badge/github.com/juusechec/jwt-beego)](https://goreportcard.com/report/github.com/juusechec/jwt-beego)
[![GoDoc Reference](https://godoc.org/github.com/juusechec/jwt-beego?status.svg)](http://godoc.org/github.com/juusechec/jwt-beego)
[![Coverage Status](https://coveralls.io/repos/juusechec/jwt-beego/badge.svg?branch=master)](https://coveralls.io/r/juusechec/jwt-beego?branch=master)

You can read Docs in [english](README-en_US.md).

# jwt-beego
Una implementación simple de dgrijalva/jwt-go para beego.

Pasos para implementar:

1) Crear una llave RSA con los comandos de ***generar_key.log***

2) Generar una ruta con action POST en el controlador especificado, en este caso la ruta elegida fue /user/getToken:

```go
// ./controllers/user.go
package controllers

import (
	...

	"github.com/juusechec/jwt-beego"
)

...

// @Title getToken
// @Description Get token from user session
// @Param	username		query 	string	true		"The username for get token"
// @Param	password		query 	string	true		"The password for get token"
// @Success 200 {string} Obtain Token
// @router /getToken [post]
func (u *UserController) GetToken() {
	username := u.Ctx.Input.Query("username")
	password := u.Ctx.Input.Query("password")

	tokenString := ""
	if username == "admin" && password == "mipassword" {
		et := jwtbeego.EasyToken{
			Username: username,
			Expires:  time.Now().Unix() + 3600, //Segundos
		}
		tokenString, _ = et.GetToken()
	}

	u.Data["json"] = "{'tokenString': '" + tokenString + "'}"
	u.ServeJSON()
	return
}

...
```

3) Agregar la validación del token en cada controlador que se necesite. Esto se hace a través de la función ***Prepare***.

```go
// ./controllers/my_restricted_controller.go
package controllers

import (
	...

	"github.com/juusechec/jwt-beego"
)

func (c *TipoCancelacionSemestreController) Prepare() {
	tokenString := c.Ctx.Input.Query("tokenString")
	// O puede ser leído de una cabecera HEADER!!
	// tokenString := c.Ctx.Request.Header.Get("X-JWTtoken")

	et := jwtbeego.EasyToken{}
	valido, _, _ := et.ValidateToken(tokenString)
	if !valido {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Permission Deny"
		c.ServeJSON()
	}
	return
}

...
```

Con esto ya estaría terminado, pero si se quiere hacer que todos los controladores tengan la misma validación se puede hacer:

1) Configurar un nuevo paquete.

```go
//./myBeego/controller.go

//Se crea un espacio de nombres llamado myBeego
package myBeego

//Se agrega la biblioteca de beego
import (
	...
	"github.com/astaxie/beego"
)

//Se genera un tipo Controller que hereda de beego.Controller
type Controller struct {
	beego.Controller
}

//Se redefine lo que hace la función Prepare
//* es un apuntador al igual que en C
//& hace referencia a la dirección de memoria
//La iniciación de una variable o funcion con * se traduce en que almacena
//u := 10 //var z *int  //z = &u //fmt.Println(z)//0x1040e0f8
//var s *string //var r **string = &s //fmt.Println(r)//0x1040a120
func (c *Controller) Prepare() {
	//Lo que quieras hacer en todos los controladores
	tokenString := c.Ctx.Input.Query("tokenString")
	// O puede ser leído de una cabecera HEADER!!
	// tokenString := c.Ctx.Request.Header.Get("X-JWTtoken")

	et := jwtbeego.EasyToken{}
	valido, _, _ := et.ValidateToken(tokenString)
	if !valido {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Permission Deny"
		c.ServeJSON()
	}
	return
}

```

2) Configurar llamado del nuevo Controller en todos los controladores:

```go
//./controllers/miObjeto.go

package controllers

import (
	...
	"github.com/juusechec/jwt-beego-implementation/myBeego"
	"github.com/juusechec/jwt-beego"
)

type MiObjetoController struct {
	//beego.Controller
	myBeego.Controller
}

...
```

## Ejemplo de uso:
- https://github.com/juusechec/jwt-beego-implementation

Se basa en:
* https://github.com/someone1/gcp-jwt-go
* https://github.com/dgrijalva/jwt-go
