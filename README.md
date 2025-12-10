# PassGen-Go – Ultra Hızlı Şifre Üretici

Go ile yazılmış, 1M+ şifre/saniye üreten CLI tool. Güç ölçer ve clipboard desteği var.

## Kurulum
go install github.com/zaferertasglobal-hub/passgen-go@latest


## Kullanım
passgen-go -l 16 -c 3 -copy  # 3 tane 16 karakterli şifre, 

## Özellikler
- Parametreler: -l (uzunluk), -c (adet), -no-upper/digits/symbols
- Güç hesaplama: Harf/rakam/sembol/entropy bazlı
- Teknolojiler: Go 1.25, crypto/rand, clipboard

MIT Lisans

