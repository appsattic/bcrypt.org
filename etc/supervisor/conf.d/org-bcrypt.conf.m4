[program:org-bcrypt]
command = sudo -E -u __USER__ __PWD__/server
directory = __PWD__
user = __USER__
autostart = true
autorestart = true
startsecs=10
startretries= 3
stopwaitsec= 600
stdout_logfile = /var/log/__NAME__/stdout.log
stderr_logfile = /var/log/__NAME__/stderr.log
environment =
    PORT=__PORT__,
    HOME="/home/chilts",
    USER="chilts"
