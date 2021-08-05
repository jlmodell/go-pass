# go-pass
simple cmd line password generator/manager written in go [touch,grep,del]

*build*:  
&nbsp;&nbsp;|&nbsp;&nbsp;go build .  
&nbsp;&nbsp;|&nbsp;&nbsp;go build -o "something" .  

*commands*:  
&nbsp;&nbsp;|&nbsp;&nbsp;touch "key::string" (optional: "password::string") --length::int, -l::int  
&nbsp;&nbsp;|&nbsp;&nbsp;grep "key::string"  
&nbsp;&nbsp;|&nbsp;&nbsp;del "key::string"  

(**windows**) stores master key at *c:\temp\go-pass\\.skf*  
(linux) TODO  
(mac) TODO  

master key is used to encrypt and decrypt your passwords and each key is stored as *"\\.[key]"*
