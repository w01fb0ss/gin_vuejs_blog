/*
 * Copyright (c) 2019.
 * Author:Alex	xiaoshenlong1910@gmail.com

   .--,       .--,
 ( (  \.---./  ) )
  '.__/o   o\__.'
     {=  ^  =}
      >  -  <
     /       \
    //       \\
   //|   .   |\\
   "'\       /'"_.-~^`'-.
      \  _  /--'         `
    ___)( )(___
   (((__) (__)))    高山仰止,景行行止.虽不能至,心向往之。

*/

package models

type User struct {
	BaseModel
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUser(username, password string) bool {
	var user User
	db.Select("id").Where(User{Username: username, Password: password}).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func IsExist(username, email, password string) bool {
	var user User
	db.Select("id").Where(User{Username: username, Email:email, Password: password}).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

// func AddUser(username, email, password string) error {
//
// }
