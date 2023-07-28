# JWT Command line

Command line application for testing JWT. You can generate JWT and sign it.
You can also generate public and private key for testing purposes.
## Commands
### Help
```shell
./jwt help
```

### Generate Public/Private Key
`genkeys` will generate private and public key to stdout.

Flags
```
--keypath string   path to directory where keys will be stored (default ".")
--privatekey string   private key file name (default "private.pem")
--publickey string   public key file name (default "public.pem")

```

```shell
jwt genkeys
```
Will generate private and public key to stdout.

Specify file path, this will generate `private.pem` and `public.pem` in current directory.
```shell
jwt genkeys --keypath .
```

Verify keys
```shell
openssl rsa -in path/to/rsa_key.pem -text -noout
```

Specify file name
```shell
./jwt genkeys --keypath . --privatekey pk --publickey puk
```

This will generate private and public key in current directory.
These keys can be used for signing and verifying JWT (testing purposes only).

## Generate Sample token

`gen` command will generate sample token.
```shell
./jwt gen 
```
This token does not contain custom claims. Just standard
claims. Output looks like this

```
=== Generating Simple Token ===
Header
        typ : JWT 
        alg : HS256 
Custom Claims
Standard Claims
         Id: 1 
         Audience: Recipient
         Issuer: Sample
         Issued at: 1970-01-20T14:30:54+01:00
         Not Before: 1970-01-20T14:32:20+01:00
         Expires at: 1970-01-20T23:20:49+01:00
 
Signed string: 
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
.eyJDdXN0b21DbGFpbXMiOnt9LCJhdWQiOiJSZWNpcGllbnQiLCJleHAiOjE3MjIwNDkzMzUsImp0aSI6IjEiLCJpYXQiOjE2OTAyNTQxMzUsImlzcyI6IlNhbXBsZSIsIm5iZiI6MTY5MDM0MDUzNSwic3ViIjoiVXNlciJ9
.XSNQLGX6Gfdk_PToao9KBrHAC7aWBeqjaT3zDwWrfR4
```
Default secret is `AllYourBase`
You can change the secret with `--secret` flag.
```shell
./jwt gen --secret mysecret
```

Change the signing method
```shell
./jwt gen --signingmethod HS384
```

## Encode JWT

```shell
./jwt encode '{"sub":"1234567890","name":"John Doe","admin":true}' AllYourBase
```


## Decode JWT

```shell
./jwt decode eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdWQiLCJleHAiOjE3MjIxNTIxNjgsImlhdCI6MTY5MDM1Njk2OCwiaXNzIjoiaXNzIiwibmJmIjoxNjkwNDQzMzY4LCJzdWIiOiJzdWIifQ._1L7ZTk4QpybaCk4rx2pgTwl1cGaRl8W9AUH_T3TfT0 AllYourBase
```

# JWT Samples
JSON Web tokens defined in [RFC 7519](https://datatracker.ietf.org/doc/html/rfc7519) . 
JWT represents set of claims. 

JWT stands for "JSON Web Token." It is a compact and self-contained way of representing 
information between two parties in a secure manner as a JSON object. 
JWT is commonly used for authentication and authorization in web applications and APIs.

The JWT is typically issued by an authentication server when a user logs in or 
requests access to certain protected resources. The client (usually a web browser or mobile app) 
then includes the JWT in the Authorization header when making subsequent requests to the server. 
The server can then validate the JWT to authenticate the user and authorize access to the requested resources.

Since JWTs are digitally signed, they are tamper-proof. This means that the server can trust the information 
contained in the token without the need to store session information on the server side. 
This makes JWT a stateless and scalable approach for user authentication and authorization in distributed systems. 
However, it's essential to keep the secret used for signing JWTs secure to prevent unauthorized access and tampering.
## JWT Structure
JWT is a string consisting of three parts separated by dots.
```
header.payload.signature
```
### Header
Header is a JSON object containing information about the token.
```json
{
  "typ": "JWT",
  "alg": "HS256"
}
```
What is the purpose of the header? It is used to tell the receiver
how to validate the token. In this case the token is signed with HMAC
using SHA-256.
What is the purpose of the typ? It is used to tell the receiver
what is the type of the token. In this case it is JWT.
What other types are there? There is JWE (JSON Web Encryption).

### Payload
Payload is a JSON object containing claims.
```json
{
  "sub": "1234567890",
  "name": "John Doe",
  "admin": true
}
```
### Signature
Signature is a hash of header and payload. 
```shell
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```
## JWT Claims
Claims are key-value pairs holding information about a subject.

### Registered Claims
Registered claims are predefined claims.
```
iss (issuer)
sub (subject)
aud (audience)
exp (expiration time)
nbf (not before)
iat (issued at)
jti (JWT ID)
```
### Public Claims
Public claims are defined by RFC 7519. 
```
https://www.iana.org/assignments/jwt/jwt.xhtml
```
### Private Claims
Private claims are custom claims defined by the user.

## JOSE Header
Javascript Object Signing and Encryption
Example JOSE header
```json
{
  "typ": "JWT",
  "alg": "HS256"
}
```

MAC is message authentication code.
HMAC is hash based message authentication code.
HMAC is symmetric signature, you have to pick a secret phrase
which will be used in signing. In other words you have to have
a shared key, so the other party can validate the message.

To overcome the issue with shared keys (how to share it in secure manner) 
you can also use public key cryptography.

Signing algorithms
```
+--------------+-------------------------------+--------------------+
| "alg" Param  | Digital Signature or MAC      | Implementation     |
| Value        | Algorithm                     | Requirements       |
+--------------+-------------------------------+--------------------+
| HS256        | HMAC using SHA-256            | Required           |
| HS384        | HMAC using SHA-384            | Optional           |
| HS512        | HMAC using SHA-512            | Optional           |
| RS256        | RSASSA-PKCS1-v1_5 using       | Recommended        |
|              | SHA-256                       |                    |
| RS384        | RSASSA-PKCS1-v1_5 using       | Optional           |
|              | SHA-384                       |                    |
| RS512        | RSASSA-PKCS1-v1_5 using       | Optional           |
|              | SHA-512                       |                    |
| ES256        | ECDSA using P-256 and SHA-256 | Recommended+       |
| ES384        | ECDSA using P-384 and SHA-384 | Optional           |
| ES512        | ECDSA using P-521 and SHA-512 | Optional           |
| PS256        | RSASSA-PSS using SHA-256 and  | Optional           |
|              | MGF1 with SHA-256             |                    |
| PS384        | RSASSA-PSS using SHA-384 and  | Optional           |
|              | MGF1 with SHA-384             |                    |
| PS512        | RSASSA-PSS using SHA-512 and  | Optional           |
|              | MGF1 with SHA-512             |                    |
| none         | No digital signature or MAC   | Optional           |
|              | performed                     |                    |
+--------------+-------------------------------+--------------------+
```

## JWT Claims
Claims are key-value pairs holding information about a subject.

Commandline application for testing JWT. You can generate JWT and sign it.
