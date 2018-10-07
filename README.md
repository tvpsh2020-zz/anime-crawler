# Anime Crawler

Get anime information from DMHY website periodically, server will send notification email to you if the anime have update information. For now, this application is better run on local computer because the config file only been read once when initializing the application, I will keep enhancing this application.

## Requirement

- Golang 1.7+

## Install

### 1. Pull from GitHub

```
$ go get github.com/tvpsh2020/anime-crawler
```

### 2. Copy config file from example

```
$ cp ./config.yaml.default config.yaml
```

### 3. Set up SMTP config for notification

See more detail to get your GMail work on: https://www.digitalocean.com/community/tutorials/how-to-use-google-s-smtp-server

Of course you can use any other SMTP provider.

### 4. Set up query string 

I wrote two examples in the config file, `queryString` will be used to request the anime search result page, `keywords` will be used to filter out your target from the result page, you can have more than one `keywords` at same time, the methodology of using `keywords` is UNION

## Monitoring

You can see memory status and config status in these route:

- localhost:8080/api/stat
- localhost:8080/api/config
