# What is this?

This is a web app templete using "echo" framework.

# Requirement

Go : 1.11+
DB : PostgreSQL (Changeable db/setup.go if necessary)

# Sample Request

Sample Command

```
curl http://localhost:1234/api/v1/ -H "Authorization: API-TOKEN web-app-user-token"
```

Log

```
{"User":"WebAppUser","level":"info","msg":"Request accepted","time":"2020-07-20T19:28:45+09:00"}
```

# Reference

ORM : [GORM](http://gorm.io/ja_JP/docs/index.html)