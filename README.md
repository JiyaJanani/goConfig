# goConfig

We can create our own config file.

main.go
config
|- config.development.json
|- config.production.json

type Configuration struct {
    Port              int
    Static_Variable   string
    Connection_String string
}

Our config file
{
     "Port": 8080
     "Static_Variable": "static"
}

we have to get information from the json :
//filename is the path to the json config file
file, err := os.Open(filename) if err != nil {  return err }  
decoder := json.NewDecoder(file) 
err = decoder.Decode(&configuration) 
if err != nil {  return err }

 from the environment variables
configuration.Connection_String = os.Getenv("Connection_String")

If you don’t want to manage it by yourself, you can use gonfig or viper.

Here we are using VIPER

go get github.com/spf13/viper
On your code you will need to declare your variables

var ( 
    Port            string                   
)

Viper will be able to read in a toml file, or in the env variables, so if you want to have a config file for your Development environment, but use the env vars, for Production, you can write a code like :

if os.Getenv("ENVIRONMENT") == "DEV" {        
     viper.SetConfigName("config")  
     viper.SetConfigType("toml") 
    viper.AddConfigPath(filepath.Dir(dirname))  
     viper.ReadInConfig() 
} else {  viper.AutomaticEnv() }
And then to init your variable :

// To add a default value :
viper.SetDefault("APP_PORT", "8080")
//To get from the toml file or env var
Port = viper.GetString("APP_PORT")


refer link 
https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152
