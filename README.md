# TOTP Algorithm impelementation

## How to use

    go run cmd/authenticator [KEY]
where ```[KEY]``` is a base32 string with the secret key agreed between server and client.

Example:
    go run cmd/authenticator ABCABCABCABCABCA
