# GASKN

GASKN stands for Go Api Starter Kit eNgine, is a simple go starterkit to speed up your job. this project is still under development, you can make a Pull Request into this project.

# How to run this project?

To run this project copy the .env.example file into .env, then do the configuration as you need

After you create the configuration file, create a database in MySQL or PostgreSQL with the appropriate name in the configuration file above.

Run the command in the root directory

```
go run main.go
```

or if you use makefile run the following command

```
make watch
```

This command has a "hot reload" feature, but you will need the <b>reflect</b> library to run the command
