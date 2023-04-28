# SSOService

Single sign on service as part of E_uprava class project assigment. Developed by [petar__k](https://www.linkedin.com/in/petar-komord%C5%BEi%C4%87-23765a233/)

## Getting started
To run this project you will need:
- [Docker+Docker compose](https://www.docker.com/)
- [GoLang](https://go.dev/dl/) version: 1.18 

To start your service you can see attached docker-compose.yaml that I wrote for all services(including this one)
in [this repo](https://github.com/EUPRAVA-TIM1/DockerCompose). Ports can be seen in `.env` there or in `config.go` file of this project
To be able to retrieve secret and validate token that has been generated in your service you will first need to manually 
add your service as a issuer in redis db. You can do that by following this set of instructions:

After running compose check for your redis container id by typing following command:
```
- docker ps -a
```
When you find your container id copy it and paste it in this command as a container-id:
```
docker exec -it <your-container-id> sh
```
Once inside terminal type in following command for as many services as you would like.Name isn't important(except that it needs to start with **issuer:**), but it would be great if were service name + something random:
```
mongo-cli set issuer:<your-issuer-name> "true"
```
Once that's set you are good to go since MySql data will be filled with script upon first compose-up.
#### <u>Note: all passwords in sql script are bycript encoded and are: Lozinka123</u>

## Endpoints
- `GET /sso/Login` Loges in user and returns generated JWT token as a response.\

**Expects** json in this format:
```
{
    "JMBG" : <jmbg-here>,
    "Password" : <password-here>
}
```
**Returns** json in this format:
```
{
    "token" : <jwt-token-here>
}
```
- `GET /sso/Secret` Returns secret that is currently in use for encrypting JWT tokens (not using HTTPS for sake of simplicity)\

  **Expects** `X-Service-Name` header (same as name you provided for a redis as issuer) so it will return secret only to issuers registered in db:\
**Returns** json in this format:
```
{
    "secret" : <jwt-token-here>,
    "expiresAt" : <date and time until secret is valid>
}
```

- `GET /sso/Whoami` Returns currently logged in gradjanin based on JWT subject\

  **Expects** `Authorization` header with JWT token(with or without Barrer)\
  **Returns** json in this format:
```
{
    "ime": "",
    "prezime": "",
    "jmbg": "",
    "adresa": "",
    "brojTelefona": "",
    "email": "",
    "opstina": {
        "PTT": "",
        "Naziv": ""
    }
}
```

- `GET /sso/User/{jmbg}` Returns currently logged in gradjanin based on JWT subject\

  **Expects** `jmbg` as a part of an url\
  **Returns** json in this format:
```
{
    "ime": "",
    "prezime": "",
    "jmbg": "",
    "adresa": "",
    "brojTelefona": "",
    "email": "",
    "opstina": {
        "PTT": "",
        "Naziv": ""
    }
}
```