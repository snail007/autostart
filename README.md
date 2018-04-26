# Autostart

[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/snail007/autostart/) [![license](https://img.shields.io/github/license/snail007/autostart.svg?style=plastic)]() [![download_count](https://img.shields.io/github/downloads/snail007/autostart/total.svg?style=plastic)](https://github.com/snail007/autostart/releases) [![download](https://img.shields.io/github/release/snail007/autostart.svg?style=plastic)](https://github.com/snail007/autostart/releases)  
  
autostart tools to set your application auto startup on desktop login,only for desktop version of linux , windows , mac.     

# Behavior
On Linux and BSD, it creates a .desktop file in $XDG_CONFIG_HOME/autostart (i.e. $HOME/.config/autostart). See http://askubuntu.com/questions/48321/how-do-i-start-applications-automatically-on-login  
On macOS, it creates a launchd job. See http://blog.gordn.org/2015/03/implementing-run-on-login-for-your-node.html  
On Windows, it creates a link to the program in %USERPROFILE%\Start Menu\Programs\Startup  

# Usage  
On linux and mac , the command is autostart.   
On windows , the command is autostart.exe.  

## 1.Enable Application autostart.   
***linux&mac***   
`autostart enable -k "demokey" -n "my demo application" -c "echo \"autostart\">~/autostart.txt"`   
help:    
`autostart enable --help`  

***windows***    
`autostart.exe enable -k test -n test -c "c:\\windows\explorer.exe c:"`   
help:    
`autostart.exe enable --help`   

## 2.Disable Application autostart.   
***linux&mac***   
`autostart disable -k "demokey"`    
help:     
`autostart disable --help`   

***windows***    
`autostart.exe disable -k "demokey"`    
help:     
`autostart.exe disable --help`   


