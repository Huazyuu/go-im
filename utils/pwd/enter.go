package pwd

import "golang.org/x/crypto/bcrypt"

func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return err.Error()
	}
	return string(hash)
}

func VerifyPwd(hash string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}
