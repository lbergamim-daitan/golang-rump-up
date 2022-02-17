# golang-rump-up

First, let’s download the Docker repository. Run the following:

sudo apt-get update

Followed by:
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

Next run:
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

To verify you have the repository run the following:
sudo apt-key fingerprint 0EBFCD88

And you should get something like this:
pub   4096R/0EBFCD88 2017-02-22
      Key fingerprint = 9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
uid                  Docker Release (CE deb) <docker@docker.com>
sub   4096R/F273FCD8 2017-02-22

Now to install Docker you just need to do the following:
sudo apt-get update && sudo apt-get install docker-ce

Now to get Docker Compose run:
sudo curl -L "https://github.com/docker/compose/releases/download/1.22.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && sudo chmod +x /usr/local/bin/docker-compose

To test if your installation was setup correctly, run:
docker-compose --version

And you should get something similar to:
docker-compose version 1.22.0, build 1719ceb

After

install
sudo apt install mysql-client-core-8.0

go to your repo folder
$cd <patch/to/golang-rump-up>

Connect to mysql
$ mysql -h 127.0.0.1 -P 3306 -u root -p

Select database
$ USE db

Execute sql file
$ source mysql-tables.sql 


