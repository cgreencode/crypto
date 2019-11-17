// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testdata

var PEMBytes = map[string][]byte{
	"dsa": []byte(`-----BEGIN DSA PRIVATE KEY-----
MIIBuwIBAAKBgQD6PDSEyXiI9jfNs97WuM46MSDCYlOqWw80ajN16AohtBncs1YB
lHk//dQOvCYOsYaE+gNix2jtoRjwXhDsc25/IqQbU1ahb7mB8/rsaILRGIbA5WH3
EgFtJmXFovDz3if6F6TzvhFpHgJRmLYVR8cqsezL3hEZOvvs2iH7MorkxwIVAJHD
nD82+lxh2fb4PMsIiaXudAsBAoGAQRf7Q/iaPRn43ZquUhd6WwvirqUj+tkIu6eV
2nZWYmXLlqFQKEy4Tejl7Wkyzr2OSYvbXLzo7TNxLKoWor6ips0phYPPMyXld14r
juhT24CrhOzuLMhDduMDi032wDIZG4Y+K7ElU8Oufn8Sj5Wge8r6ANmmVgmFfynr
FhdYCngCgYEA3ucGJ93/Mx4q4eKRDxcWD3QzWyqpbRVRRV1Vmih9Ha/qC994nJFz
DQIdjxDIT2Rk2AGzMqFEB68Zc3O+Wcsmz5eWWzEwFxaTwOGWTyDqsDRLm3fD+QYj
nOwuxb0Kce+gWI8voWcqC9cyRm09jGzu2Ab3Bhtpg8JJ8L7gS3MRZK4CFEx4UAfY
Fmsr0W6fHB9nhS4/UXM8
-----END DSA PRIVATE KEY-----
`),
	"ecdsa": []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEINGWx0zo6fhJ/0EAfrPzVFyFC9s18lBt3cRoEDhS3ARooAoGCCqGSM49
AwEHoUQDQgAEi9Hdw6KvZcWxfg2IDhA7UkpDtzzt6ZqJXSsFdLd+Kx4S3Sx4cVO+
6/ZOXRnPmNAlLUqjShUsUBBngG0u2fqEqA==
-----END EC PRIVATE KEY-----
`),
	"ecdsap256": []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIAPCE25zK0PQSnsgVcEbM1mbKTASH4pqb5QJajplDwDZoAoGCCqGSM49
AwEHoUQDQgAEWy8TxGcIHRh5XGpO4dFVfDjeNY+VkgubQrf/eyFJZHxAn1SKraXU
qJUjTKj1z622OxYtJ5P7s9CfAEVsTzLCzg==
-----END EC PRIVATE KEY-----
`),
	"ecdsap384": []byte(`-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDBWfSnMuNKq8J9rQLzzEkx3KAoEohSXqhE/4CdjEYtoU2i22HW80DDS
qQhYNHRAduygBwYFK4EEACKhZANiAAQWaDMAd0HUd8ZiXCX7mYDDnC54gwH/nG43
VhCUEYmF7HMZm/B9Yn3GjFk3qYEDEvuF/52+NvUKBKKaLbh32AWxMv0ibcoba4cz
hL9+hWYhUD9XIUlzMWiZ2y6eBE9PdRI=
-----END EC PRIVATE KEY-----
`),
	"ecdsap521": []byte(`-----BEGIN EC PRIVATE KEY-----
MIHcAgEBBEIBrkYpQcy8KTVHNiAkjlFZwee90224Bu6wz94R4OBo+Ts0eoAQG7SF
iaygEDMUbx6kTgXTBcKZ0jrWPKakayNZ/kigBwYFK4EEACOhgYkDgYYABADFuvLV
UoaCDGHcw5uNfdRIsvaLKuWSpLsl48eWGZAwdNG432GDVKduO+pceuE+8XzcyJb+
uMv+D2b11Q/LQUcHJwE6fqbm8m3EtDKPsoKs0u/XUJb0JsH4J8lkZzbUTjvGYamn
FFlRjzoB3Oxu8UQgb+MWPedtH9XYBbg9biz4jJLkXQ==
-----END EC PRIVATE KEY-----
`),
	"rsa": []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC8A6FGHDiWCSREAXCq6yBfNVr0xCVG2CzvktFNRpue+RXrGs/2
a6ySEJQb3IYquw7HlJgu6fg3WIWhOmHCjfpG0PrL4CRwbqQ2LaPPXhJErWYejcD8
Di00cF3677+G10KMZk9RXbmHtuBFZT98wxg8j+ZsBMqGM1+7yrWUvynswQIDAQAB
AoGAJMCk5vqfSRzyXOTXLGIYCuR4Kj6pdsbNSeuuRGfYBeR1F2c/XdFAg7D/8s5R
38p/Ih52/Ty5S8BfJtwtvgVY9ecf/JlU/rl/QzhG8/8KC0NG7KsyXklbQ7gJT8UT
Ojmw5QpMk+rKv17ipDVkQQmPaj+gJXYNAHqImke5mm/K/h0CQQDciPmviQ+DOhOq
2ZBqUfH8oXHgFmp7/6pXw80DpMIxgV3CwkxxIVx6a8lVH9bT/AFySJ6vXq4zTuV9
6QmZcZzDAkEA2j/UXJPIs1fQ8z/6sONOkU/BjtoePFIWJlRxdN35cZjXnBraX5UR
fFHkePv4YwqmXNqrBOvSu+w2WdSDci+IKwJAcsPRc/jWmsrJW1q3Ha0hSf/WG/Bu
X7MPuXaKpP/DkzGoUmb8ks7yqj6XWnYkPNLjCc8izU5vRwIiyWBRf4mxMwJBAILa
NDvRS0rjwt6lJGv7zPZoqDc65VfrK2aNyHx2PgFyzwrEOtuF57bu7pnvEIxpLTeM
z26i6XVMeYXAWZMTloMCQBbpGgEERQpeUknLBqUHhg/wXF6+lFA+vEGnkY+Dwab2
KCXFGd+SQ5GdUcEMe9isUH6DYj/6/yCDoFrXXmpQb+M=
-----END RSA PRIVATE KEY-----
`),
	"pkcs8": []byte(`-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCitzS2KiRQTccf
VApb0mbPpo1lt29JjeLBYAehXHWfQ+w8sXpd8e04n/020spx1R94yg+v0NjXyh2R
NFXNBYdhNei33VJxUeKNlExaecvW2yxfuZqka+ZxT1aI8zrAsjh3Rwc6wayAJS4R
wZuzlDv4jZitWqwD+mb/22Zwq/WSs4YX5dUHDklfdWSVnoBfue8K/00n8f5yMTdJ
vFF0qAJwf9spPEHla0lYcozJk64CO5lRkqfLor4UnsXXOiA7aRIoaUSKa+rlhiqt
1EMGYiBjblPt4SwMelGGU2UfywPb4d85gpQ/s8SBARbpPxNVs2IbHDMwj70P3uZc
74M3c4VJAgMBAAECggEAFIzY3mziGzZHgMBncoNXMsCRORh6uKpvygZr0EhSHqRA
cMXlc3n7gNxL6aGjqc7F48Z5RrY0vMQtCcq3T2Z0W6WoV5hfMiqqV0E0h3S8ds1F
hG13h26NMyBXCILXl8Cqev4Afr45IBISCHIQTRTaoiCX+MTr1rDIU2YNQQumvzkz
fMw2XiFTFTgxAtJUAgKoTqLtm7/T+az7TKw+Hesgbx7yaJoMh9DWGBh4Y61DnIDA
fcxJboAfxxnFiXvdBVmzo72pCsRXrWOsjW6WxQmCKuXHvyB1FZTmMaEFNCGSJDa6
U+OCzA3m65loAZAE7ffFHhYgssz/h9TBaOjKO0BX1QKBgQDZiCBvu+bFh9pEodcS
VxaI+ATlsYcmGdLtnZw5pxuEdr60iNWhpEcV6lGkbdiv5aL43QaGFDLagqeHI77b
+ITFbPPdCiYNaqlk6wyiXv4pdN7V683EDmGWSQlPeC9IhUilt2c+fChK2EB/XlkO
q8c3Vk1MsC6JOxDXNgJxylNpswKBgQC/fYBTb9iD+uM2n3SzJlct/ZlPaONKnNDR
pbTOdxBFHsu2VkfY858tfnEPkmSRX0yKmjHni6e8/qIzfzLwWBY4NmxhNZE5v+qJ
qZF26ULFdrZB4oWXAOliy/1S473OpQnp2MZp2asd0LPcg/BNaMuQrz44hxHb76R7
qWD0ebIfEwKBgQCRCIiP1pjbVGN7ZOgPS080DSC+wClahtcyI+ZYLglTvRQTLDQ7
LFtUykCav748MIADKuJBnM/3DiuCF5wV71EejDDfS/fo9BdyuKBY1brhixFTUX+E
Ww5Hc/SoLnpgALVZ/7jvWTpIBHykLxRziqYtR/YLzl+IkX/97P2ePoZ0rwKBgHNC
/7M5Z4JJyepfIMeVFHTCaT27TNTkf20x6Rs937U7TDN8y9JzEiU4LqXI4HAAhPoI
xnExRs4kF04YCnlRDE7Zs3Lv43J3ap1iTATfcymYwyv1RaQXEGQ/lUQHgYCZJtZz
fTrJoo5XyWu6nzJ5Gc8FLNaptr5ECSXGVm3Rsr2xAoGBAJWqEEQS/ejhO05QcPqh
y4cUdLr0269ILVsvic4Ot6zgfPIntXAK6IsHGKcg57kYm6W9k1CmmlA4ENGryJnR
vxyyqA9eyTFc1CQNuc2frKFA9It49JzjXahKc0aDHEHmTR787Tmk1LbuT0/gm9kA
L4INU6g+WqF0fatJxd+IJPrp
-----END PRIVATE KEY-----
`),
	"ed25519": []byte(`-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACA+3f7hS7g5UWwXOGVTrMfhmxyrjqz7Sxxbx7I1j8DvvwAAAJhAFfkOQBX5
DgAAAAtzc2gtZWQyNTUxOQAAACA+3f7hS7g5UWwXOGVTrMfhmxyrjqz7Sxxbx7I1j8Dvvw
AAAEAaYmXltfW6nhRo3iWGglRB48lYq0z0Q3I3KyrdutEr6j7d/uFLuDlRbBc4ZVOsx+Gb
HKuOrPtLHFvHsjWPwO+/AAAAE2dhcnRvbm1AZ2FydG9ubS14cHMBAg==
-----END OPENSSH PRIVATE KEY-----
`),
	"rsa-openssh-format": []byte(`-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAlwAAAAdzc2gtcn
NhAAAAAwEAAQAAAIEAwa48yfWFi3uIdqzuf9X7C2Zxfea/Iaaw0zIwHudpF8U92WVIiC5l
oEuW1+OaVi3UWfIEjWMV1tHGysrHOwtwc34BPCJqJknUQO/KtDTBTJ4Pryhw1bWPC999Lz
a+yrCTdNQYBzoROXKExZgPFh9pTMi5wqpHDuOQ2qZFIEI3lT0AAAIQWL0H31i9B98AAAAH
c3NoLXJzYQAAAIEAwa48yfWFi3uIdqzuf9X7C2Zxfea/Iaaw0zIwHudpF8U92WVIiC5loE
uW1+OaVi3UWfIEjWMV1tHGysrHOwtwc34BPCJqJknUQO/KtDTBTJ4Pryhw1bWPC999Lza+
yrCTdNQYBzoROXKExZgPFh9pTMi5wqpHDuOQ2qZFIEI3lT0AAAADAQABAAAAgCThyTGsT4
IARDxVMhWl6eiB2ZrgFgWSeJm/NOqtppWgOebsIqPMMg4UVuVFsl422/lE3RkPhVkjGXgE
pWvZAdCnmLmApK8wK12vF334lZhZT7t3Z9EzJps88PWEHo7kguf285HcnUM7FlFeissJdk
kXly34y7/3X/a6Tclm+iABAAAAQE0xR/KxZ39slwfMv64Rz7WKk1PPskaryI29aHE3mKHk
pY2QA+P3QlrKxT/VWUMjHUbNNdYfJm48xu0SGNMRdKMAAABBAORh2NP/06JUV3J9W/2Hju
X1ViJuqqcQnJPVzpgSL826EC2xwOECTqoY8uvFpUdD7CtpksIxNVqRIhuNOlz0lqEAAABB
ANkaHTTaPojClO0dKJ/Zjs7pWOCGliebBYprQ/Y4r9QLBkC/XaWMS26gFIrjgC7D2Rv+rZ
wSD0v0RcmkITP1ZR0AAAAYcHF1ZXJuYUBMdWNreUh5ZHJvLmxvY2FsAQID
-----END OPENSSH PRIVATE KEY-----`),
	"user": []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEILYCAeq8f7V4vSSypRw7pxy8yz3V5W4qg8kSC3zJhqpQoAoGCCqGSM49
AwEHoUQDQgAEYcO2xNKiRUYOLEHM7VYAp57HNyKbOdYtHD83Z4hzNPVC4tM5mdGD
PLL8IEwvYu2wq+lpXfGQnNMbzYf9gspG0w==
-----END EC PRIVATE KEY-----
`),
	"ca": []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAvg9dQ9IRG59lYJb+GESfKWTch4yBpr7Ydw1jkK6vvtrx9jLo
5hkA8X6+ElRPRqTAZSlN5cBm6YCAcQIOsmXDUn6Oj1lVPQAoOjTBTvsjM3NjGhvv
52kHTY0nsMsBeY9q5DTtlzmlYkVUq2a6Htgf2mNi01dIw5fJ7uTTo8EbNf7O0i3u
c9a8P19HaZl5NKiWN4EIZkfB2WdXYRJCVBsGgQj3dE/GrEmH9QINq1A+GkNvK96u
vZm8H1jjmuqzHplWa7lFeXcx8FTVTbVb/iJrZ2Lc/JvIPitKZWhqbR59yrGjpwEp
Id7bo4WhO5L3OB0fSIJYvfu+o4WYnt4f3UzecwIDAQABAoIBABRD9yHgKErVuC2Q
bA+SYZY8VvdtF/X7q4EmQFORDNRA7EPgMc03JU6awRGbQ8i4kHs46EFzPoXvWcKz
AXYsO6N0Myc900Tp22A5d9NAHATEbPC/wdje7hRq1KyZONMJY9BphFv3nZbY5apR
Dc90JBFZP5RhXjTc3n9GjvqLAKfFEKVmPRCvqxCOZunw6XR+SgIQLJo36nsIsbhW
QUXIVaCI6cXMN8bRPm8EITdBNZu06Fpu4ZHm6VaxlXN9smERCDkgBSNXNWHKxmmA
c3Glo2DByUr2/JFBOrLEe9fkYgr24KNCQkHVcSaFxEcZvTggr7StjKISVHlCNEaB
7Q+kPoECgYEA3zE9FmvFGoQCU4g4Nl3dpQHs6kaAW8vJlrmq3xsireIuaJoa2HMe
wYdIvgCnK9DIjyxd5OWnE4jXtAEYPsyGD32B5rSLQrRO96lgb3f4bESCLUb3Bsn/
sdgeE3p1xZMA0B59htqCrvVgN9k8WxyevBxYl3/gSBm/p8OVH1RTW/ECgYEA2f9Z
95OLj0KQHQtxQXf+I3VjhCw3LkLW39QZOXVI0QrCJfqqP7uxsJXH9NYX0l0GFTcR
kRrlyoaSU1EGQosZh+n1MvplGBTkTSV47/bPsTzFpgK2NfEZuFm9RoWgltS+nYeH
Y2k4mnAN3PhReCMwuprmJz8GRLsO3Cs2s2YylKMCgYEA2UX+uO/q7jgqZ5UJW+ue
1H5+W0aMuFA3i7JtZEnvRaUVFqFGlwXin/WJ2+WY1++k/rPrJ+Rk9IBXtBUIvEGw
FC5TIfsKQsJyyWgqx/jbbtJ2g4s8+W/1qfTAuqeRNOg5d2DnRDs90wJuS4//0JaY
9HkHyVwkQyxFxhSA/AHEMJECgYA2MvyFR1O9bIk0D3I7GsA+xKLXa77Ua53MzIjw
9i4CezBGDQpjCiFli/fI8am+jY5DnAtsDknvjoG24UAzLy5L0mk6IXMdB6SzYYut
7ak5oahqW+Y9hxIj+XvLmtGQbphtxhJtLu35x75KoBpxSh6FZpmuTEccs31AVCYn
eFM/DQKBgQDOPUwbLKqVi6ddFGgrV9MrWw+SWsDa43bPuyvYppMM3oqesvyaX1Dt
qDvN7owaNxNM4OnfKcZr91z8YPVCFo4RbBif3DXRzjNNBlxEjHBtuMOikwvsmucN
vIrbeEpjTiUMTEAr6PoTiVHjsfS8WAM6MDlF5M+2PNswDsBpa2yLgA==
-----END RSA PRIVATE KEY-----
`),
}

var SSHCertificates = map[string][]byte{
	// The following are corresponding certificates for the private keys above, signed by the CA key
	// Generated by the following commands:
	//
	// 1. Assumes "rsa" key above in file named "rsa", write out the public key to "rsa.pub":
	//    ssh-keygen -y -f rsa > rsa.pub
	//
	// 2. Assumes "ca" key above in file named "ca", sign a cert for "rsa.pub":
	//    ssh-keygen -s ca -h -n host.example.com -V +500w -I host.example.com-key rsa.pub
	"rsa": []byte(`ssh-rsa-cert-v01@openssh.com AAAAHHNzaC1yc2EtY2VydC12MDFAb3BlbnNzaC5jb20AAAAgLjYqmmuTSEmjVhSfLQphBSTJMLwIZhRgmpn8FHKLiEIAAAADAQABAAAAgQC8A6FGHDiWCSREAXCq6yBfNVr0xCVG2CzvktFNRpue+RXrGs/2a6ySEJQb3IYquw7HlJgu6fg3WIWhOmHCjfpG0PrL4CRwbqQ2LaPPXhJErWYejcD8Di00cF3677+G10KMZk9RXbmHtuBFZT98wxg8j+ZsBMqGM1+7yrWUvynswQAAAAAAAAAAAAAAAgAAABRob3N0LmV4YW1wbGUuY29tLWtleQAAABQAAAAQaG9zdC5leGFtcGxlLmNvbQAAAABZHN8UAAAAAGsjIYUAAAAAAAAAAAAAAAAAAAEXAAAAB3NzaC1yc2EAAAADAQABAAABAQC+D11D0hEbn2Vglv4YRJ8pZNyHjIGmvth3DWOQrq++2vH2MujmGQDxfr4SVE9GpMBlKU3lwGbpgIBxAg6yZcNSfo6PWVU9ACg6NMFO+yMzc2MaG+/naQdNjSewywF5j2rkNO2XOaViRVSrZroe2B/aY2LTV0jDl8nu5NOjwRs1/s7SLe5z1rw/X0dpmXk0qJY3gQhmR8HZZ1dhEkJUGwaBCPd0T8asSYf1Ag2rUD4aQ28r3q69mbwfWOOa6rMemVZruUV5dzHwVNVNtVv+ImtnYtz8m8g+K0plaGptHn3KsaOnASkh3tujhaE7kvc4HR9Igli9+76jhZie3h/dTN5zAAABDwAAAAdzc2gtcnNhAAABALeDea+60H6xJGhktAyosHaSY7AYzLocaqd8hJQjEIDifBwzoTlnBmcK9CxGhKuaoJFThdCLdaevCeOSuquh8HTkf+2ebZZc/G5T+2thPvPqmcuEcmMosWo+SIjYhbP3S6KD49aLC1X0kz8IBQeauFvURhkZ5ZjhA1L4aQYt9NjL73nqOl8PplRui+Ov5w8b4ldul4zOvYAFrzfcP6wnnXk3c1Zzwwf5wynD5jakO8GpYKBuhM7Z4crzkKSQjU3hla7xqgfomC5Gz4XbR2TNjcQiRrJQ0UlKtX3X3ObRCEhuvG0Kzjklhv+Ddw6txrhKjMjiSi/Yyius/AE8TmC1p4U= host.example.com
`),
}

var PEMEncryptedKeys = []struct {
	Name              string
	EncryptionKey     string
	IncludesPublicKey bool
	PEMBytes          []byte
}{
	0: {
		Name:          "rsa-encrypted",
		EncryptionKey: "r54-G0pher_t3st$",
		PEMBytes: []byte(`-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-128-CBC,3E1714DE130BC5E81327F36564B05462

MqW88sud4fnWk/Jk3fkjh7ydu51ZkHLN5qlQgA4SkAXORPPMj2XvqZOv1v2LOgUV
dUevUn8PZK7a9zbZg4QShUSzwE5k6wdB7XKPyBgI39mJ79GBd2U4W3h6KT6jIdWA
goQpluxkrzr2/X602IaxLEre97FT9mpKC6zxKCLvyFWVIP9n3OSFS47cTTXyFr+l
7PdRhe60nn6jSBgUNk/Q1lAvEQ9fufdPwDYY93F1wyJ6lOr0F1+mzRrMbH67NyKs
rG8J1Fa7cIIre7ueKIAXTIne7OAWqpU9UDgQatDtZTbvA7ciqGsSFgiwwW13N+Rr
hN8MkODKs9cjtONxSKi05s206A3NDU6STtZ3KuPDjFE1gMJODotOuqSM+cxKfyFq
wxpk/CHYCDdMAVBSwxb/vraOHamylL4uCHpJdBHypzf2HABt+lS8Su23uAmL87DR
yvyCS/lmpuNTndef6qHPRkoW2EV3xqD3ovosGf7kgwGJUk2ZpCLVteqmYehKlZDK
r/Jy+J26ooI2jIg9bjvD1PZq+Mv+2dQ1RlDrPG3PB+rEixw6vBaL9x3jatCd4ej7
XG7lb3qO9xFpLsx89tkEcvpGR+broSpUJ6Mu5LBCVmrvqHjvnDhrZVz1brMiQtU9
iMZbgXqDLXHd6ERWygk7OTU03u+l1gs+KGMfmS0h0ZYw6KGVLgMnsoxqd6cFSKNB
8Ohk9ZTZGCiovlXBUepyu8wKat1k8YlHSfIHoRUJRhhcd7DrmojC+bcbMIZBU22T
Pl2ftVRGtcQY23lYd0NNKfebF7ncjuLWQGy+vZW+7cgfI6wPIbfYfP6g7QAutk6W
KQx0AoX5woZ6cNxtpIrymaVjSMRRBkKQrJKmRp3pC/lul5E5P2cueMs1fj4OHTbJ
lAUv88ywr+R+mRgYQlFW/XQ653f6DT4t6+njfO9oBcPrQDASZel3LjXLpjjYG/N5
+BWnVexuJX9ika8HJiFl55oqaKb+WknfNhk5cPY+x7SDV9ywQeMiDZpr0ffeYAEP
LlwwiWRDYpO+uwXHSFF3+JjWwjhs8m8g99iFb7U93yKgBB12dCEPPa2ZeH9wUHMJ
sreYhNuq6f4iWWSXpzN45inQqtTi8jrJhuNLTT543ErW7DtntBO2rWMhff3aiXbn
Uy3qzZM1nPbuCGuBmP9L2dJ3Z5ifDWB4JmOyWY4swTZGt9AVmUxMIKdZpRONx8vz
I9u9nbVPGZBcou50Pa0qTLbkWsSL94MNXrARBxzhHC9Zs6XNEtwN7mOuii7uMkVc
adrxgknBH1J1N+NX/eTKzUwJuPvDtA+Z5ILWNN9wpZT/7ed8zEnKHPNUexyeT5g3
uw9z9jH7ffGxFYlx87oiVPHGOrCXYZYW5uoZE31SCBkbtNuffNRJRKIFeipmpJ3P
7bpAG+kGHMelQH6b+5K1Qgsv4tpuSyKeTKpPFH9Av5nN4P1ZBm9N80tzbNWqjSJm
S7rYdHnuNEVnUGnRmEUMmVuYZnNBEVN/fP2m2SEwXcP3Uh7TiYlcWw10ygaGmOr7
MvMLGkYgQ4Utwnd98mtqa0jr0hK2TcOSFir3AqVvXN3XJj4cVULkrXe4Im1laWgp
-----END RSA PRIVATE KEY-----
`),
	},

	1: {
		Name:          "dsa-encrypted",
		EncryptionKey: "qG0pher-dsa_t3st$",
		PEMBytes: []byte(`-----BEGIN DSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-128-CBC,7CE7A6E4A647DC01AF860210B15ADE3E

hvnBpI99Hceq/55pYRdOzBLntIEis02JFNXuLEydWL+RJBFDn7tA+vXec0ERJd6J
G8JXlSOAhmC2H4uK3q2xR8/Y3yL95n6OIcjvCBiLsV+o3jj1MYJmErxP6zRtq4w3
JjIjGHWmaYFSxPKQ6e8fs74HEqaeMV9ONUoTtB+aISmgaBL15Fcoayg245dkBvVl
h5Kqspe7yvOBmzA3zjRuxmSCqKJmasXM7mqs3vIrMxZE3XPo1/fWKcPuExgpVQoT
HkJZEoIEIIPnPMwT2uYbFJSGgPJVMDT84xz7yvjCdhLmqrsXgs5Qw7Pw0i0c0BUJ
b7fDJ2UhdiwSckWGmIhTLlJZzr8K+JpjCDlP+REYBI5meB7kosBnlvCEHdw2EJkH
0QDc/2F4xlVrHOLbPRFyu1Oi2Gvbeoo9EsM/DThpd1hKAlb0sF5Y0y0d+owv0PnE
R/4X3HWfIdOHsDUvJ8xVWZ4BZk9Zk9qol045DcFCehpr/3hslCrKSZHakLt9GI58
vVQJ4L0aYp5nloLfzhViZtKJXRLkySMKdzYkIlNmW1oVGl7tce5UCNI8Nok4j6yn
IiHM7GBn+0nJoKTXsOGMIBe3ulKlKVxLjEuk9yivh/8=
-----END DSA PRIVATE KEY-----
`),
	},

	2: {
		Name:              "ed25519-encrypted",
		EncryptionKey:     "password",
		IncludesPublicKey: true,
		PEMBytes: []byte(`-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAACmFlczI1Ni1jdHIAAAAGYmNyeXB0AAAAGAAAABDKj29BlC
ocEWuVhQ94/RjoAAAAEAAAAAEAAAAzAAAAC3NzaC1lZDI1NTE5AAAAIIw1gSurPTDwZidA
2AIjQZgoQi3IFn9jBtFdP10/Jj7DAAAAoFGkQbB2teSU7ikUsnc7ct2aH3pitM359lNVUh
7DQbJWMjbQFbrBYyDJP+ALj1/RZmP2yoIf7/wr99q53/pm28Xp1gGP5V2RGRJYCA6kgFIH
xdB6KEw1Ce7Bz8JaDIeagAGd3xtQTH3cuuleVxCZZnk9NspsPxigADKCls/RUiK7F+z3Qf
Lvs9+PH8nIuhFMYZgo3liqZbVS5z4Fqhyzyq4=
-----END OPENSSH PRIVATE KEY-----
`),
	},
}

// SKData contains a list of PubKeys backed by U2F/FIDO2 Security Keys and their test data.
var SKData = []struct {
	Name         string
	PubKey       []byte
	HexData      []byte
	HexSignature []byte
}{
	{
		Name:         "sk-ecdsa-sha2-nistp256@openssh.com",
		PubKey:       []byte("sk-ecdsa-sha2-nistp256@openssh.com AAAAInNrLWVjZHNhLXNoYTItbmlzdHAyNTZAb3BlbnNzaC5jb20AAAAIbmlzdHAyNTYAAABBBGRNqlFgED/pf4zXz8IzqA6CALNwYcwgd4MQDmIS1GOtn1SySFObiuyJaOlpqkV5FeEifhxfIC2ejKKtNyO4CysAAAAEc3NoOg== user@host"),
		HexData:      []byte("00000020A4DE1F50DE0EF3F66DCD156C78F5C93B07EEE89D5B5A6531656E835FA1C87B323200000006736B696E6E650000000E7373682D636F6E6E656374696F6E000000097075626C69636B65790100000022736B2D65636473612D736861322D6E69737470323536406F70656E7373682E636F6D0000007F00000022736B2D65636473612D736861322D6E69737470323536406F70656E7373682E636F6D000000086E697374703235360000004104644DAA5160103FE97F8CD7CFC233A80E8200B37061CC207783100E6212D463AD9F54B248539B8AEC8968E969AA457915E1227E1C5F202D9E8CA2AD3723B80B2B000000047373683A"),
		HexSignature: []byte("0000007800000022736B2D65636473612D736861322D6E69737470323536406F70656E7373682E636F6D000000490000002016CC1A3070E180621CB206C2C6313D1CC5F094DB844A61D06001E243C608875F0000002100E4BD45D6B9DAA11489AEA8D76C222AA3FD6D50FBFFDA8049526D5D61F63B2C5601000000F9"),
	},
	{
		Name:         "sk-ssh-ed25519@openssh.com",
		PubKey:       []byte("sk-ssh-ed25519@openssh.com AAAAGnNrLXNzaC1lZDI1NTE5QG9wZW5zc2guY29tAAAAIJjzc2a20RjCvN/0ibH6UpGuN9F9hDvD7x182bOesNhHAAAABHNzaDo= user@host"),
		HexData:      []byte("000000204CFE6EA65CCB99B69348339165C7F38E359D95807A377EEE8E603C71DC3316FA3200000006736B696E6E650000000E7373682D636F6E6E656374696F6E000000097075626C69636B6579010000001A736B2D7373682D65643235353139406F70656E7373682E636F6D0000004A0000001A736B2D7373682D65643235353139406F70656E7373682E636F6D0000002098F37366B6D118C2BCDFF489B1FA5291AE37D17D843BC3EF1D7CD9B39EB0D847000000047373683A"),
		HexSignature: []byte("000000670000001A736B2D7373682D65643235353139406F70656E7373682E636F6D000000404BF5CA0CAA553099306518732317B3FE4BA6C75365BC0CB02019FBE65A1647016CBD7A682C26928DF234C378ADDBC5077B47F72381144840BF00FB2DA2FB6A0A010000009E"),
	},
}
