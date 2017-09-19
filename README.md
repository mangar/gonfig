# Gonfig

Get the config from a private repoository.


## Extra packages

  go get gopkg.in/kyokomi/emoji.v1


## Debug mode

First export a ENV variable: ```export FlagDebug=true``` and then run the program



## Configuration

__1__

You can create a config.yml file located at: ```~/.gonfig/config.yml``` and put the and complete that with: 

```

skol_circles:
  key: ""
  
```


__2__

Use the same file structure and specify when calling the command.




## TODO

- help no formato de command line
- config.json com uma lista de projetos e nao apenas um
- retornar erro se nao encontrar o projeto no config.
