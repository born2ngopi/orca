# Orca

Orca is tools for auto generate commit message with local llm with ollama.

its very simple to use, just run orca in your terminal and it will generate commit message for you.

> Note: you need to have ollama and llm installed in your local machine. and its depend on your local llm to generate the message.

```
$ orca
```

or if you want preview the message before commit, you can use --preview flag

```
$ orca --preview
```

you can also change the model with flag --model or -m

```
$ orca --model=codellama
```