package main

import (
    "crypto/rand"
    "flag"
    "fmt"
    "math"
    "math/big"
    "os"
    "strings"

    "github.com/atotto/clipboard"
)

const (
    lower   = "abcdefghijklmnopqrstuvwxyz"
    upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    digits  = "0123456789"
    symbols = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

func main() {
    length := flag.Int("l", 16, "Şifre uzunluğu")
    count := flag.Int("c", 1, "Kaç tane şifre üretilecek")
    noUpper := flag.Bool("no-upper", false, "Büyük harf kullanma")
    noDigits := flag.Bool("no-digits", false, "Rakam kullanma")
    noSymbols := flag.Bool("no-symbols", false, "Özel karakter kullanma")
    copy := flag.Bool("copy", false, "İlk şifreyi panoya kopyala")
    flag.Parse()

    charset := lower
    if !*noUpper {
        charset += upper
    }
    if !*noDigits {
        charset += digits
    }
    if !*noSymbols {
        charset += symbols
    }

    if len(charset) == 0 {
        fmt.Println("Hata: En az bir karakter seti seçmelisin!")
        os.Exit(1)
    }

    fmt.Printf("passgen-go – %d karakterli %d şifre üretiliyor...\n\n", *length, *count)

    for i := 0; i < *count; i++ {
        password := generatePassword(*length, charset)
        strength := calculateStrength(password)

        fmt.Printf("%2d: %s  →  %s\n", i+1, password, strength)

        if *copy && i == 0 {
            clipboard.WriteAll(password)
            fmt.Println("   İlk şifre panoya kopyalandı!")
        }
    }
}

func generatePassword(length int, charset string) string {
    result := make([]byte, length)
    max := big.NewInt(int64(len(charset)))

    for i := range result {
        n, _ := rand.Int(rand.Reader, max)
        result[i] = charset[n.Int64()]
    }
    return string(result)
}

func calculateStrength(password string) string {
    var score float64
    if len(password) >= 12 {
        score += 2
    }
    if strings.ContainsAny(password, upper) {
        score += 1
    }
    if strings.ContainsAny(password, digits) {
        score += 1
    }
    if strings.ContainsAny(password, symbols) {
        score += 2
    }
    entropy := float64(len(password)) * math.Log2(float64(len(password)))
    if entropy > 80 {
        score += 2
    }

    switch {
    case score >= 7:
        return "ÇOK GÜÇLÜ"
    case score >= 5:
        return "GÜÇLÜ"
    case score >= 3:
        return "ORTA"
    default:
        return "ZAYIF"
    }
}