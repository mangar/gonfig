# Gonfig

Get the config from a private repoository.


## Usage

    go run main.go projectDemo ../output microservices/all/development microservices/users/development


## Making

    make
    make install


## Extra packages

    go get gopkg.in/kyokomi/emoji.v1




## Debug mode

First export a ENV variable: ```export FlagDebug=true``` and then run the program




## Configuration

__1__

You can create a config.yml file located at: ```~/.gonfig/config.json``` and put the and complete that with: 

```

[
  {
    "project":"ProjectName",
    "repo" : "https://USERNAME:APP_PASSWORD@bitbucket.org/REPOSITORY_NAME.git"
  }
]
  
```


__2__

Use the same file structure and specify when calling the command.



## TODO
