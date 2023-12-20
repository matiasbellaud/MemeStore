package handler

import (
	"errors"
    "regexp"
)

var (
	IncorectDescription = errors.New("Description do not match whith the rules")
	IncorectMail = errors.New("Mail do not match whith a corectly made mail adress")
	IncorectUsername = errors.New("Username do not match whith the rules")
	IncorectPassword = errors.New("Password do not match whith the rules")
	IncorctVerificationword = errors.New("Verification word do not match whith the rules")
)

func CorectInformation(description, mail, username, password, verificaztionword string) error {
	matched, err := regexp.MatchString(`.`, description)
	if err != nil {
		return err
	}
	if matched {
		matched, err := regexp.MatchString(`.*@.*\..*`, mail)
		if err != nil {
			return err
		}
		if matched {
			matched, err := regexp.MatchString(`.`, username)
			if err != nil {
				return err
			}
			if matched {
				matched, err := GoodPassword(password)
				if err != nil {
					return err
				}
				if matched {
					matched, err := regexp.MatchString(`.`, verificaztionword)
					if err != nil {
						return err
					}
					if matched {
						return nil
					}
					return IncorctVerificationword
				}
				return IncorectPassword
			}
			return IncorectUsername
		}
		return IncorectMail
	}
	return IncorectDescription
}

func GoodPassword(password string) (goodPassword bool, err error) {
	if len(password) < 12 {
		return false, nil
	}
	lowercaseLetter := 0
	upercaseLetter := 0
	digit := 0
	specialCaractere := 0
	
	for caractere := range(password) {
		matched, err := regexp.MatchString(`[a-z]`, string(password[caractere]))
		if err != nil {
			return false, err
		}
		if matched {
			lowercaseLetter += 1
		}
		matched, err = regexp.MatchString(`[A-Z]`, string(password[caractere]))
		if err != nil {
			return false, err
		}
		if matched {
			upercaseLetter += 1
		}
		matched, err = regexp.MatchString(`[0-9]`, string(password[caractere]))
		if err != nil {
			return false, err
		}
		if matched {
			digit += 1
		}
		matched =  byte(33) <=password[caractere] && password[caractere]<= byte(47) || 
		byte(58) <=password[caractere] && password[caractere]<= byte(64) || 
		byte(91) <=password[caractere] && password[caractere]<= byte(96) || 
		byte(123) <=password[caractere] && password[caractere]<= byte(126)
		if matched {
			specialCaractere += 1
		}
	}
	if lowercaseLetter > 0 && upercaseLetter > 0 && digit >= 4 && specialCaractere >= 0{
		return true, nil
	} else {
		return false, nil
	}
}