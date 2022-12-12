package clients

import (
	"net"
)

type Client struct {
  Name string
  Conection net.Conn
}

// Tells the index of a user in the user array
func FindIndex(userName string, userArr *[]Client) int {
  for i := 0; i < len(*userArr); i++ {
    if (*userArr)[i].Name == userName {
      return i
    }
  }
  return -1
}

// Deletes a user from the user array and close it connection
func DeleteUser(user Client, userArr *[]Client) {
  findIndex := FindIndex(user.Name, userArr)
  if findIndex == -1 {
    return
  }
  slice1 := (*userArr)[0:findIndex]
  slice2 := (*userArr)[(findIndex + 1):]
  *userArr = append(slice1, slice2...)
  user.Conection.Close()
}

// Adds a user to the chat
func AddUser(name string, conn net.Conn, userArr *[]Client) {
  *userArr = append(*userArr, Client{Name: name, Conection: conn})
}