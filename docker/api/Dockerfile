FROM golang:1.8


ADD env.sh /usr/local/bin/env.sh
RUN chmod u+x /usr/local/bin/env.sh && /usr/local/bin/env.sh

ADD ../../../augehorus/ ~/develop/src/github.com/KoganezawaRyouta/augehorus

ADD run.sh /usr/local/bin/run.sh
RUN chmod u+x /usr/local/bin/run.sh
ENTRYPOINT ["/usr/local/bin/run.sh"]
