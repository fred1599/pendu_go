package main

import (
    "os"
    "fmt"
    "strings"
    "unicode/utf8"
    "bufio"
)

func word_to_secret(word string) string {
    secret_word := strings.Repeat("*", utf8.RuneCountInString(word))
    return secret_word
}

func replace_char(secret string, word string, char byte) string {
    length := utf8.RuneCountInString(word)
    byte_word := []byte(word)
    byte_secret := []byte(secret)
    for i:= 0; i < length; i++ {
        if byte_word[i] == char {
            byte_secret[i] = char
        }
    }
    return string(byte_secret)
}

func main() {
        word := "test"
        secret := word_to_secret(word)
        for word != secret {
            fmt.Println(secret)
            fmt.Println("Entrer votre lettre: ")
            scanner := bufio.NewScanner(os.Stdin)
            if scanner.Scan() {
                c := scanner.Bytes()
                if len(c) != 0 {
                    secret = replace_char(secret, word, c[0])
                }
            }
        }
        fmt.Println(word)
}
