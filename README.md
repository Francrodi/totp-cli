# TOTP Algorithm impelementation

[Time-based One Time Password](https://datatracker.ietf.org/doc/html/rfc6238) implementation, similar to Google or Microsoft Authenticator 

## How to use

    go run cmd/authenticator [KEY]
where ```[KEY]``` is a base32 string with the secret key agreed between server and client.

Example:

    go run cmd/authenticator ABCABCABCABCABCA
