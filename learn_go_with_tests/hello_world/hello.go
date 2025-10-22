package main

import (
	"fmt"
)

type language string

const (
	ENGLISH language = "en"
	HINDI language= "hi"
	FRENCH language= "fr"
	SPANISH language = "sp"
	GERMAN language= "de"
	GUJARATI  language= "gu"
	JAPANESE language= "ja"
	KOREAN language= "ko"
)

const (
	englishPrefix  = "Hello"        // en
	hindiPrefeix    = "नमस्ते"       // hi
	frenchPrefic   = "Bonjour"      // fr
	spanishPrefix  = "Hola"         // sp
	germanPrefix   = "Hallo"        // de
	gujaratiPrefix = "નમસ્તે"       // gu
	japanesePrefix = "こんにちは"     // ja
	koreanPrefix   = "안녕하세요"      // ko
)

func Hello(name string,lang language)string{
	var prefix string
	switch lang{
		case ENGLISH:
			prefix =englishPrefix		
		case HINDI:
			prefix = hindiPrefeix
		case FRENCH:
			prefix = frenchPrefic
		case SPANISH:
			prefix = spanishPrefix
		case GERMAN:
			prefix = germanPrefix
		case GUJARATI:
			prefix = gujaratiPrefix
		case JAPANESE:
			prefix = japanesePrefix
		case KOREAN:
			prefix = koreanPrefix
		default:
			prefix = englishPrefix
	}
	if name == "" || name == " "{
		name = "World"
	}
	return prefix+" "+name
}

func main(){
	fmt.Println(Hello("Dhwanil",ENGLISH))
}