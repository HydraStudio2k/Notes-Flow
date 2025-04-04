package models

// models for working with the database

/*

	            User table:
----------------------------------------------
| ID | Name   | Username | Password   | Uuid |
----------------------------------------------
| 1  | Bob    | @user    | 12345      | uuid |
----------------------------------------------

* Username & uuid are unique fields that should not be repeated, the rest can be repeated.
* The username field is the author field in the table below. This is the relationship between the user and their note.

*/

type User struct {
	Id       int
	Name     string
	Username string
	Password string
	Uuid     string
}

/*

	Note table:
----------------------------------------
| ID | Author | Status | Note   | Uuid |
----------------------------------------
| 1  | @user  | public | #note1 | uuid |
----------------------------------------

* The author and uuid fields are unique fields and should not be repeated, other fields may be repeated.
* The status field has only 2 states: public and private.

*/

type Note struct {
	Id     int
	Author string
	Status string
	Note   string
	Uuid   string
}
