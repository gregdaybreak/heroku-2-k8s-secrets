# heroku-2-k8s-secrets
Tool to allow a person to Get Config Vars from Heroku and create a Secret file for Kubernetes

Usage:

1. Get your API key from Heroku at https://dashboard.heroku.com/account

2. Run ```go get github.com/bgentry/heroku-go``` to download, build, and install the heroku-go package.

3. build the go package with ```go build heroku-2-k8s-secrets.go```

4. Run the app
   ```./heroku-2-k8s-secrets -u <username> -p <api-key> -a <app-name> -s <secret name>```
   
5. You will now have a yaml file that you can apply to your cluster with ```kubectl apply -f <file>```
