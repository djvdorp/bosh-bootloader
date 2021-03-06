package manifests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestManifests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "boshinit/manifests")
}

const (
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAy/NxInVJGqGATgD+JNGw94uShbYG9ZVs5JT0JxWbNih69o3P
QvRnVgYpV7JPj6PDvSq9MJfgBL8DfBPRyZz7u1NMVeuLGb84Lc//ZOpxNH7ItWFC
7AgSzhMdYFWwc2a5wNKTnz/OGejvFmVCpEj+OeeSsFnqabgK1DHj+cVCnmHpV+qx
TZggko8Z3MYrJyHj3oW3TkU5KWCta4TlXMIkIF/XaoFRJjXlTcKtCct371iV4MTP
beX1VxWHmKrJBngUGAIPptUJ4jmshkci+gkrWkD/l0Oka/TXVVOC7ONA8kijVd2k
q2s5BONqxOsHpeGCspmf/AwGitRhmsuUiKaVawIDAQABAoIBAQCXqC2/ftFegStk
Va6l+FuMp/fLf+DK61Mu4mhvS8y/x/gvsKGnWxAgUrKZaJlq6U7bMpW+NaE6RkO0
FxTPw/EYicdguKcV/TAsONfj1cVaUNC0t19JHnPFU0dr5Cwyk/ean7twFsUOVTy8
PY4Bldmde0qD4kjCVj2PLo4ko/92v2ghEZ09U8O0QjupBWe5CH6BjqcwC71Zey3k
lrHYy2bC/XM0MEUO4ruHTqwWCyNEfVgUMWK6PcrG5gw626UdhzoAxmAq7mQdvytx
DnwVu785ywKCKcfFZbjK2e6WJQNrcdcgDA6XKYxpahr5e0dBsrHKGAt5j25U9C9U
iuDVZO9pAoGBAMy7W8ZXUa94Qrn4daGrecH4osBZZKbWc4WemQQes1irovRbavUl
wJUV9MAFws6SB2lm6W4TFENYbXWxDQD+qzxfOOokfwQR7Cm4tJaIDSWtcGbQ3qpL
0CNc7Jgj1uYRqDOHzxH+t+9ylPjJAH5nrO/IQw6G5tgKUo16cp5jJTVNAoGBAP8G
BWkjCXP3Ln1zJXvesqV06uNA8q1O2jn+RB028vNIdqbcMqwgcLhO6BLjDxGZSI4+
f+R5dTLIHp8dPEtSFSjUc68U0PConodbNRGW64RgmJ42Bt+EgbwSpz084aukRF4f
nv5TVWlGNQSkhM5A1mIL+Yb9ft0GL4gm6a11LjmXAoGBALp+lVw4oIVZ5Fap2OW7
YT9rsT3McJ51zWfTkAmrua21M8yqFeVYTXTBOmFNHiaz6SJ/h3Yo1RV/0L4b2P+l
/PASwbKwKi/X76wVwBM7vdYrhq9x5tN0GhaMiE0SoKiVPwYp2VvID1+EneV+m/+J
i4QlhcgO5Ou5g3ezKgPTxsolAoGAGLatqY2iqeFHWRLijAl9yHj3FkTB/7eHgF06
npYnnxnjnbHzGykdo07KQKJOYIc5N8eovyxiBiTMiuDbafUvZcCI4Wuj/95nVnip
QOYeNrrhr1tO+TQvGlm4aT/QIsENew2Xa9AJk+Ug2C1VTWONIv1EFifUtniV/JzA
wnV7oEsCgYEAoxhOzmI6SOn5iGt3PfWchCcL+gkqRmBOptiydkdO+/KNdlXoNl7P
1zDpDTuE1kucxFJSratwD74gyVrwxNTsx/W94mYAzao+wA3EgozczQvw4Dfwj/+k
tmdF1gd0MYOQk/q0ssdNLLX3yl4GBGpsJlcOpz7qVHFB6yOOmggGsOU=
-----END RSA PRIVATE KEY-----`

	certificate = `-----BEGIN CERTIFICATE-----
MIIDLDCCAhSgAwIBAgICBNIwDQYJKoZIhvcNAQELBQAwPDEMMAoGA1UEBhMDVVNB
MRYwFAYDVQQKEw1DbG91ZCBGb3VuZHJ5MRQwEgYDVQQDEws1Mi4wLjExMi4xMjAe
Fw0xNjAzMTAwMDUwNTdaFw0xODAzMTAwMDUwNTdaMDwxDDAKBgNVBAYTA1VTQTEW
MBQGA1UEChMNQ2xvdWQgRm91bmRyeTEUMBIGA1UEAxMLNTIuMC4xMTIuMTIwggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDL83EidUkaoYBOAP4k0bD3i5KF
tgb1lWzklPQnFZs2KHr2jc9C9GdWBilXsk+Po8O9Kr0wl+AEvwN8E9HJnPu7U0xV
64sZvzgtz/9k6nE0fsi1YULsCBLOEx1gVbBzZrnA0pOfP84Z6O8WZUKkSP4555Kw
WeppuArUMeP5xUKeYelX6rFNmCCSjxncxisnIePehbdORTkpYK1rhOVcwiQgX9dq
gVEmNeVNwq0Jy3fvWJXgxM9t5fVXFYeYqskGeBQYAg+m1QniOayGRyL6CStaQP+X
Q6Rr9NdVU4Ls40DySKNV3aSrazkE42rE6wel4YKymZ/8DAaK1GGay5SIppVrAgMB
AAGjODA2MA4GA1UdDwEB/wQEAwIFIDATBgNVHSUEDDAKBggrBgEFBQcDATAPBgNV
HREECDAGhwQ0AHAMMA0GCSqGSIb3DQEBCwUAA4IBAQBdnVFV+yNJrHuuy+n+lLuM
wGRHOroks/iKIbVTISfgnA9jG4iWzijggyvWiA9YR3DbzPSbMfPmAHdEVoxtPTrW
OeUGMY9DoLEO1O7qxib75JS0RAxaAOa/x+w5BPNF3/wWxf2gQ60rNzjYJCEAKnJD
RkE1ITlQyATKr3WTz13DOWtsQ48I5Cb84ID6vHkTstABp5i2/ppEcP9aNAYgMu5I
j7UZfBzqZeNR1GtNNRTcXwEORY4rUR4dBCoOWOgvlcDv3BT0nGz1p/7nwvMDlkQ/
NNsrYQTmjEpgtOVxy3ifqad3mxkiEkXc67uWVoJY0jJT+FUemJP0cZTVTif3jncA
-----END CERTIFICATE-----`
)
