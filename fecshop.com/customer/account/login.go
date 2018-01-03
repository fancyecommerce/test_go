package account


func Login(email, password string) string {
    msg := "email:" + email + ", password: " + password + " ";
    msg += "login success"
    return msg
}