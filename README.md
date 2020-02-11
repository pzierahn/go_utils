# GoGo

## Deps

```shell
go get github.com/SSSaaS/sssa-golang
go get github.com/deckarep/golang-set
```

## Shamir's Secret Sharing

https://github.com/SSSaaS/sssa-golang

```go
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {

	b := make([]rune, n)

	for inx := range b {
		b[ inx ] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func sssaas() {

	text := RandStringRunes(1024 * 4)
	fmt.Println("msg =", text)

	parts, _ := sssa.Create(3, 9, text)

	for inx, part := range parts {
		fmt.Println("inx=" + strconv.Itoa(inx) + " part=" + part)
	}

	msg, _ := sssa.Combine(parts[:3])

	fmt.Println("msg =", msg)
}

```