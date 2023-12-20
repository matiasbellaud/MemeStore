package handler

import (
	"errors"
    "regexp"
)

// set the news errors case used by the foloming functions
var (
	IncorectDescription = errors.New("Description do not match whith the rules")
	IncorectMail = errors.New("Mail do not match whith a corectly made mail adress")
	IncorectUsername = errors.New("Username do not match whith the rules")
	IncorectPassword = errors.New("Password do not match whith the rules")
	IncorctVerificationword = errors.New("Verification word do not match whith the rules")
)

func CorectInformation(description, mail, username, password, verificaztionword string) error {
	/*
	func Corectinformation verify if the information put by the user in the form are in the same form as we need, a valid mail adress, 
	a good password...
	--------------------------------------------------------------
	input : description, mail  username, password, verificationword as string and also the user information
	output : error case
	--------------------------------------------------------------
	step by step the function see if the user information match whith the regexp gived to him, especialy for mail.
	error case from the regexp.Matching handel
	if the information dosen't match whith regexp, the function return the error in the information

	after the func pass all the test it return a nil error, because the information dosent have one

	! the password dosen't use regexp rule but an other function write down there !
	*/
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
	/*
	the function GoodPassword verify if the password put are folowing our rules for password
	--------------------------------------------
	input : the password as string
	output : the boolean representing if the password folowing the rules, ans also error case
	--------------------------------------------
	at first, the programe verify if the password length is superior or equal to 12, if the passwors lenth is lower than 12 it return 
	the password error

	next the function set the variables used to count the password lowerkase, upercase, digit, special caractéres
	the function walk acrose the password letter by letter and if the caracére is in one of the 4 family of letter it add 1 to the count of 
	it family, the fucntion use regexp to determine if letter is in on of the 4 family.
	if the letter dosen't bee in the 4 family the program return password error and false

	once the letter see one by one, whe return true if the password have 1 or + upercase and lowercase and special caractére and 4 or + digit,
	the program return nil error.
	*/
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
			continue
		}
		matched, err = regexp.MatchString(`[A-Z]`, string(password[caractere]))
		if err != nil {
			return false, err
		}
		if matched {
			upercaseLetter += 1
			continue
		}
		matched, err = regexp.MatchString(`[0-9]`, string(password[caractere]))
		if err != nil {
			return false, err
		}
		if matched {
			digit += 1
			continue
		}
		matched =  byte(33) <=password[caractere] && password[caractere]<= byte(47) || 
		byte(58) <=password[caractere] && password[caractere]<= byte(64) || 
		byte(91) <=password[caractere] && password[caractere]<= byte(96) || 
		byte(123) <=password[caractere] && password[caractere]<= byte(126)
		if matched {
			specialCaractere += 1
			continue
		}
		return false, nil
	}
	return lowercaseLetter > 0 && upercaseLetter > 0 && digit >= 4 && specialCaractere >= 0 , nil
}