# JWT Command line

Command line application for testing JWT. You can generate JWT and sign it.
## Commands
### Help
```shell
./jwt help
```

### Generate Public/Private Key
```shell
jwt genkeys --keypath . --keyname test
```
This will generate private and public key in current directory.
Named testprivate.pem and testpublic.pem. 
These keys can be used for signing and verifying JWT (testing purposes only).

## Generate token

`gen` command will generate sample token.
```shell
./jwt gen 
```
This token does not contain custom claims. Just standard
claims. Output looks like this

```
Header
        typ : JWT 
        alg : HS256 
Custom Claims
Standard Claims
         Id: 1 
         Audience: Recipient
         Issuer: Sample
         Issued at: 2022-11-23T09:15:58+01:00
         Not Before: 2022-11-23T09:15:58+01:00
         Expires at: 2023-11-23T09:15:58+01:00
```

Change the signing method
```shell
./jwt gen --signingmethod HS384
```

 

# JWT Samples
JSON Web tokens defined in [RFC 7519](https://datatracker.ietf.org/doc/html/rfc7519) . 
JWT represents set of claims. 
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
