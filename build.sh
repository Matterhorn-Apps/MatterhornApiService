# https://stackoverflow.com/questions/59144120/gopath-go-mod-exists-but-should-not-in-aws-elastic-beanstalk
sudo rm /var/app/current/go.*

go build -o bin/application .