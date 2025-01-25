# use in n/vim to restart on save:
# :autocmd BufWritePost * silent! !./autoload.sh
#!/bin/bash
pkill sigma-firma.com || true
go build -o sigma-firma.com
echo http://localhost:10529
./sigma-firma.com >> log.txt 2>&1 &
