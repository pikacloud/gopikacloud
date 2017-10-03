package gopikacloud

// User defines a pikacloud user
type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

// Me returns current user
func (client *Client) Me() (User, error) {
	user := User{}
	if err := client.Get("me/", &user); err != nil {
		return User{}, err
	}
	return user, nil
}
