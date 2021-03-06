package auth

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA60Nll48e9AjnDMtYwFxC6FpGTgB+05YczWNo2vHAQoQ2CK8o
cIdmnX713aqZgMymGueyETE063dNIwHNd4hI8eNb6yT5036MUIxqFfCeoiY9L5v7
dUWt2I2Y2zO0GyFNbKmsSUgY/uFc4Pehc/rCdxa+yK4XINj3GwfZMTzZj9mbHmG1
aUrl2BHQ+atWJPZ5RGfJ894MOuVig+TczP9q2oq3FwihPSymlSoXMtU0swXP9oPW
ihC4Jl6fooeoktFaB4aKMtbwYoI2bFI1Z/HtcPvpKTqA2N4I06uUBAd7fYQ6S67e
sfCWuz41GRsmt4/ow+n4VDbEL9+fvTbuvBUScwIDAQABAoIBAQCIZMOCUZRlcL/b
OA0VQKf9o1C9KiZdLts17Bs/TAblyVa2RC5Sxv0j9BZA1AY4mOz4Djm7I+cFWF4d
HI9tbFK2AdUph39bMTPSAwYCAEVHxpmUf38SZzo5oSXsd5ZEYFNN2zSqK6VqJ2e8
eXReFBQoDMTA/cUk4wSn5C30EM4lSavIahsWjjhhQ3TJEHq+Db8V86mkTEYY0lVG
QHWmVeXgswaeSwi71Uek1fYCqJX9Y6B7BJggOTYvy0VrdFY7Gsv2W38NbzAdfNHJ
ap3Zvbfs8l9OCoXLMn3OpzYcJAP9r095REhEETZ0daEgSAWmkdEeMyrVbe08Hn8V
3atSQe1xAoGBAPV6nPqaTZvxBiAZPn+t7dU2hBkzSVQzy16iVhYV3JEy1tSybCb3
1s9bRRIgGyFy6GVcbxP8y37sZOH1N6yESapew1v9TMJ4ns1AXWKzOBh1sKpeq082
rreNPmTs4FrWH5Rsg9fI8yhSNmT7PtsPhh0n6mc1OjMHRTAeFpetlUPJAoGBAPVY
soxF+cIwHFbq8OKTaObeG217qUlW9gWFahHbHKZq2SDc8wzuj10+2zJ+p444jQ+7
htQoo+EOd4okV8FchpnXZ5XGdb8x2QM947NhPdpvSCvSXpo9GMXkB1Fc5qWH2Fck
7OiFRmyggicWfQty1zsVwP5NSjVIWm7WavPRLypbAoGBAPEgYshK/4b5RxIKmfii
6WaxKSrz1MjQI6kufq8tBD7gGGRPaqsQccG3Hm7CeMPvclX2aOnSPDXNkP7fd7OP
MikW8oEOm/K+twZYfiKwdp6YJFQFr/KcfRyxsiNDBtJTWCvcN9mjey4VK3cf0Esv
wYYsKjT89NEhCBZOx9Rv3pn5AoGAfVRP8QQV77Kv516gqr1uivDOxgCzm3lGt89N
Pc3WNtTCEd2McwCyT4m6Y3L7mNUeifjbiSkxfdfsDK8//NCNGa1FhucYYxDNsLA+
zvVPqXPLREd0eAnZy80OR+vTDp/0TZ/ZPXZHYVok4l7EhcKmW9JjAgTtcYaCW7hA
bkhfdZ0CgYAArK17pmGcfzxOPfGjo409xEHc/LJRwxbL95xj9FStUrU2eWXw1jrP
ViOsqXovpbpG9aNPvQxYFbV9xQjCwpYqFlDB+UH1g7SlCOtXfYjzs6efyizQCOBw
/ASilbauT62xLbRvYY0F7M399gwAJocHJrAqOdf3PoCBLjRONno2QQ==
-----END RSA PRIVATE KEY-----`

func TestJWTTokenGen_GenerateToken(t1 *testing.T) {

	pem, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t1.Fatal(err)
	}
	gen := NewJWTTokenGen("coolcar/auth", pem)

	token, err := gen.GenerateToken("hahhhh", 7200)

	if err != nil {
		t1.Fatal(err)
	}
	t1.Log(token)

}
