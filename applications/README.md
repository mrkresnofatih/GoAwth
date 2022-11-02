# How To Run

```bash
# run goawth backend
cd GoAwthApplication
go run main.go
cd ..

# run goawth frontend
cd go-awth-fe-nextjs
yarn dev
cd ..
```

1. open goawth app player signup page: `http://localhost:3000/player`
2. signup as a player
3. open goawth app developer signup page: `http://localhost:3000/developer`
4. signup as a developer & login
5. create a developer app with success redirect as `http://localhost:1324/auth/authenticate`
6. save `applicationId` & `secret` of your created developer app & paste it in `./ClientFullstackApplication/endpoints/AuthAuthenticateEndpoint.go`

```go
authGrantResponse, err := a.AuthService.AuthenticateGrant(models.AuthAuthenticateGrantRequestModel{
    GrantId:           grantId,
    ApplicationId:     "043af8b1-914c-45e5-a01e-cef860ed6875", // GoAwth AppId
    ApplicationSecret: "cec274c5-a6e5-4055-940f-bd3173833438", // GoAwth secret
})
```

& in `./ClientFullstackApplication/endpoints/AuthLoginEndpoint.go` change the applicationid query param:

```go
return func(c echo.Context) error {
    return c.HTML(http.StatusOK, "<a href=\"http://localhost:3000/oauth2?applicationid=043af8b1-914c-45e5-a01e-cef860ed6875&granttype=grantid&scopes=openid\">login</a>")
}
```

7. run client app

```bash
# run client app
cd ClientFullstackApplication
go run main.go
cd ..
```

8. open client app login page: `http://localhost:1324/auth/login`
9. click login, you will be redirected to `http://localhost:3000/oauth2?applicationid=...`
10. login using the player you created earlier
11. you will see a consent form, click agree
12. you will be redirected back to `http://localhost:1324/auth/authenticate?grantid=...`, and will see a json data of your profile