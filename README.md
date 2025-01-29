# Telegram Bot Template

## Basic Go Telegram Bot Structure to Start Developing Your Bot

```bash
git clone https://github.com/morf1lo/tgbotbase.git
```

#### Replace `module github.com/morf1lo/tgbotbase` line from `go.mod` file with your package name

### Structure

- `/cmd`/`bot` - *`package main` for starting your bot*
- `/i18n` - ***<*language*>.yaml** files with bot texts translation*
- `/internal` - *all your bot business logic*
    - `/config` - *configuration **structs***
    - `/handler` - *telegram actions handlers*
    - `/localization` - *dynamic text translation*
    - `/model` - *your database models*
    - `/repository` - *working with databases*
        - `/postgres` - *working with PostgreSQL (for example)*
        - `/redisrepo` - *Redis*
    - `/service` - ***main** logic, that you will use in **handlers***

`.env` - *your environment variables*
`bot.yaml` - *your bot config*
