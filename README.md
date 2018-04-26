# autostart
autostart tools to set your application auto startup on desktop login,only for desktop version of linux&windows&mac.     

# Behavior
On Linux and BSD, it creates a .desktop file in $XDG_CONFIG_HOME/autostart (i.e. $HOME/.config/autostart). See http://askubuntu.com/questions/48321/how-do-i-start-applications-automatically-on-login  
On macOS, it creates a launchd job. See http://blog.gordn.org/2015/03/implementing-run-on-login-for-your-node.html  
On Windows, it creates a link to the program in %USERPROFILE%\Start Menu\Programs\Startup  
