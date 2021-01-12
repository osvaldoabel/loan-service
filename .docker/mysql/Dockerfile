FROM mysql:5.7

COPY smartmei.sql /docker-entrypoint-initdb.d/

ONBUILD RUN chown -R mysql:mysql ./dbdata
