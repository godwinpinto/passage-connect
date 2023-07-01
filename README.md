# Passage Connect
A seamless passwordless authentication for Servers!!! Yes, that is the reason for removing the or and adding Beyond in my cover image.

SSH(secured by 1Password)->Connect(Passage)->2FA(Google Authenticator)

OR

SSH(secured by 1Password)->2FA(Google Authenticator)->Connect(Passage)

Experience the convenience and security of a seamless passwordless authentication solution for servers! Say goodbye to the hassle of managing weak or guessable passwords when it comes to server access.

Here's how the enhanced server authentication route unfolds:

SSH, already fortified by the robust security of 1Password, forms the initial layer of defence against unauthorized access.

Enter Passkeys, powered by Passage, the ultimate passwordless authentication solution. Passkeys revolutionize the way you authenticate yourself to servers. No more remembering or storing passwords manually. Instead, Passkeys use your official computer, phone, or other trusted devices as the medium for authentication.

The power of 2FA (Two-Factor Authentication) further strengthens the server authentication process. With Google Authenticator, you can enjoy an additional layer of security that verifies your identity using a time-based OTP (One-Time Password) generated on your trusted device.

By combining SSH, Passkeys, and 2FA, the entire server authentication journey becomes exceptionally robust. And the best part? You no longer need to carry separate devices or tokens to unlock access. Your existing computer or phone becomes the key to seamless and passwordless server authentication.

This was build towards the Hackathon by 1Password


### Installation Passage Connect Server Component
1. Create an AWS t2 micro instance for hosting server and making it public
2. Create an elastic IP and point to this instance
3. Point your domain name using CNAME to this elastic IP
4. Wait for domain to resolve and then ping to test "ping yourdomain" from local computer
5. Install git
6. Install Docker
7. Install Certbot to help you create SSL for your domain with LetsEncrypt
8. "mkdir cert" in home directory of user
9. copy the cert and private key (generated from certbot) into the folder cert and rename to cert.pem and key.pem
10. Follow below steps
```shell
   git clone https://github.com/godwinpinto/passage-connect.git
   cd passage-connect/setup
   #This will build and create a docker image of passage-connect-server with SSL and copy the cert from cert directory
   sudo bash install-connect-server.sh
```
You Server docker images is now ready.
```
sudo docker run -e PASSAGE_APP_ID=<YOUR_PASSAGE_APP_ID> -e REACT_APP_PASSAGE_APP_ID=<YOUR_PASSAGE_APP_ID> -e PASSAGE_API_KEY=<YOUR_PASSAGE_KEY> -p 443:443 -d passage-connect
```
This will make your server run in background now

Test by checking https://youdomain/. register and login to confirm working as expected. 
Note: You will get a message after login "No active sessions found". This means its working

### Installation Passage Connect Client
WARNING: Keep two SSH terminals open at all times so that if you are locked out then you can undo the changes
This is installation where you want passage authentication to take place

1. Create an AWS t2 micro instance for hosting server and login using the SSH terminal
2. ssh -i your_path_to_aws_private_key ec2-user@awsdomain
3. Install git, make,gcc,libpam0g-dev, amazon-linux-extras, glibc-static, tar, gzip, pam-devel
4. Do not install go from yum, instead do the following
```
curl -O https://dl.google.com/go/go1.19.2.linux-amd64.tar.gz
tar -xf go1.19.2.linux-amd64.tar.gz
#set this in in environment config 
export PATH=$PATH:$HOME/go/bin
```
5. Next building the client
```
   git clone https://github.com/godwinpinto/passage-connect.git
   cd passage-connect/setup
   #This will build the client
   sudo bash install-client.sh
```
6. create a file in users home directory ".connect"
```
PASSAGE_APP_ID=<YOUR PASSAGE APP ID>
USER=<YOUR PASSAGE USER ID WHO WILL BE VALIDATED>
```
7. For PAM authentication test you need make an entry in /etc/pam.d/sshd file
8. Add "session required pam_passage.so"
9. For sshd mode you need to make entry in ssh_config file located in /etc/ssh/sshd_config
10. ForceCommand /home/sshd-force-command.sh
11. Now restart sshd "systemctl service restart sshd"
NOTE: Dont use PAM mode and SSH mode together. just use one of the mode (you dont want dual authentication)

Testing:
1. Login to your client with ssh
2. Go to you web page and login. "You will get a message You are not logged in and return to your terminal"

##Notes:
If the documentation lacks / misses installation steps, feel free to point out
